// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/atc0005/send2teams/internal/webhookurl"
)

// TeamsSubmissionTimeout is the timeout value for sending messages to
// Microsoft Teams.
func (c Config) TeamsSubmissionTimeout() time.Duration {
	delay := time.Duration(c.Retries) * time.Duration(c.RetriesDelay)

	// Fallback to 1 if retry behavior was disabled or retry delay is set too
	// short.
	if delay <= 0 {
		delay = time.Duration(1)
	}

	return delay * teamsSubmissionTimeoutMultiplier
}

// UserAgent returns a string usable as-is as a custom user agent for plugins
// provided by this project.
func (c Config) UserAgent() string {

	// Default User Agent: (Go-http-client/1.1)
	// https://datatracker.ietf.org/doc/html/draft-ietf-httpbis-p2-semantics-22#section-5.5.3
	return fmt.Sprintf(
		"%s/%s",
		c.App.Name,
		c.App.Version,
	)

}

// WebhookURL attempts to transparently process the given input for the target
// Microsoft Teams webhook URL as:
//
//   - a single base64 string
//   - multiple base64 strings ("segments") separated by commas
//   - an unencoded webhook URL
//
// If a decode attempt is successful, the decoded value is used for message
// delivery. If unsuccessful the original input value is provided as-is.
func (c Config) WebhookURL() string {
	webhookURL, err := webhookurl.DecodeBase64(c.webhookURL)
	if err != nil {
		// If base64 decoding fails return the original value as-is.
		return c.webhookURL
	}

	return strings.TrimSpace(string(webhookURL))
}
