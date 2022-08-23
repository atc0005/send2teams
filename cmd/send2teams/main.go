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
	"fmt"
	"log"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/adaptivecard"
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

	// This should only trigger if user specifies large retry values.
	if cfg.TeamsSubmissionTimeout() > config.DefaultNagiosNotificationTimeout {
		if !cfg.SilentOutput {
			log.Printf(
				"WARNING: app cancellation timeout value of %v greater than default Nagios command timeout value!",
				cfg.TeamsSubmissionTimeout(),
			)
		}
	}

	ctxSubmissionTimeout, cancel := context.WithTimeout(context.Background(), cfg.TeamsSubmissionTimeout())
	defer cancel()

	// Create Microsoft Teams client
	mstClient := goteamsnotify.NewTeamsClient()

	// Override User Agent.
	mstClient.SetUserAgent(cfg.UserAgent())

	// Disable webhook URL validation if requested by user.
	mstClient.SkipWebhookURLValidationOnSend(cfg.DisableWebhookURLValidation)

	// Convert EOL (useful for output from scripts) in the incoming text if
	// user requested it.
	if cfg.ConvertEOL {
		cfg.MessageText = adaptivecard.ConvertEOL(cfg.MessageText)

		// Not 100% safe to apply across the board.
		//
		// It is unlikely, but not impossible that someone would submit raw
		// text with break statements. When you consider that the flag is
		// named "convert-eol", it is entirely reasonable that the user would
		// expect break statements to remain untouched.
		//
		// cfg.MessageText = adaptivecard.ConvertBreakToEOL(cfg.MessageText)
	}

	card, err := adaptivecard.NewTextBlockCard(cfg.MessageText, cfg.MessageTitle, true)
	if err != nil {
		if !cfg.SilentOutput {
			log.Printf(
				"\n\nERROR: Failed to create new card using specified text/title values for %q channel in the %q team: %v\n\n",
				cfg.Channel,
				cfg.Team,
				err,
			)
		}
		// Regardless of silent flag, explicitly note unsuccessful results
		appExitCode = 1
		return
	}
	card.SetFullWidth()

	if len(cfg.UserMentions) > 0 {
		// Process user mention details specified by user, create user mention
		// values that we can attach to the card.
		userMentions := make([]adaptivecard.Mention, 0, len(cfg.UserMentions))
		for _, mention := range cfg.UserMentions {
			userMention, err := adaptivecard.NewMention(mention.Name, mention.ID)
			if err != nil {
				if !cfg.SilentOutput {
					log.Printf("\n\nERROR: Failed to process user mention for %q channel in the %q team: %v\n\n",
						cfg.Channel, cfg.Team, err)
				}
				// Regardless of silent flag, explicitly note unsuccessful results
				appExitCode = 1
				return
			}
			userMentions = append(userMentions, userMention)
		}

		// Add user mention collection to card.
		if err := card.AddMention(true, userMentions...); err != nil {
			if !cfg.SilentOutput {
				log.Printf("\n\nERROR: Failed to add user mentions to message for %q channel in the %q team: %v\n\n",
					cfg.Channel, cfg.Team, err)
			}
			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}
	}

	// If provided, use target URLs and their descriptions to add labelled
	// URL "buttons" to Microsoft Teams message.
	if len(cfg.TargetURLs) > 0 {

		// Create dedicated container for all action items.
		actionsContainer := adaptivecard.NewContainer()
		actionsContainer.Separator = false
		actionsContainer.Style = adaptivecard.ContainerStyleEmphasis
		actionsContainer.Spacing = adaptivecard.SpacingExtraLarge

		actions := make([]adaptivecard.Action, 0, len(cfg.TargetURLs))

		for i := range cfg.TargetURLs {

			urlAction, err := adaptivecard.NewActionOpenURL(
				cfg.TargetURLs[i].URL.String(),
				cfg.TargetURLs[i].Description,
			)
			if err != nil {
				if !cfg.SilentOutput {
					log.Printf(
						"\n\nERROR: Failed to process openURL action for %q channel in the %q team: %v\n\n",
						cfg.Channel,
						cfg.Team,
						err,
					)
				}
				// Regardless of silent flag, explicitly note unsuccessful results
				appExitCode = 1
				return
			}
			actions = append(actions, urlAction)
		}

		if err := actionsContainer.AddAction(true, actions...); err != nil {
			if !cfg.SilentOutput {
				log.Printf(
					"\n\nERROR: Failed to add openURL action to container for %q channel in the %q team: %v\n\n",
					cfg.Channel,
					cfg.Team,
					err,
				)
			}
			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}

		if err := card.AddContainer(false, actionsContainer); err != nil {
			if !cfg.SilentOutput {
				log.Printf("\n\nERROR: Failed to add actions container to card for %q channel in the %q team: %v\n\n",
					cfg.Channel, cfg.Team, err)
			}
			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}
	}

	// If requested, skip appending the branding trailer to messages.
	if !cfg.DisableBrandingTrailer {

		// Process branding trailer content.
		//
		// NOTE: Unlike MessageCard text which has benefited from \r\n
		// (windows), \r (mac) and \n (unix) conversion to <br> statements in
		// the past, <br> statements in Adaptive Card text remain as-is in the
		// final rendered message. This is not useful.
		trailerText := fmt.Sprintf(
			"\n\n%s",
			config.MessageTrailer(cfg.Sender),
		)

		trailerContainer := adaptivecard.NewContainer()
		trailerContainer.Separator = true
		trailerContainer.Spacing = adaptivecard.SpacingExtraLarge

		trailerTextBlock := adaptivecard.NewTextBlock(trailerText, true)
		trailerTextBlock.Size = adaptivecard.SizeSmall
		trailerTextBlock.Weight = adaptivecard.WeightLighter

		if err := trailerContainer.AddElement(false, trailerTextBlock); err != nil {
			if !cfg.SilentOutput {
				log.Printf("\n\nERROR: Failed to add text block to trailer container for card for %q channel in the %q team: %v\n\n",
					cfg.Channel, cfg.Team, err)
			}
			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}
		if err := card.AddContainer(false, trailerContainer); err != nil {
			if !cfg.SilentOutput {
				log.Printf("\n\nERROR: Failed to add trailer container to card for %q channel in the %q team: %v\n\n",
					cfg.Channel, cfg.Team, err)
			}
			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}
	}

	message, err := adaptivecard.NewMessageFromCard(card)
	if err != nil {
		if !cfg.SilentOutput {
			log.Printf(
				"\n\nERROR: Failed to create new message from card for %q channel in the %q team: %v\n\n",
				cfg.Channel,
				cfg.Team,
				err,
			)

			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}
	}

	if cfg.VerboseOutput {
		if err := message.Prepare(); err != nil {
			log.Printf("\n\nERROR: Failed to prepare message for %q channel in the %q team: %v\n\n",
				cfg.Channel, cfg.Team, err)

			// Regardless of silent flag, explicitly note unsuccessful results
			appExitCode = 1
			return
		}

		log.Println(message.PrettyPrint())
	}

	// Submit message card using Microsoft Teams client, retry submission if
	// needed up to specified number of retry attempts.
	sendErr := mstClient.SendWithRetry(ctxSubmissionTimeout, cfg.WebhookURL, message, cfg.Retries, cfg.RetriesDelay)

	switch {

	case cfg.IgnoreInvalidResponse &&
		errors.Is(sendErr, goteamsnotify.ErrInvalidWebhookURLResponseText):

		if !cfg.SilentOutput {
			log.Printf(
				"WARNING: invalid response received from %q endpoint", cfg.WebhookURL)
			log.Printf("ignoring error response as requested: \n%s", sendErr)
		}

	// If an error occurred and we were not expecting one.
	case sendErr != nil:
		// Display error output if silence is not requested
		if !cfg.SilentOutput {
			log.Printf("\n\nERROR: Failed to submit message to %q channel in the %q team: %v\n\n",
				cfg.Channel, cfg.Team, sendErr)

			if cfg.VerboseOutput {
				log.Printf("[Config]: %+v\n[Error]: %v", cfg, sendErr)
			}

		}

		// Regardless of silent flag, explicitly note unsuccessful results
		appExitCode = 1
		return

	default:
		if !cfg.SilentOutput {
			// Emit basic success message
			log.Println("Message successfully sent!")
		}

	}

	if cfg.VerboseOutput {
		log.Printf("Configuration used: %#v\n", cfg)
		log.Printf("Webhook URL: %s\n", cfg.WebhookURL)
		log.Printf("Message values sent: %#v\n", message)
	}

}
