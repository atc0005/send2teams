// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import "time"

// TeamsSubmissionTimeout is the timeout value for sending messages to
// Microsoft Teams.
func (c Config) TeamsSubmissionTimeout() time.Duration {

	return time.Duration(c.Retries) *
		time.Duration(c.RetriesDelay) *
		teamsSubmissionTimeoutMultiplier
}
