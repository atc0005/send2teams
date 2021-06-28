// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"context"
	"errors"
	"log"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/send2teams/internal/config"
)

func main() {

	// Configure our logger to use more verbose, specific format to
	// differentiate between loggers from other imported packages
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[send2teams] ")

	// Toggle library debug logging output
	// goteamsnotify.EnableLogging()
	goteamsnotify.DisableLogging()

	cfg, cfgErr := config.NewConfig()
	switch {
	case errors.Is(cfgErr, config.ErrVersionRequested):
		config.Branding()
		os.Exit(0)
	case cfgErr != nil:
		log.Fatalf("failed to initialize application: %s", cfgErr)
	}

	if cfg.VerboseOutput {
		log.Printf("Configuration: %s\n", cfg)
	}

	// Emulate returning exit code from main function by "queuing up" a
	// default exit code that matches expectations, but allow explicitly
	// setting the exit code in such a way that is compatible with using
	// deferred function calls throughout the application.
	var appExitCode int
	defer func(code *int) {
		var exitCode int
		if code != nil {
			exitCode = *code
		}
		os.Exit(exitCode)
	}(&appExitCode)

	// Convert EOL if user requested it (useful for converting script output)
	if cfg.ConvertEOL {
		cfg.MessageText = goteamsnotify.ConvertEOLToBreak(cfg.MessageText)
	}

	// Setup base message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = cfg.MessageTitle
	msgCard.Text = cfg.MessageText
	msgCard.ThemeColor = cfg.ThemeColor

	// If provided, use target URLs and their descriptions to add labelled URL
	// "buttons" to Microsoft Teams message.
	if cfg.TargetURLs != nil {

		// Create dedicated section for all potentialAction items
		actionSection := goteamsnotify.NewMessageCardSection()
		actionSection.StartGroup = true

		for i := range cfg.TargetURLs {

			pa := goteamsnotify.NewMessageCardPotentialAction(
				goteamsnotify.PotentialActionOpenURIType,
				cfg.TargetURLs[i].Description,
			)

			pa.MessageCardPotentialActionOpenURI.Targets =
				[]goteamsnotify.MessageCardPotentialActionOpenUriTarget{
					{
						OS:  "default",
						URI: cfg.TargetURLs[i].URL.String(),
					},
				}

			if err := actionSection.AddPotentialAction(pa); err != nil {
				log.Println("error encountered when adding target URL to message:", err)
				appExitCode = 1

				return
			}
		}

		if err := msgCard.AddSection(actionSection); err != nil {
			log.Println("error encountered when adding section value:", err)
			appExitCode = 1
			return
		}
	}

	// Create branding trailer section
	trailerSection := goteamsnotify.NewMessageCardSection()
	trailerSection.Text = config.MessageTrailer(cfg.Sender)
	trailerSection.StartGroup = true

	// Add branding trailer section, bail if unexpected error occurs
	if err := msgCard.AddSection(trailerSection); err != nil {
		log.Println("error encountered when adding section value:", err)
		appExitCode = 1
		return
	}

	// This should only trigger if user specifies large retry values.
	if cfg.TeamsSubmissionTimeout() > config.DefaultNagiosNotificationTimeout {
		log.Printf(
			"WARNING: app cancellation timeout value of %v greater than default Nagios command timeout value!",
			cfg.TeamsSubmissionTimeout(),
		)
	}

	ctxSubmissionTimeout, cancel := context.WithTimeout(context.Background(), cfg.TeamsSubmissionTimeout())
	defer cancel()

	// Create Microsoft Teams client
	mstClient := goteamsnotify.NewClient()

	// Submit message card using Microsoft Teams client, retry submission
	// if needed up to specified number of retry attempts.
	if err := mstClient.SendWithRetry(ctxSubmissionTimeout, cfg.WebhookURL, msgCard, cfg.Retries, cfg.RetriesDelay); err != nil {

		// Display error output if silence is not requested
		if !cfg.SilentOutput {
			log.Printf("\n\nERROR: Failed to submit message to %q channel in the %q team: %v\n\n",
				cfg.Channel, cfg.Team, err)

			if cfg.VerboseOutput {
				log.Printf("[Config]: %+v\n[Error]: %v", cfg, err)
			}

		}

		// Regardless of silent flag, explicitly note unsuccessful results
		appExitCode = 1
		return
	}

	if !cfg.SilentOutput {
		// Emit basic success message
		log.Println("Message successfully sent!")
	}

	if cfg.VerboseOutput {
		log.Printf("Configuration used: %#v\n", cfg)
		log.Printf("Webhook URL: %s\n", cfg.WebhookURL)
		log.Printf("MessageCard values sent: %#v\n", msgCard)
	}

}
