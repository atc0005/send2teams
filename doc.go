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

USAGE

Help output is below. See the README for examples.

	Usage of T:\github\send2teams\send2teams.exe:
	-channel string
			The target channel where we will send a message
	-color string
			The hex color code used to set the desired trim color on submitted messages
	-message string
			The (optionally) Markdown-formatted message to submit
	-team string
			The name of the Team containing our target channel
	-title string
			The title for the message to submit
	-url string
			The Webhook URL provided by a preconfigured Connector

*/
package main
