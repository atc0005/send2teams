<!-- omit in toc -->
# send2teams

Small CLI tool used to submit messages to Microsoft Teams.

[![Latest Release](https://img.shields.io/github/release/atc0005/send2teams.svg?style=flat-square)](https://github.com/atc0005/send2teams/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/send2teams.svg)](https://pkg.go.dev/github.com/atc0005/send2teams)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/send2teams)](https://github.com/atc0005/send2teams)
[![Lint and Build](https://github.com/atc0005/send2teams/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/send2teams/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/send2teams/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/send2teams/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [How to install it](#how-to-install-it)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
- [Configuration Options](#configuration-options)
  - [Webhook URLs](#webhook-urls)
    - [Expected format](#expected-format)
    - [How to create a webhook URL (Connector)](#how-to-create-a-webhook-url-connector)
  - [Command-line](#command-line)
- [Limitations](#limitations)
  - [message size](#message-size)
- [Examples](#examples)
  - [One-off](#one-off)
  - [Using an invalid flag](#using-an-invalid-flag)
  - [Specifying url, description pairs](#specifying-url-description-pairs)
  - [User mentions](#user-mentions)
    - [One mention](#one-mention)
    - [Multiple mentions](#multiple-mentions)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

First of all, many thanks to the developer/contributors of the original
`dasrick/go-teams-notify` package. While this project now uses a fork of that
original project, this project would likely not have been possible without the
efforts of the original developer.

This project provides:

- `send2teams`
  - Small CLI tool used to submit messages to Microsoft Teams. `send2teams` is
    intended for use by Nagios, scripts or other actions that may need to
    submit pass/fail results to a MS Teams channel.

Prior to `v0.4.7` this project also provided a `teams` subpackage. All of that
functionality has since been migrated to the `atc0005/go-teams-notify`
project. All client code for that package has been updated to use
`atc0005/go-teams-notify` in place of the previous `teams` subpackage of this
project.

## Features

- single binary, no outside dependencies
- minimal configuration
- very few build dependencies
- optional conversion of messages with Windows, Mac or Linux newlines to
  increase compatibility with Teams formatting
- message delivery retry support with retry and retry delay values
  configurable via flag
- support for user mentions
- optional support for noting a sending application as the source of the
  message
- optional support for specifying target `url`, `description` comma-separated
  pairs for use as labelled "buttons" within a Microsoft Teams message
- optional support for omitting the "branding" trailer from generated messages

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 10
- Ubuntu Linux 18.04+

## How to install it

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
   - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/send2teams`
   1. `cd send2teams`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - for current operating system
     - `go build -mod=vendor ./cmd/send2teams/`
       - *forces build to use bundled dependencies in top-level `vendor`
         folder*
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for Windows
      - `make windows`
   - for Linux
     - `make linux`
1. Copy the applicable binary to whatever systems needs to run it
   - if using `Makefile`: look in `/tmp/release_assets/send2teams/`
   - if using `go build`: look in `/tmp/send2teams/`

**NOTE**: Depending on which `Makefile` recipe you use the generated binary
may be compressed and have an `xz` extension. If so, you should decompress the
binary first before deploying it (e.g., `xz -d send2teams-linux-amd64.xz`).

### Using release binaries

1. Download the [latest
   release](https://github.com/atc0005/send2teams/releases/latest) binaries
1. Decompress binaries
   - e.g., `xz -d send2teams-linux-amd64.xz`
1. Deploy
   - Place `send2teams` in a location of your choice
     - e.g., `/usr/local/bin/send2teams`

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

## Configuration Options

### Webhook URLs

#### Expected format

Valid webhook URLs for Microsoft Teams use one of several (confirmed) FQDNs
patterns:

- `outlook.office.com`
- `outlook.office365.com`
- `*.webhook.office.com`
  - e.g., `example.webhook.office.com`

Using a webhook URL with any of these FQDN patterns appears to give identical
results.

Here are complete, equivalent example webhook URLs from Microsoft's
documentation using the FQDNs above:

- <https://outlook.office.com/webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>
- <https://outlook.office365.com/webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>
- <https://example.webhook.office.com/webhookb2/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>
  - note the `webhookb2` sub-URI specific to this FQDN pattern

All of these patterns should pass the default validation applied to
user-specified webhook URLs.

#### How to create a webhook URL (Connector)

1. Open Microsoft Teams
1. Navigate to the channel where you wish to receive incoming messages from
   this application
1. Select `⋯` next to the channel name and then choose Connectors.
1. Scroll through the list of Connectors to Incoming Webhook, and choose Add.
1. Enter a name for the webhook, upload an image to associate with data from
   the webhook, and choose Create.
1. Copy the webhook URL to the clipboard and save it. You'll need the webhook
   URL for sending information to Microsoft Teams.
   - NOTE: While you can create another easily enough, you should treat this
     webhook URL as sensitive information as anyone with this unique URL is
     able to send messages (without authentication) into the associated
     channel.
1. Choose Done.

Credit:
[docs.microsoft.com](https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/connectors-using#setting-up-a-custom-incoming-webhook),
[gist comment from
shadabacc3934](https://gist.github.com/chusiang/895f6406fbf9285c58ad0a3ace13d025#gistcomment-3562501)

### Command-line

Currently `send2teams` only supports command-line configuration flags.
Requests for other configuration sources will be considered.

| Flag                       | Required | Default       | Possible                                                  | Description                                                                                                                                       |
| -------------------------- | -------- | ------------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`                | No       | N/A           | N/A                                                       | Display Help; show available flags.                                                                                                               |
| `v`, `version`             | No       | `false`       | `true`, `false`                                           | Whether to display application version and then immediately exit application.                                                                     |
| `channel`                  | No       | `unspecified` | *valid Microsoft Teams channel name*                      | The target channel where we will send a message. If not specified, defaults to `unspecified`.                                                     |
| `color`                    | No       | `NotUsed`     | N/A                                                       | NOOP; this setting is no longer used. Values specified for this flag are ignored.                                                                 |
| `message`                  | Yes      |               | *valid message string*                                    | The (optionally) Markdown-formatted message to submit.                                                                                            |
| `team`                     | No       | `unspecified` | *valid Microsoft Teams team name*                         | The name of the Team containing our target channel. If not specified, defaults to `unspecified`.                                                  |
| `title`                    | No       |               | *valid title string*                                      | The (optional) title for the message to submit.                                                                                                   |
| `sender`                   | No       |               | *valid application or script name*                        | The (optional) sending application name or generator of the message this app will attempt to deliver.                                             |
| `url`                      | Yes      |               | [*valid Microsoft Office 365 Webhook URL*](#webhook-urls) | The Webhook URL provided by a pre-configured Connector.                                                                                           |
| `target-url`               | No       |               | *valid comma-separated `url`, `description` pair*         | The target URL and label (specified as comma separated pair) usually visible as a button towards the bottom of the Microsoft Teams message.       |
| `verbose`                  | No       | `false`       | `true`, `false`                                           | Whether detailed output should be shown after message submission success or failure                                                               |
| `silent`                   | No       | `false`       | `true`, `false`                                           | Whether ANY output should be shown after message submission success or failure                                                                    |
| `convert-eol`              | No       | `false`       | `true`, `false`                                           | Whether messages with Windows, Mac and Linux newlines are updated to use break statements before message submission                               |
| `disable-url-validation`   | No       | `false`       | `true`, `false`                                           | Whether webhook URL validation should be disabled. Useful when submitting generated JSON payloads to a service like <https://httpbin.org/>.       |
| `disable-branding-trailer` | No       | `false`       | `true`, `false`                                           | Whether the branding trailer should be omitted from all messages generated by this application.                                                   |
| `ignore-invalid-response`  | No       | `false`       | `true`, `false`                                           | Whether an invalid response from remote endpoint should be ignored. This is expected if submitting a message to a non-standard webhook URL.       |
| `retries`                  | No       | `2`           | *positive whole number*                                   | The number of attempts that this application will make to deliver messages before giving up.                                                      |
| `retries-delay`            | No       | `2`           | *positive whole number*                                   | The number of seconds that this application will wait before making another delivery attempt.                                                     |
| `user-mention`             | No       |               | *one or more valid comma-separated `name`, `id` pairs*    | The DisplayName and ID of the recipient (specified as comma separated pair) for a user mention. May be repeated to create multiple user mentions. |

## Limitations

### message size

Per official documentation (see [references](#references)), each message sent
to Microsoft Teams can be approximately 28 KB. This includes the message
itself (text, image links, etc.), @-mentions, and reactions.

## Examples

### One-off

This example illustrates the basics of using the application to submit a
single message. This can serve as a starting point for use with Nagios,
scripts or any other tool that calls out to others in order to perform its
tasks.

The same example, shown split over multiple lines for readability (e.g., shell
script):

```console
send2teams \
  --silent \
  --channel "Alerts" \
  --team "Support" \
  --message "System XYZ is down!" \
  --title "System outage alert" \
  --sender "Nagios" \
  --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

and on a single line (e.g., one-off via terminal or batch file):

```console
send2teams.exe --silent --channel "Alerts" --team "Support" --message "System XYZ is down!" --title "System outage alert" --sender "Nagios" --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

Note:

- remove the `-silent` flag in order to see pass or failure output
- use the `-verbose` flag to see the JSON payload submitted to Microsoft Teams
- check the exit code (`$?`) to determine overall success/failure result

### Using an invalid flag

Accidentally typing the wrong flag results in a message like this one:

```ShellSession
flag provided but not defined: -fake-flag
```

### Specifying url, description pairs

```console
send2teams \
  --silent \
  --channel "Alerts" \
  --team "Support" \
  --message "Useful starting points" \
  --title "Learn more about Go" \
  --sender "Nagios" \
  --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz" \
  --target-url "https://go.dev/, Go Homepage" \
  --target-url "https://github.com/dariubs/GoBooks, Awesome Go Books"
```

and on a single line (e.g., one-off via terminal or batch file):

```console
send2teams.exe --silent --channel "Alerts" --team "Support" --message "Useful starting points" --title "Learn more about Go" --sender "Nagios" --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz" --target-url "https://go.dev/, Go Homepage" --target-url "https://github.com/dariubs/GoBooks, Awesome Go Books"
```

### User mentions

#### One mention

This example illustrates mentioning a user along with providing a brief
message.

The example, shown split over multiple lines for readability (e.g., shell
script):

```console
send2teams \
  --silent \
  --channel "Alerts" \
  --team "Support" \
  --message "System XYZ is down!" \
  --user-mention "John Doe,john.doe@example.com" \
  --sender "Nagios" \
  --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

and on a single line (e.g., one-off via terminal or batch file):

```console
send2teams --silent --channel "Alerts" --team "Support" --message "System XYZ is down!" --user-mention "John Doe,john.doe@example.com" --sender "Nagios" --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

Note:

- remove the `-silent` flag in order to see pass or failure output
- use the `-verbose` flag to see the JSON payload submitted to Microsoft Teams
- check the exit code (`$?`) to determine overall success/failure result

#### Multiple mentions

This example illustrates mentioning multiple users along with providing a
brief message. The `--user-mention` flag is repeated for each user mention.

Though valid syntax, repeating the same user mention does not increase the
number of times the same user is notified of a user mention.

The example, shown split over multiple lines for readability (e.g., shell
script):

```console
send2teams \
  --silent \
  --channel "Alerts" \
  --team "Support" \
  --message "System XYZ is down!" \
  --user-mention "John Doe,john.doe@example.com" \
  --user-mention "Jane Doe,jane.doe@example.com" \
  --sender "Nagios" \
  --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

and on a single line (e.g., one-off via terminal or batch file):

```console
send2teams --silent --channel "Alerts" --team "Support" --message "System XYZ is down!" --user-mention "John Doe,john.doe@example.com" --user-mention "Jane Doe,jane.doe@example.com" --sender "Nagios" --url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

Note:

- remove the `-silent` flag in order to see pass or failure output
- use the `-verbose` flag to see the JSON payload submitted to Microsoft Teams
- check the exit code (`$?`) to determine overall success/failure result

## License

From the [LICENSE](LICENSE) file:

```license
MIT License

Copyright 2021 Adam Chalkley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```

## References

- Related projects
  - <https://github.com/atc0005/go-teams-notify/>

- Webhook / Office 365
  - <https://sankalpit.com/how-to-get-channel-webhook-url/>
  - <https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference>
  - <https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/outgoingwebhook>
  - <https://docs.microsoft.com/en-us/outlook/actionable-messages/send-via-connectors>
  - <https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/connectors-using>
    - <https://gist.github.com/chusiang/895f6406fbf9285c58ad0a3ace13d025#gistcomment-3562510>
  - <https://messagecardplayground.azurewebsites.net/>
  - <https://docs.microsoft.com/en-us/microsoftteams/limits-specifications-teams>

- General Go topics of interest
  - <https://stackoverflow.com/questions/38807903/how-do-i-handle-plain-text-http-get-response-in-golang>
  - <https://stackoverflow.com/questions/32042989/go-lang-differentiate-n-and-line-break>
    - <https://stackoverflow.com/a/42793954/903870>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/send2teams>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
