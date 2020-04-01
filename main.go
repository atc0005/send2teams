// Copyright 2019 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"flag"
	"log"
	"os"

	//goteamsnotify "gopkg.in/dasrick/go-teams-notify.v1"

	// temporarily use our fork while developing changes for potential
	// inclusion in the upstream project
	goteamsnotify "github.com/atc0005/go-teams-notify"
	"github.com/atc0005/send2teams/config"
	"github.com/atc0005/send2teams/teams"
)

func main() {

	// Configure our logger to use more verbose, specific format to
	// differentiate between the go-teams-notify package logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Add extra trailing whitespace to align log message prefix with package
	// log message prefix
	log.SetPrefix("[send2teams]   ")

	// Toggle library debug logging output
	goteamsnotify.EnableLogging()
	// goteamsnotify.DisableLogging()

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

	// setup message card
	// msgCard := goteamsnotify.NewMessageCard()
	// msgCard.Title = cfg.MessageTitle
	// msgCard.Text = "placeholder (top-level text content)"
	// msgCard.ThemeColor = cfg.ThemeColor

	send := func(cfg *config.Config, msgCard goteamsnotify.MessageCard) {
		// Submit message card
		if err := teams.SendMessage(cfg.WebhookURL, msgCard); err != nil {

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
	}

	// Assign test cases
	testCases := []goteamsnotify.MessageCard{
		testCase1(cfg),
		testCase2(cfg),
	}

	for _, testCase := range testCases {
		send(cfg, testCase)
	}

}
