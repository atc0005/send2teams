// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

module github.com/atc0005/send2teams

go 1.23.0

// Allow for testing local changes before they're published.
//
// replace github.com/atc0005/go-teams-notify/v2 => ../go-teams-notify

require github.com/atc0005/go-teams-notify/v2 v2.13.0
