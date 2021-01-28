// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-teams-notify
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

/*

Package goteamsnotify is used to send messages to Microsoft Teams (channels)

PROJECT HOME

See our GitHub repo (https://github.com/atc0005/go-teams-notify) for the
latest code, to file an issue or submit improvements for review and potential
inclusion into the project.


PURPOSE

Send messages to a Microsoft Teams channel.


FEATURES

• Generate messages with one or more sections, Facts (key/value pairs) or images (hosted externally)

• Submit messages to Microsoft Teams


EXAMPLE

	import (
	"github.com/atc0005/go-teams-notify/v2"
	)

	func main() {
		_ = sendTheMessage()
	}

	func sendTheMessage() error {
		// init the client
		mstClient := goteamsnotify.NewClient()

		// setup webhook url
		webhookUrl := "https://outlook.office.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"

		// setup message card
		msgCard := goteamsnotify.NewMessageCard()
		msgCard.Title = "Hello world"
		msgCard.Text = "Here are some examples of formatted stuff like "+
			"<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
		msgCard.ThemeColor = "#DF813D"

		// send
		return mstClient.Send(webhookUrl, msgCard)
	}

*/
package goteamsnotify
