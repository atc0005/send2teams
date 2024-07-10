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
  - [Setup a connection to Microsoft Teams](#setup-a-connection-to-microsoft-teams)
    - [Overview](#overview-1)
    - [Workflow connectors](#workflow-connectors)
      - [Workflow webhook URL format](#workflow-webhook-url-format)
      - [How to create a Workflow connector webhook URL](#how-to-create-a-workflow-connector-webhook-url)
        - [Using Teams client Workflows context option](#using-teams-client-workflows-context-option)
        - [Using Teams client app](#using-teams-client-app)
        - [Using Power Automate web UI](#using-power-automate-web-ui)
    - [O365 connectors](#o365-connectors)
      - [O365 webhook URL format](#o365-webhook-url-format)
      - [How to create an O365 connector webhook URL](#how-to-create-an-o365-connector-webhook-url)
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

### Setup a connection to Microsoft Teams

#### Overview

> [!WARNING]
>
> Microsoft announced July 3rd, 2024 that Office 365 (O365) connectors within
Microsoft Teams would be [retired in 3
months][o365-connector-retirement-announcement] and replaced by Power Automate
workflows (or just "Workflows" for short).

Quoting from the microsoft365dev blog:

> We will gradually roll out this change in waves:
>
> - Wave 1 - effective August 15th, 2024: All new Connector creation will be
>   blocked within all clouds
> - Wave 2 - effective October 1st, 2024: All connectors within all clouds
>   will stop working

As noted, Existing O365 connector webhook URLs *should* continue to work until
2024-10-01.

#### Workflow connectors

##### Workflow webhook URL format

Valid Power Automate Workflow URLs used to submit messages to Microsoft Teams
use this format:

- `https://*.logic.azure.com:443/workflows/GUID_HERE/triggers/manual/paths/invoke?api-version=YYYY-MM-DD&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=SIGNATURE_HERE`

Example URL from the LinkedIn [Bring Microsoft Teams incoming webhook security to
the next level with Azure Logic App][linkedin-teams-webhook-security-article]
article:

- `https://webhook-jenkins.azure-api.net/manual/paths/invoke?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=f2QjZY50uoRnX6PIpyPT3xk`

##### How to create a Workflow connector webhook URL

> [!TIP]
>
> Use a dedicated "service" account not tied to a specific team member to help
ensure that the Workflow connector is long lived.

The [initial O365 retirement blog
post][o365-connector-retirement-announcement] provides a list of templates
which guide you through the process of creating a Power Automate Workflow
webhook URL.

###### Using Teams client Workflows context option

1. Navigate to a channel or chat
1. Select the ellipsis on the channel or chat
1. Select `Workflows`
1. Type `when a webhook request`
1. Select the appropriate template
   - `Post to a channel when a webhook request is received`
   - `Post to a chat when a webhook request is received`
1. Verify that `Microsoft Teams` is successfully enabled
1. Select `Next`
1. Select an appropriate value from the `Microsoft Teams Team` drop-down list.
1. Select an appropriate `Microsoft Teams Channel` drop-down list.
1. Select `Create flow`
1. Copy the new workflow URL
1. Select `Done`

###### Using Teams client app

1. Open `Workflows` application in teams
1. Select `Create` across the top of the UI
1. Choose `Notifications` at the left
1. Select `Post to a channel when a webhook request is received`
1. Verify that `Microsoft Teams` is successfully enabled
1. Select `Next`
1. Select an appropriate value from the `Microsoft Teams Team` drop-down list.
1. Select an appropriate `Microsoft Teams Channel` drop-down list.
1. Select `Create flow`
1. Copy the new workflow URL
1. Select `Done`

###### Using Power Automate web UI

[This][workflow-channel-post-from-webhook-request] template walks you through
the steps of creating a new Workflow using the
<https://make.powerautomate.com/> web UI:

1. Select or create a new connection (e.g., <user@example.com>) to Microsoft
   Teams
1. Select `Create`
1. Select an appropriate value from the `Microsoft Teams Team` drop-down list.
1. Select an appropriate `Microsoft Teams Channel` drop-down list.
1. Select `Create`
1. If prompted, read the info message (e.g., "Your flow is ready to go") and
   dismiss it.
1. Select `Edit` from the menu across the top
   - alternatively, select `My flows` from the side menu, then select `Edit`
     from the "More commands" ellipsis
1. Select `When a Teams webhook request is received` (e.g., left click)
1. Copy the `HTTP POST URL` value
   - this is your *private* custom Workflow connector URL
   - by default anyone can `POST` a request to this Workflow connector URL
     - while this access setting can be changed it will prevent this library
       from being used to submit webhook requests

#### O365 connectors

##### O365 webhook URL format

> [!WARNING]
>
> O365 connector webhook URLs are deprecated and [scheduled to be
retired][o365-connector-retirement-announcement] on 2024-10-01.

Valid (***deprecated***) O365 webhook URLs for Microsoft Teams use one of several
(confirmed) FQDNs patterns:

- `outlook.office.com`
- `outlook.office365.com`
- `*.webhook.office.com`
  - e.g., `example.webhook.office.com`

Using an O365 webhook URL with any of these FQDN patterns appears to give
identical results.

Here are complete, equivalent example webhook URLs from Microsoft's
documentation using the FQDNs above:

- <https://outlook.office.com/webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>
- <https://outlook.office365.com/webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>
- <https://example.webhook.office.com/webhookb2/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>
  - note the `webhookb2` sub-URI specific to this FQDN pattern

All of these patterns when provided to this library should pass the default
validation applied. See the example further down for the option of disabling
webhook URL validation entirely.

##### How to create an O365 connector webhook URL

> [!WARNING]
>
> O365 connector webhook URLs are deprecated and [scheduled to be
retired][o365-connector-retirement-announcement] on 2024-10-01.

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

| Flag                       | Required | Default       | Possible                                                      | Description                                                                                                                                       |
| -------------------------- | -------- | ------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`                | No       | N/A           | N/A                                                           | Display Help; show available flags.                                                                                                               |
| `v`, `version`             | No       | `false`       | `true`, `false`                                               | Whether to display application version and then immediately exit application.                                                                     |
| `channel`                  | No       | `unspecified` | *valid Microsoft Teams channel name*                          | The target channel where we will send a message. If not specified, defaults to `unspecified`.                                                     |
| `color`                    | No       | `NotUsed`     | N/A                                                           | NOOP; this setting is no longer used. Values specified for this flag are ignored.                                                                 |
| `message`                  | Yes      |               | *valid message string*                                        | The (optionally) Markdown-formatted message to submit.                                                                                            |
| `team`                     | No       | `unspecified` | *valid Microsoft Teams team name*                             | The name of the Team containing our target channel. If not specified, defaults to `unspecified`.                                                  |
| `title`                    | No       |               | *valid title string*                                          | The (optional) title for the message to submit.                                                                                                   |
| `sender`                   | No       |               | *valid application or script name*                            | The (optional) sending application name or generator of the message this app will attempt to deliver.                                             |
| `url`                      | Yes      |               | [*valid Webhook URL*](#setup-a-connection-to-microsoft-teams) | The Webhook URL provided by a pre-configured Connector.                                                                                           |
| `target-url`               | No       |               | *valid comma-separated `url`, `description` pair*             | The target URL and label (specified as comma separated pair) usually visible as a button towards the bottom of the Microsoft Teams message.       |
| `verbose`                  | No       | `false`       | `true`, `false`                                               | Whether detailed output should be shown after message submission success or failure                                                               |
| `silent`                   | No       | `false`       | `true`, `false`                                               | Whether ANY output should be shown after message submission success or failure                                                                    |
| `convert-eol`              | No       | `false`       | `true`, `false`                                               | Whether messages with Windows, Mac and Linux newlines are updated to use break statements before message submission                               |
| `disable-url-validation`   | No       | `false`       | `true`, `false`                                               | Whether webhook URL validation should be disabled. Useful when submitting generated JSON payloads to a service like <https://httpbin.org/>.       |
| `disable-branding-trailer` | No       | `false`       | `true`, `false`                                               | Whether the branding trailer should be omitted from all messages generated by this application.                                                   |
| `ignore-invalid-response`  | No       | `false`       | `true`, `false`                                               | Whether an invalid response from remote endpoint should be ignored. This is expected if submitting a message to a non-standard webhook URL.       |
| `retries`                  | No       | `2`           | *positive whole number*                                       | The number of attempts that this application will make to deliver messages before giving up.                                                      |
| `retries-delay`            | No       | `2`           | *positive whole number*                                       | The number of seconds that this application will wait before making another delivery attempt.                                                     |
| `user-mention`             | No       |               | *one or more valid comma-separated `name`, `id` pairs*        | The DisplayName and ID of the recipient (specified as comma separated pair) for a user mention. May be repeated to create multiple user mentions. |

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

<!--
  TODO: Refresh/replace these ref links after 2024-10-01 when O365 connectors are scheduled to be retired.
-->
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

[o365-connector-retirement-announcement]: <https://devblogs.microsoft.com/microsoft365dev/retirement-of-office-365-connectors-within-microsoft-teams/> "Retirement of Office 365 connectors within Microsoft Teams"
[workflow-channel-post-from-webhook-request]: <https://make.preview.powerautomate.com/galleries/public/templates/d271a6f01c2545a28348d8f2cddf4c8f/post-to-a-channel-when-a-webhook-request-is-received> "Post to a channel when a webhook request is received"
[linkedin-teams-webhook-security-article]: <https://www.linkedin.com/pulse/bring-microsoft-teams-incoming-webhook-security-next-level-kinzelin> "Bring Microsoft Teams incoming webhook security to the next level with Azure Logic App"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
