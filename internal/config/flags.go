package config

import "flag"

// handleFlagsConfig wraps flag setup code into a bundle for potential ease of
// use and future testability
func (c *Config) handleFlagsConfig() {

	flag.BoolVar(&c.VerboseOutput, "verbose", defaultVerboseOutput, verboseOutputFlagHelp)
	flag.BoolVar(&c.SilentOutput, "silent", defaultSilentOutput, silentOutputFlagHelp)
	flag.BoolVar(&c.ConvertEOL, "convert-eol", defaultConvertEOL, convertEOLFlagHelp)
	flag.StringVar(&c.Team, "team", defaultTeamName, teamNameFlagHelp)
	flag.StringVar(&c.Channel, "channel", defaultChannelName, channelNameFlagHelp)
	flag.StringVar(&c.WebhookURL, "url", defaultWebhookURL, webhookURLFlagHelp)
	flag.StringVar(&c.ThemeColor, "color", defaultMessageThemeColor, themeColorFlagHelp)
	flag.StringVar(&c.MessageTitle, "title", defaultMessageTitle, titleFlagHelp)
	flag.StringVar(&c.MessageText, "message", defaultMessageText, messageFlagHelp)
	flag.IntVar(&c.Retries, "retries", defaultRetries, retriesFlagHelp)
	flag.IntVar(&c.RetriesDelay, "retries-delay", defaultRetriesDelay, retriesDelayFlagHelp)
	flag.BoolVar(&c.ShowVersion, "version", defaultDisplayVersionAndExit, versionFlagHelp)
	flag.BoolVar(&c.ShowVersion, "v", defaultDisplayVersionAndExit, versionFlagHelp+" (shorthand)")

	flag.Usage = flagsUsage()

	// parse flag definitions from the argument list
	flag.Parse()

}
