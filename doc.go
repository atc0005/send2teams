// Copyright 2020 Adam Chalkley
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

• exported `teams` package to handle formatting text content as code for proper display within Microsoft Teams

•  message delivery retry support with retry and retry delay values configurable via flag


USAGE

Help output is below. See the README for examples.

	send2teams dev build
	https://github.com/atc0005/send2teams

	Usage of "send2teams.exe":
	-channel string
			The target channel where we will send a message.
	-color string
			The hex color code used to set the desired trim color on submitted messages. (default "#832561")
	-convert-eol
			Whether messages with Windows, Mac and Linux newlines are updated to use break statements before message submission.
	-message string
			The message to submit. This message may be provided in Markdown format.
	-retries int
			The number of attempts that this application will make to deliver messages before giving up. (default 2)
	-retries-delay int
			The number of seconds that this application will wait before making another delivery attempt. (default 2)
	-silent
			Whether ANY output should be shown after message submission success or failure.
	-team string
			The name of the Team containing our target channel.
	-title string
			The title for the message to submit.
	-url string
			The Webhook URL provided by a preconfigured Connector.
	-v    Whether to display application version and then immediately exit application. (shorthand)
	-verbose
			Whether detailed output should be shown after message submission success or failure.
	-version
			Whether to display application version and then immediately exit application.

*/
package main
