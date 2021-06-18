// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

/*

send2teams is a small CLI tool used to submit messages to Microsoft Teams.

PROJECT HOME

See our GitHub repo (https://github.com/atc0005/send2teams) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

PURPOSE

send2teams is intended for use by Nagios, scripts or other actions that may
need to submit pass/fail results to a MS Teams channel.

FEATURES

• single binary, no outside dependencies

• minimal configuration

• very few build dependencies

• optional conversion of messages with Windows, Mac or Linux newlines to `<br>` to increase compatibility with Teams formatting

• message delivery retry support with retry and retry delay values configurable via flag

• optional support for specifying url/description pairs for display in Microsoft Teams messages

USAGE

See our main README for supported settings and examples.

*/
package main
