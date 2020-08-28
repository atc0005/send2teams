// Copyright 2020 Enrico Hoffmann
// Copyright 2020 Adam Chalkley
//
// https:#github.com/atc0005/go-teams-notify
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package goteamsnotify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// logger is a package logger that can be enabled from client code to allow
// logging output from this package when desired/needed for troubleshooting
var logger *log.Logger

// Known webhook URL prefixes for submitting messages to Microsoft Teams
const (
	WebhookURLOfficecomPrefix = "https://outlook.office.com"
	WebhookURLOffice365Prefix = "https://outlook.office365.com"
)

// DefaultWebhookSendTimeout specifies how long the message operation may take
// before it times out and is cancelled.
const DefaultWebhookSendTimeout = 5 * time.Second

// API - interface of MS Teams notify
type API interface {
	Send(webhookURL string, webhookMessage MessageCard) error
	SendWithContext(ctx context.Context, webhookURL string, webhookMessage MessageCard) error
}

type teamsClient struct {
	httpClient *http.Client
}

func init() {
	// Disable logging output by default unless client code explicitly
	// requests it
	logger = log.New(os.Stderr, "[goteamsnotify] ", 0)
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

// NewClient - create a brand new client for MS Teams notify
func NewClient() API {
	client := teamsClient{
		httpClient: &http.Client{
			// We're using a context instead of setting this directly
			// Timeout: DefaultWebhookSendTimeout,
		},
	}
	return &client
}

// Send is a wrapper function around the SendWithContext method in order to
// provide backwards compatibility.
func (c teamsClient) Send(webhookURL string, webhookMessage MessageCard) error {
	// Create context that can be used to emulate existing timeout behavior.
	ctx, cancel := context.WithTimeout(context.Background(), DefaultWebhookSendTimeout)
	defer cancel()

	return c.SendWithContext(ctx, webhookURL, webhookMessage)
}

// SendWithContext posts a notification to the provided MS Teams webhook URL.
// The http client request honors the cancellation or timeout of the provided
// context.
func (c teamsClient) SendWithContext(ctx context.Context, webhookURL string, webhookMessage MessageCard) error {
	logger.Printf("SendWithContext: Webhook message received: %#v\n", webhookMessage)

	// Validate input data
	if valid, err := IsValidInput(webhookMessage, webhookURL); !valid {
		return err
	}

	// prepare message
	webhookMessageByte, _ := json.Marshal(webhookMessage)
	webhookMessageBuffer := bytes.NewBuffer(webhookMessageByte)

	// Basic, unformatted JSON
	//logger.Printf("SendWithContext: %+v\n", string(webhookMessageByte))

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, webhookMessageByte, "", "\t"); err != nil {
		return err
	}
	logger.Printf("SendWithContext: Payload for Microsoft Teams: \n\n%v\n\n", prettyJSON.String())

	// prepare request (error not possible)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, webhookURL, webhookMessageBuffer)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	// do the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		logger.Println(err)
		return err
	}

	if ctx.Err() != nil {
		logger.Println("SendWithContext: Context has expired after Do(req):", time.Now().Format("15:04:05"))
	}

	// Make sure that we close the response body once we're done with it
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	// Get the response body, then convert to string for use with extended
	// error messages
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Println(err)
		return err
	}
	responseString := string(responseData)

	if res.StatusCode >= 299 {
		// 400 Bad Response is likely an indicator that we failed to provide a
		// required field in our JSON payload. For example, when leaving out
		// the top level MessageCard Summary or Text field, the remote API
		// returns "Summary or Text is required." as a text string. We include
		// that response text in the error message that we return to the
		// caller.

		err = fmt.Errorf("error on notification: %v, %q", res.Status, responseString)
		logger.Println(err)
		return err
	}

	// log the response string
	logger.Printf("SendWithContext: Response string from Microsoft Teams API: %v\n", responseString)

	return nil
}

// helper --------------------------------------------------------------------------------------------------------------

// IsValidInput is a validation "wrapper" function. This function is intended
// to run current validation checks and offer easy extensibility for future
// validation requirements.
func IsValidInput(webhookMessage MessageCard, webhookURL string) (bool, error) {
	// validate url
	if valid, err := IsValidWebhookURL(webhookURL); !valid {
		return false, err
	}

	// validate message
	if valid, err := IsValidMessageCard(webhookMessage); !valid {
		return false, err
	}

	return true, nil
}

// IsValidWebhookURL performs validation checks on the webhook URL used to
// submit messages to Microsoft Teams.
func IsValidWebhookURL(webhookURL string) (bool, error) {
	switch {
	case strings.HasPrefix(webhookURL, WebhookURLOfficecomPrefix):
	case strings.HasPrefix(webhookURL, WebhookURLOffice365Prefix):
	default:
		u, err := url.Parse(webhookURL)
		if err != nil {
			return false, fmt.Errorf(
				"unable to parse webhook URL %q: %w",
				webhookURL,
				err,
			)
		}
		userProvidedWebhookURLPrefix := u.Scheme + "://" + u.Host

		return false, fmt.Errorf(
			"webhook URL does not contain expected prefix; got %q, expected one of %q or %q",
			userProvidedWebhookURLPrefix,
			WebhookURLOfficecomPrefix,
			WebhookURLOffice365Prefix,
		)
	}

	return true, nil
}

// IsValidMessageCard performs validation/checks for known issues with
// MessardCard values.
func IsValidMessageCard(webhookMessage MessageCard) (bool, error) {
	if (webhookMessage.Text == "") && (webhookMessage.Summary == "") {
		// This scenario results in:
		// 400 Bad Request
		// Summary or Text is required.
		return false, fmt.Errorf("invalid message card: summary or text field is required")
	}

	return true, nil
}
