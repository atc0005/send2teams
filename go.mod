// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

module github.com/atc0005/send2teams

// TODO: Review this after atc0005/go-teams-notify#103 and any follow-up PRs are merged.
replace github.com/atc0005/go-teams-notify/v2 => github.com/nmaupu/go-teams-notify/v2 v2.4.3-0.20210426083400-b9e05e82de91

require (
	github.com/atc0005/go-teams-notify/v2 v2.5.0
	github.com/davecgh/go-spew v1.1.1 // indirect
)

go 1.14
