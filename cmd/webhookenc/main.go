// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/atc0005/send2teams/internal/webhookurl"
)

// validateResults asserts that:
//
//  1. that the raw segments when combined match the original URL
//  2. the decoded URL matches the original URL
//
// otherwise the application is aborted.
func validateResults(encodedURL string, rawSegments []string, originalURL string) {
	if strings.Join(rawSegments, "") != originalURL {
		panic("Error: parsed/reconstructed URL does not match original URL")
	}

	decodedURL, err := webhookurl.DecodeBase64(encodedURL)
	if err != nil {
		panic(err)
	}

	if strings.TrimSpace(string(decodedURL)) != strings.TrimSpace(originalURL) {
		fmt.Println("Original URL:", originalURL)
		fmt.Println("Decoded URL:", string(decodedURL))

		panic("Error: Failed to decode base64 encoded URL back to original URL")
	}
}

func main() {
	exampleURL := `https://defaultccb6deedbd294b388979d72780f62d.3b.environment.api.powerplatform.com:443/powerautomate/automations/direct/workflows/1d3ada0d8a334289b6bd8bfa6ee63bb0/triggers/manual/paths/invoke?api-version=1&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`

	if len(os.Args) < 2 || strings.TrimSpace(os.Args[1]) == "" {
		appBasename := filepath.Base(os.Args[0])

		fmt.Println("Error: Please provide input webhook URL for encoding.")
		fmt.Printf("\nExample:\n\n")
		fmt.Printf("%s '%s'\n", appBasename, exampleURL)
		return
	}

	rawURL := strings.TrimSpace(os.Args[1])

	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		panic(err)
	}

	part1 := fmt.Sprintf(
		`%v://%v`,
		u.Scheme,
		u.Host,
	)

	part2 := fmt.Sprintf(
		`%v?`,
		u.Path,
	)

	part3 := u.RawQuery

	verboseOutputTemplate := `
    _powerautomateworkflowurl_part1of3
    %v,
    %d base64 encoded characters

    _powerautomateworkflowurl_part2of3
    %v,
    %d base64 encoded characters

    _powerautomateworkflowurl_part3of3
    %v
    %d base64 encoded characters

`

	simpleOutputTemplate := `
    _powerautomateworkflowurl_part1of3    %v
    _powerautomateworkflowurl_part2of3    %v
    _powerautomateworkflowurl_part3of3    %v

`

	// encode values
	part1Encoded := webhookurl.EncodeToBase64String([]byte(part1))
	part2Encoded := webhookurl.EncodeToBase64String([]byte(part2))
	part3Encoded := webhookurl.EncodeToBase64String([]byte(part3))

	combinedBase64, err := webhookurl.JoinBase64Segments(part1Encoded, part2Encoded, part3Encoded)
	if err != nil {
		panic(err)
	}

	// Hard stop if new URL doesn't decode back to the original URL.
	validateResults(combinedBase64, []string{part1, part2, part3}, rawURL)

	fmt.Printf("Provided URL:\n%v\n\nBreaks down to these Custom Object Variables:\n", rawURL)
	fmt.Printf(
		verboseOutputTemplate,
		part1Encoded,
		utf8.RuneCountInString(part1Encoded),
		part2Encoded,
		utf8.RuneCountInString(part2Encoded),
		part3Encoded,
		utf8.RuneCountInString(part3Encoded),
	)

	fmt.Println("Copy/paste into Nagios contact entry config:")
	fmt.Printf(
		simpleOutputTemplate,
		part1Encoded,
		part2Encoded,
		part3Encoded,
	)

	fmt.Println("NOTE: We split into multiple base64 encoded values to comply with Nagios XI DB field limitations for Custom Object Variables.")

	fmt.Printf(
		"\nCombined (comma separated) input string for testing with send2teams:\n\n'%s,%s,%s'\n",
		part1Encoded, part2Encoded, part3Encoded,
	)

	fmt.Printf(
		`
Use like so:

  ./send2teams \
    --silent \
    --channel "Alerts" \
    --team "Support" \
    --message "System XYZ is down!" \
    --title "System outage alert" \
    --sender "Nagios" \
    --url "%s,%s,%s"
`,
		part1Encoded, part2Encoded, part3Encoded,
	)

	fmt.Printf(
		`
Alternatively, if you are not storing the encoded webhook URL in a Nagios database field you can also use a single base64 string like so:

  ./send2teams \
    --silent \
    --channel "Alerts" \
    --team "Support" \
    --message "System XYZ is down!" \
    --title "System outage alert" \
    --sender "Nagios" \
    --url "%s"
`,
		combinedBase64,
	)
}
