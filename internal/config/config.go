// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

const (
	versionFlagHelp       = "Whether to display application version and then immediately exit application."
	verboseOutputFlagHelp = "Whether detailed output should be shown after message submission success or failure."
	silentOutputFlagHelp  = "Whether ANY output should be shown after message submission success or failure."
	convertEOLFlagHelp    = "Whether messages with Windows, Mac and Linux newlines are updated to use break statements before message submission."
	teamNameFlagHelp      = "The name of the Team containing our target channel. Used in log messages. If not specified, defaults to \"unspecified\"."
	channelNameFlagHelp   = "The target channel where we will send a message. Used in log messages. If not specified, defaults to \"unspecified\"."
	webhookURLFlagHelp    = "The Webhook URL provided by a preconfigured Connector."
	themeColorFlagHelp    = "The hex color code used to set the desired trim color on submitted messages."
	titleFlagHelp         = "The title for the message to submit."
	messageFlagHelp       = "The message to submit. This message may be provided in Markdown format."
	senderFlagHelp        = "The (optional) sending application name or generator of the message this app will attempt to deliver."
	retriesFlagHelp       = "The number of attempts that this application will make to deliver messages before giving up."
	retriesDelayFlagHelp  = "The number of seconds that this application will wait before making another delivery attempt."
)

// Default flag settings if not overridden by user input
const (
	defaultMessageThemeColor     string = "#832561"
	defaultSilentOutput          bool   = false
	defaultVerboseOutput         bool   = false
	defaultConvertEOL            bool   = false
	defaultTeamName              string = "unspecified"
	defaultChannelName           string = "unspecified"
	defaultWebhookURL            string = ""
	defaultMessageTitle          string = ""
	defaultMessageText           string = ""
	defaultSender                string = ""
	defaultDisplayVersionAndExit bool   = false
	defaultRetries               int    = 2
	defaultRetriesDelay          int    = 2
)

// Overridden via Makefile for release builds
var version string = "dev build"

// Primarily used with branding
const myAppName string = "send2teams"
const myAppURL string = "https://github.com/atc0005/" + myAppName

// teamsSubmissionTimeoutMultiplier is the timeout value for sending messages
// to Microsoft Teams. This value is used along with user specified (or
// default) retries and retries delay values to calculate a context with the
// desired timeout value.
const teamsSubmissionTimeoutMultiplier time.Duration = 2 * time.Second

// DefaultNagiosNotificationTimeout is the default timeout value for Nagios 3
// and 4 installations. This is our *default* timeout ceiling.
const DefaultNagiosNotificationTimeout time.Duration = 30 * time.Second

// brandingTextPrefix is the lead-in or prefix text used to brand or give
// credit to this application for messages delivered to a Microsoft Teams
// channel.
const brandingTextPrefix string = "Message delivered by"

// brandingTextSuffix is the lead-out or suffix text used to give credit to
// the application responsible for generating the text or messages that this
// one will attempt to deliver to a Microsoft Teams channel.
const brandingTextSuffix string = "on behalf of"

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags provided by the user.
type Config struct {

	// Team is the human-readable name of the Microsoft Teams "team" that
	// contains the channel we wish to post a message to. This is used in
	// informational output produced by this application only; the remote API
	// does not receive this value.
	Team string

	// Channel is human-readable name of the channel within a specific
	// Microsoft Teams "team". This is used in informational output produced
	// by this application only; the remote API does not receive this value.
	Channel string

	// WebhookURL is the full URL used to submit messages to the Teams channel
	// This URL is in the form of https://outlook.office.com/webhook/xxx or
	// https://outlook.office365.com/webhook/xxx. This URL is REQUIRED in
	// order for this application to function and needs to be created in
	// advance by adding/configuring a Webhook Connector in a Microsoft Teams
	// channel that you wish to submit messages to using this application.
	WebhookURL string

	// ThemeColor is a hex color code string representing the desired border
	// trim color for our submitted messages.
	ThemeColor string

	// MessageTitle is the text shown on the top portion of the message "card"
	// that is displayed in Microsoft Teams for the message that we send.
	MessageTitle string

	// MessageText is an (optionally) Markdown-formatted string representing
	// the message that we will submit.
	MessageText string

	// Sender is an optional value provided to indicate what application was
	// responsible for generating the message that this one will attempt to
	// deliver.
	Sender string

	// Retries is the number of attempts that this application will make
	// to deliver messages before giving up.
	Retries int

	// RetriesDelay is the number of seconds to wait between retry attempts.
	RetriesDelay int

	// Whether detailed output should be shown after message submission
	// success or failure.
	VerboseOutput bool

	// Whether ANY output should be shown after message submission success or
	// failure.
	SilentOutput bool

	// Whether messages with Windows, Mac and Linux newlines are updated to
	// use break statements before message submission.
	ConvertEOL bool

	// ShowVersion is a flag indicating whether the user opted to display only
	// the version string and then immediately exit the application
	ShowVersion bool
}

// Branding is responsible for emitting application name, version and origin
func Branding() {
	fmt.Fprintf(flag.CommandLine.Output(), "\n%s %s\n%s\n\n", myAppName, version, myAppURL)
}

// MessageTrailer generates a branded "footer" for use with submitted Teams
// messages. If specified, the sending or "generator" application is credited
// as the source of the message, while this application is credited as the
// delivery agent/mechanism.
func MessageTrailer(sender string) string {
	var onBehalfOf string
	if strings.TrimSpace(sender) != "" {
		onBehalfOf = fmt.Sprintf(" %s %s ", brandingTextSuffix, sender)
	}

	return fmt.Sprintf(
		"%s [%s](%s) (%s) at %s%s",
		brandingTextPrefix,
		myAppName,
		myAppURL,
		version,
		time.Now().Format(time.RFC3339),
		onBehalfOf,
	)
}

// flagsUsage displays branding information and general usage details
func flagsUsage() func() {

	return func() {

		myBinaryName := filepath.Base(os.Args[0])

		Branding()

		fmt.Fprintf(flag.CommandLine.Output(), "Usage of \"%s\":\n",
			myBinaryName,
		)
		flag.PrintDefaults()

	}
}

func (c Config) String() string {
	return fmt.Sprintf(
		"Team=%q, "+
			"Channel=%q, "+
			"WebhookURL=%q, "+
			"ThemeColor=%q, "+
			"MessageTitle=%q, "+
			"MessageText=%q, "+
			"Sender=%q, "+
			"Retries=%q, "+
			"RetriesDelay=%q, "+
			"AppTimeout=%q",
		c.Team,
		c.Channel,
		c.WebhookURL,
		c.ThemeColor,
		c.MessageTitle,
		c.MessageText,
		c.Sender,
		strconv.Itoa(c.Retries),
		strconv.Itoa(c.RetriesDelay),
		c.TeamsSubmissionTimeout(),
	)
}

// NewConfig is a factory function that produces a new Config object based
// on user provided flag values.
func NewConfig() (*Config, error) {
	cfg := Config{}

	cfg.handleFlagsConfig()

	// Return immediately if user just wants version details
	if cfg.ShowVersion {
		return &cfg, nil
	}

	// log.Debug("Validating configuration ...")
	if err := cfg.Validate(); err != nil {
		flag.Usage()
		return nil, err
	}
	// log.Debug("Configuration validated")

	return &cfg, nil
}

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	if c.SilentOutput && c.VerboseOutput {
		return fmt.Errorf("unsupported: You cannot have both silent and verbose output")
	}

	// Expected pattern: #832561
	if len(c.ThemeColor) < len(defaultMessageThemeColor) {

		expectedLength := len(defaultMessageThemeColor)
		actualLength := len(c.ThemeColor)
		return fmt.Errorf("provided message theme color too short; got message %q of length %d, expected length of %d",
			c.ThemeColor, actualLength, expectedLength)
	}

	// Note: This is separate from goteamsnotify.IsValidMessageCard() That
	// function specifically checks the results of creating and fleshing out a
	// MessageCard value, this validation check is more concerned with the
	// specific value supplied via flag input.
	if c.MessageTitle == "" {
		return fmt.Errorf("message title too short")
	}

	// Note: This is separate from goteamsnotify.IsValidMessageCard() That
	// function specifically checks the results of creating and fleshing out a
	// MessageCard value, this validation check is more concerned with the
	// specific value supplied via flag input.
	if c.MessageText == "" {
		return fmt.Errorf("message content too short")
	}

	// Team and Channel names are optional. If provided, use as-is.

	// Sender is optional. If provided, use as-is.

	if c.Retries < 0 {
		return fmt.Errorf("retries too short")
	}

	if c.RetriesDelay < 0 {
		return fmt.Errorf("retries delay too short")
	}

	// Create Microsoft Teams client
	mstClient := goteamsnotify.NewClient()

	if err := mstClient.ValidateWebhook(c.WebhookURL); err != nil {
		return fmt.Errorf("webhook URL validation failed: %w", err)
	}

	// Indicate that we didn't spot any problems
	return nil

}
