// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"

	//goteamsnotify "gopkg.in/dasrick/go-teams-notify.v1"

	// temporarily use our fork while developing changes for potential
	// inclusion in the upstream project
	goteamsnotify "github.com/atc0005/go-teams-notify"
	"github.com/atc0005/send2teams/internal/config"
	"github.com/atc0005/send2teams/teams"
)

func main() {

	// Configure our logger to use more verbose, specific format to
	// differentiate between loggers from other imported packages
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[send2teams] ")

	// Toggle library debug logging output
	// goteamsnotify.EnableLogging()
	// teams.DisableLogging()

	goteamsnotify.DisableLogging()
	teams.DisableLogging()

	//log.Debug("Initializing application")

	cfg, err := config.NewConfig()
	switch {
	// TODO: How else to guard against nil cfg object?
	case cfg != nil && cfg.ShowVersion:
		config.Branding()
		os.Exit(0)
	case err == nil:
		// do nothing for this one
	case errors.Is(err, flag.ErrHelp):
		os.Exit(0)
	default:
		log.Printf("failed to initialize application: %s", err)
		os.Exit(1)
	}

	if cfg.VerboseOutput {
		log.Printf("Configuration: %s\n", cfg)
	}

	// Convert EOL if user requested it (useful for converting script output)
	if cfg.ConvertEOL {
		cfg.MessageText = teams.ConvertEOLToBreak(cfg.MessageText)
	}

	// Setup base message card
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = cfg.MessageTitle
	msgCard.Text = cfg.MessageText
	msgCard.ThemeColor = cfg.ThemeColor

	// Create branding trailer section
	trailerSection := goteamsnotify.NewMessageCardSection()
	trailerSection.Text = config.MessageTrailer()
	trailerSection.StartGroup = true

	// Add branding trailer section, bail if unexpected error occurs
	if err := msgCard.AddSection(trailerSection); err != nil {
		log.Println("error encountered when adding section value:", err)
		os.Exit(1)
	}

	ctxSubmissionTimeout, cancel := context.WithTimeout(context.Background(), config.TeamsSubmissionTimeout)
	defer cancel()

	// Submit message card, retry submission if needed up to specified number
	// of retry attempts or until context expires.
	if err := teams.SendMessage(ctxSubmissionTimeout, cfg.WebhookURL, msgCard, cfg.Retries, cfg.RetriesDelay); err != nil {

		// Display error output if silence is not requested
		if !cfg.SilentOutput {
			log.Printf("\n\nERROR: Failed to submit message to %q channel in the %q team: %v\n\n",
				cfg.Channel, cfg.Team, err)

			if cfg.VerboseOutput {
				log.Printf("[Config]: %+v\n[Error]: %v", cfg, err)
			}

		}

		// Regardless of silent flag, explicitly note unsuccessful results
		os.Exit(1)
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
