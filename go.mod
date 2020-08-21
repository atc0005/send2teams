// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

module github.com/atc0005/send2teams

// replace github.com/atc0005/go-teams-notify => ../go-teams-notify

require (
	//gopkg.in/dasrick/go-teams-notify.v1 v1.2.0

	// temporarily use our fork while developing changes for potential
	// inclusion in the upstream project
	//
	// temporarily use local copy instead of pinning to a specific commit in
	// our test branch
	//github.com/atc0005/go-teams-notify v0.0.0
	github.com/atc0005/go-teams-notify v1.3.1-0.20200419155834-55cca556e726
	github.com/davecgh/go-spew v1.1.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 //indirect
)

go 1.14
