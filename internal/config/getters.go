// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"time"
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
