package teams

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	//goteamsnotify "gopkg.in/dasrick/go-teams-notify.v1"
	goteamsnotify "github.com/atc0005/go-teams-notify"
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

// TODO: Why is the double leading slash necessary to match on escape
// sequences in order to replace them?
//
// A: Convert double-quoted strings to backtick-quoted strings, replace
// double-backslash with single-backslash as desired.

// Used by Teams to separate lines
const breakStatement = "<br>"

// CR LF \r\n (windows)
const windowsEOLActual = "\r\n"
const windowsEOLEscaped = `\r\n`

// CF \r (mac)
const macEOLActual = "\r"
const macEOLEscaped = `\r`

// LF \n (unix)
const unixEOLActual = "\n"
const unixEOLEscaped = `\n`

// ConvertEOLToBreak converts \r\n (windows), \r (mac) and \n (unix) into <br>
// HTML/Markdown break statements
func ConvertEOLToBreak(s string) string {

	//log.Printf("ConvertEOLToBreak: Received %q", s)

	s = strings.ReplaceAll(s, windowsEOLActual, breakStatement)
	s = strings.ReplaceAll(s, windowsEOLEscaped, breakStatement)
	s = strings.ReplaceAll(s, macEOLActual, breakStatement)
	s = strings.ReplaceAll(s, macEOLEscaped, breakStatement)
	s = strings.ReplaceAll(s, unixEOLActual, breakStatement)
	s = strings.ReplaceAll(s, unixEOLEscaped, breakStatement)

	//log.Printf("ConvertEOLToBreak: Returning %q", s)

	return s
}

// SendMessage is a wrapper function for setting up and using the
// goteamsnotify client to send a message card to Microsoft Teams via a
// webhook URL.
func SendMessage(webhookURL string, message goteamsnotify.MessageCard) error {

	// init the client
	mstClient, err := goteamsnotify.NewClient()
	if err != nil {
		return err
	}

	// attempt to send message, return the pass/fail result to caller
	return mstClient.Send(webhookURL, message)
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

// validateWebhookURLPrefix ensure that known/expected prefixes are used with
// provided webhook URL
func validateWebhookURLPrefix(webhookURL string) error {

	// TODO: Inquire about merging this upstream
	// Reasons:
	//
	// Move urls to constants for easier, less error-prone references
	// User-friendly error messages
	//
	switch {
	case strings.HasPrefix(webhookURL, webhookURLOfficecomPrefix):
	case strings.HasPrefix(webhookURL, webhookURLOffice365Prefix):
	default:
		u, err := url.Parse(webhookURL)
		if err != nil {
			return fmt.Errorf(
				"unable to parse webhook URL %q: %v",
				webhookURL,
				err,
			)
		}
		userProvidedWebhookURLPrefix := u.Scheme + "://" + u.Host

		return fmt.Errorf(
			"webhook URL does not contain expected prefix; got %q, expected one of %q or %q",
			userProvidedWebhookURLPrefix,
			webhookURLOfficecomPrefix,
			webhookURLOffice365Prefix,
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

	if err := validateWebhookURLPrefix(webhookURL); err != nil {
		return err
	}

	if err := validateWebhookURLRegex(webhookURL); err != nil {
		return err
	}

	// Indicate that we didn't spot any problems
	return nil

}
