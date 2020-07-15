// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package teams

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	//goteamsnotify "gopkg.in/dasrick/go-teams-notify.v1"
	goteamsnotify "github.com/atc0005/go-teams-notify"
)

// logger is a package logger that can be enabled from client code to allow
// logging output from this package when desired/needed for troubleshooting
var logger *log.Logger

// Newline patterns stripped out of text content sent to Microsoft Teams (by
// request) and replacement break value used to provide equivalent formatting.
const (

	// CR LF \r\n (windows)
	windowsEOLActual  = "\r\n"
	windowsEOLEscaped = `\r\n`

	// CF \r (mac)
	macEOLActual  = "\r"
	macEOLEscaped = `\r`

	// LF \n (unix)
	unixEOLActual  = "\n"
	unixEOLEscaped = `\n`

	// Used by Teams to separate lines
	breakStatement = "<br>"
)

// In practice, all new webhook URLs appear to use the outlook.office.com
// FQDN. However, some older guides, and even the current official
// documentation, use outlook.office365.com in their webhook URL examples.
// https://docs.microsoft.com/en-us/outlook/actionable-messages/send-via-connectors
const webhookURLOfficecomPrefix = "https://outlook.office.com"
const webhookURLOffice365Prefix = "https://outlook.office365.com"
const webhookURLOfficialDocsSampleURI = "webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c"

// Build a regular expression that we can use to validate incoming webhook
// URLs provided by the user.
//
// Note: The regex allows for capital letters in the GUID patterns. This is
// allowed based on light testing which shows that mixed case works and the
// assumption that since Teams and Office 365 are Microsoft products case
// would be ignored (e.g., Windows, IIS do not consider 'A' and 'a' to be
// different).
var validWebhookURLRegex = `^https:\/\/outlook.office(?:365)?.com\/webhook\/[-a-zA-Z0-9]{36}@[-a-zA-Z0-9]{36}\/IncomingWebhook\/[-a-zA-Z0-9]{32}\/[-a-zA-Z0-9]{36}$`

// Even though Microsoft Teams doesn't show the additional newlines,
// https://messagecardplayground.azurewebsites.net/ DOES show the results
// as a formatted code block. Including the newlines now is an attempt at
// "future proofing" the codeblock support in MessageCard values sent to
// Microsoft Teams.
const (

	// msTeamsCodeBlockSubmissionPrefix is the prefix appended to text input
	// to indicate that the text should be displayed as a codeblock by
	// Microsoft Teams.
	msTeamsCodeBlockSubmissionPrefix string = "\n```\n"
	// msTeamsCodeBlockSubmissionPrefix string = "```"

	// msTeamsCodeBlockSubmissionSuffix is the suffix appended to text input
	// to indicate that the text should be displayed as a codeblock by
	// Microsoft Teams.
	msTeamsCodeBlockSubmissionSuffix string = "```\n"
	// msTeamsCodeBlockSubmissionSuffix string = "```"

	// msTeamsCodeSnippetSubmissionPrefix is the prefix appended to text input
	// to indicate that the text should be displayed as a code formatted
	// string of text by Microsoft Teams.
	msTeamsCodeSnippetSubmissionPrefix string = "`"

	// msTeamsCodeSnippetSubmissionSuffix is the suffix appended to text input
	// to indicate that the text should be displayed as a code formatted
	// string of text by Microsoft Teams.
	msTeamsCodeSnippetSubmissionSuffix string = "`"
)

func init() {

	// Disable logging output by default unless client code explicitly
	// requests it
	logger = log.New(os.Stderr, "[send2teams/teams] ", 0)
	logger.SetOutput(ioutil.Discard)

}

// EnableLogging enables logging output from this package. Output is muted by
// default unless explicitly requested (by calling this function).
func EnableLogging() {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetOutput(os.Stderr)
}

// DisableLogging reapplies default package-level logging settings of muting
// all logging output.
func DisableLogging() {
	logger.SetFlags(0)
	logger.SetOutput(ioutil.Discard)
}

// TryToFormatAsCodeBlock acts as a wrapper for FormatAsCodeBlock. If an
// error is encountered in the FormatAsCodeBlock function, this function will
// return the original string, otherwise if no errors occur the newly formatted
// string will be returned.
func TryToFormatAsCodeBlock(input string) string {

	result, err := FormatAsCodeBlock(input)
	if err != nil {
		logger.Printf("TryToFormatAsCodeBlock: error occurred when calling FormatAsCodeBlock: %v\n", err)
		logger.Println("TryToFormatAsCodeBlock: returning original string")
		return input
	}

	logger.Println("TryToFormatAsCodeBlock: no errors occurred when calling FormatAsCodeBlock")
	return result
}

// TryToFormatAsCodeSnippet acts as a wrapper for FormatAsCodeSnippet. If
// an error is encountered in the FormatAsCodeSnippet function, this function will
// return the original string, otherwise if no errors occur the newly formatted
// string will be returned.
func TryToFormatAsCodeSnippet(input string) string {

	result, err := FormatAsCodeSnippet(input)
	if err != nil {
		logger.Printf("TryToFormatAsCodeSnippet: error occurred when calling FormatAsCodeBlock: %v\n", err)
		logger.Println("TryToFormatAsCodeSnippet: returning original string")
		return input
	}

	logger.Println("TryToFormatAsCodeSnippet: no errors occurred when calling FormatAsCodeSnippet")
	return result
}

// FormatAsCodeBlock accepts an arbitrary string, quoted or not, and calls a
// helper function which attempts to format as a valid Markdown code block for
// submission to Microsoft Teams
func FormatAsCodeBlock(input string) (string, error) {

	if input == "" {
		return "", errors.New("received empty string, refusing to format")
	}

	result, err := formatAsCode(
		input,
		msTeamsCodeBlockSubmissionPrefix,
		msTeamsCodeBlockSubmissionSuffix,
	)

	return result, err

}

// FormatAsCodeSnippet accepts an arbitrary string, quoted or not, and calls a
// helper function which attempts to format as a single-line valid Markdown
// code snippet for submission to Microsoft Teams
func FormatAsCodeSnippet(input string) (string, error) {
	if input == "" {
		return "", errors.New("received empty string, refusing to format")
	}

	result, err := formatAsCode(
		input,
		msTeamsCodeSnippetSubmissionPrefix,
		msTeamsCodeSnippetSubmissionSuffix,
	)

	return result, err
}

// formatAsCode is a helper function which accepts an arbitrary string, quoted
// or not, a desired prefix and a suffix for the string and attempts to format
// as a valid Markdown formatted code sample for submission to Microsoft Teams
func formatAsCode(input string, prefix string, suffix string) (string, error) {

	var err error
	var byteSlice []byte

	switch {

	// required; protects against slice out of range panics
	case input == "":
		return "", errors.New("received empty string, refusing to format as code block")

	// If the input string is already valid JSON, don't double-encode and
	// escape the content
	case json.Valid([]byte(input)):
		logger.Printf("formatAsCode: input string already valid JSON; input: %+v", input)
		logger.Printf("formatAsCode: Calling json.RawMessage([]byte(input)); input: %+v", input)

		// FIXME: Is json.RawMessage() really needed if the input string is *already* JSON?
		// https://golang.org/pkg/encoding/json/#RawMessage seems to imply a different use case.
		byteSlice = json.RawMessage([]byte(input))
		//
		// From light testing, it appears to not be necessary:
		//
		// logger.Printf("formatAsCode: Skipping json.RawMessage, converting string directly to byte slice; input: %+v", input)
		// byteSlice = []byte(input)

	default:
		logger.Printf("formatAsCode: input string not valid JSON; input: %+v", input)
		logger.Printf("formatAsCode: Calling json.Marshal(input); input: %+v", input)
		byteSlice, err = json.Marshal(input)
		if err != nil {
			return "", err
		}
	}

	logger.Println("formatAsCode: byteSlice as string:", string(byteSlice))

	var prettyJSON bytes.Buffer

	logger.Println("formatAsCode: calling json.Indent")
	err = json.Indent(&prettyJSON, byteSlice, "", "\t")
	if err != nil {
		return "", err
	}
	formattedJSON := prettyJSON.String()

	logger.Println("formatAsCode: Formatted JSON:", formattedJSON)

	// handle both cases: where the formatted JSON string was not wrapped with
	// double-quotes and when it was
	codeContentForSubmission := prefix + strings.Trim(formattedJSON, "\"") + suffix

	logger.Printf("formatAsCode: formatted JSON as-is:\n%s\n\n", formattedJSON)
	logger.Printf("formatAsCode: formatted JSON wrapped with code prefix/suffix: \n%s\n\n", codeContentForSubmission)

	// err should be nil if everything worked as expected
	return codeContentForSubmission, err

}

// ConvertEOLToBreak converts \r\n (windows), \r (mac) and \n (unix) into <br>
// HTML/Markdown break statements
func ConvertEOLToBreak(s string) string {

	logger.Printf("ConvertEOLToBreak: Received %#v", s)

	s = strings.ReplaceAll(s, windowsEOLActual, breakStatement)
	s = strings.ReplaceAll(s, windowsEOLEscaped, breakStatement)
	s = strings.ReplaceAll(s, macEOLActual, breakStatement)
	s = strings.ReplaceAll(s, macEOLEscaped, breakStatement)
	s = strings.ReplaceAll(s, unixEOLActual, breakStatement)
	s = strings.ReplaceAll(s, unixEOLEscaped, breakStatement)

	logger.Printf("ConvertEOLToBreak: Returning %#v", s)

	return s
}

// SendMessage is a wrapper function for setting up and using the
// goteamsnotify client to send a message card to Microsoft Teams via a
// webhook URL.
func SendMessage(ctx context.Context, webhookURL string, message goteamsnotify.MessageCard, retries int, retriesDelay int) error {

	// NOTE: The caller is responsible for setting the desired context timeout

	// init the client
	mstClient := goteamsnotify.NewClient()

	var result error

	// initial attempt + number of specified retries
	attemptsAllowed := 1 + retries

	// attempt to send message to Microsoft Teams, retry specified number of
	// times before giving up
	for attempt := 1; attempt <= attemptsAllowed; attempt++ {

		// While the context is passed to mstClient.SendWithContext and it
		// should ensure that it is respected, we check here at the start of
		// the loop iteration (either first or subsequent) in order to return
		// early in an effort to prevent undesired message attempts
		if ctx.Err() != nil {
			msg := fmt.Sprintf(
				"SendMessage: context cancelled or expired: %v; aborting message submission after %d of %d attempts",
				ctx.Err().Error(),
				attempt,
				attemptsAllowed,
			)

			// if this is set, we're looking at the second (incomplete)
			// iteration
			if result != nil {
				msg += ": " + result.Error()
			}

			logger.Println(msg)
			return fmt.Errorf(msg)
		}

		// the result from the last attempt is returned to the caller
		result = mstClient.SendWithContext(ctx, webhookURL, message)
		if result != nil {

			ourRetryDelay := time.Duration(retriesDelay) * time.Second

			errMsg := fmt.Errorf(
				"SendMessage: Attempt %d of %d to send message failed: %v",
				attempt,
				attemptsAllowed,
				result,
			)

			logger.Println(errMsg.Error())

			// apply retry delay if our context hasn't been cancelled yet,
			// otherwise continue with the loop to allow context cancellation
			// handling logic to be applied
			if ctx.Err() == nil {
				logger.Printf(
					"SendMessage: Context not cancelled yet, applying retry delay of %v",
					ourRetryDelay,
				)
				time.Sleep(ourRetryDelay)
			}

			continue
		}

		logger.Printf(
			"SendMessage: successfully sent message after %d of %d attempts\n",
			attempt,
			attemptsAllowed,
		)
		break
	}

	return result
}

// validateWebhookLength ensures that at least the prefix + SOMETHING is
// present; test against the shorter of the two known prefixes
func validateWebhookLength(webhookURL string) error {

	// FIXME: This is made redundant by the prefix check

	if len(webhookURL) <= len(webhookURLOfficecomPrefix) {
		return fmt.Errorf("incomplete webhook URL: provided URL %q shorter than or equal to just the %q URL prefix",
			webhookURL,
			webhookURLOfficecomPrefix,
		)
	}

	return nil
}

// validateWebhookURLRegex applies a regular expression pattern check against
// the provided webhook URL to ensure that the URL matches the expected
// pattern.
func validateWebhookURLRegex(webhookURL string) error {

	// TODO: Consider retiring this validation check due to reliance on fixed
	// pattern (subject to change?)
	// This is fairly tight validation and will likely require future tending
	matched, err := regexp.MatchString(validWebhookURLRegex, webhookURL)
	if !matched {
		return fmt.Errorf(
			"webhook URL does not match expected pattern;\n"+
				"got: %q\n"+
				"expected webhook URL in one of these formats:\n"+
				"  * %q\n"+
				"  * %q\n"+
				"error: %v",
			webhookURL,
			webhookURLOfficecomPrefix+"/"+webhookURLOfficialDocsSampleURI,
			webhookURLOffice365Prefix+"/"+webhookURLOfficialDocsSampleURI,
			err,
		)
	}

	return nil
}

// ValidateWebhook applies validation checks to the specified webhook,
// returning an error for any detected issues.
func ValidateWebhook(webhookURL string) error {

	if err := validateWebhookLength(webhookURL); err != nil {
		return err
	}

	if err := validateWebhookURLRegex(webhookURL); err != nil {
		return err
	}

	// Indicate that we didn't spot any problems
	return nil

}
