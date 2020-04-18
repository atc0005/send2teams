// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

module github.com/atc0005/send2teams

require (
	//gopkg.in/dasrick/go-teams-notify.v1 v1.2.0

	// temporarily use our fork while developing changes for potential
	// inclusion in the upstream project
	github.com/atc0005/go-teams-notify v1.3.1-0.20200418135531-d35cfbe04599
	github.com/davecgh/go-spew v1.1.1 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

go 1.13
