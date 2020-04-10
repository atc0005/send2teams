# send2teams

Small CLI tool used to submit messages to Microsoft Teams.

[![Latest Release](https://img.shields.io/github/release/atc0005/send2teams.svg?style=flat-square)](https://github.com/atc0005/send2teams/releases/latest)
[![GoDoc](https://godoc.org/github.com/atc0005/send2teams?status.svg)](https://godoc.org/github.com/atc0005/send2teams)
![Validate Codebase](https://github.com/atc0005/send2teams/workflows/Validate%20Codebase/badge.svg)
![Validate Docs](https://github.com/atc0005/send2teams/workflows/Validate%20Docs/badge.svg)

- [send2teams](#send2teams)
  - [Project home](#project-home)
  - [Overview](#overview)
  - [Features](#features)
  - [Changelog](#changelog)
  - [Requirements](#requirements)
  - [How to install it](#how-to-install-it)
  - [Configuration Options](#configuration-options)
    - [Webhook URLs](#webhook-urls)
    - [Command-line](#command-line)
  - [Examples](#examples)
    - [One-off](#one-off)
  - [License](#license)
  - [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/send2teams) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Overview

First of all, many thanks to the developer/contributors of the
`dasrick/go-teams-notify` package for making this tool possible.

This project provides several resources:

- `send2teams`
  - Small CLI tool used to submit messages to Microsoft Teams. `send2teams` is
    intended for use by Nagios, scripts or other actions that may need to submit
    pass/fail results to a MS Teams channel.

- `teams` subpackage
  - The functions provided by this package extend or enhance existing
    functionality provided by the `dasrick/go-teams-notify` package. At
    present, the focus is primarily on text formatting and conversion
    functions that make it easier for externally submitted data to be
    formatted for proper display in Microsoft Teams.

## Features

- single binary, no outside dependencies
- minimal configuration
- very few build dependencies
- optional conversion of messages with Windows, Mac or Linux newlines to
  `<br>` to increase compatibility with Teams formatting
- exported `teams` package to handle formatting text content as code for
  proper display within Microsoft Teams
- message delivery retry support with retry and retry delay values
  configurable via flag

Worth noting: This project uses Go modules (vs classic `GOPATH` setup)

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

- Go 1.12+ (for building)
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

Tested using:

- Go 1.13+
- Windows 10 Version 1903
  - native
  - WSL
- Ubuntu Linux 16.04+

## How to install it

1. [Download](https://golang.org/dl/) Go
1. [Install](https://golang.org/doc/install) Go
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
   - for current operating system with default `go` build options
     - `go build ./cmd/send2teams/`
       - Go 1.14+ automatically uses bundled dependencies in top-level
         `vendor` folder
       - Go 1.11, 1.12 and 1.13 will default to fetching dependencies
     - `go build -mod=vendor ./cmd/send2teams/`
       - force build to use bundled dependencies in top-level `vendor` folder
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for Windows
      - `make windows`
   - for Linux
     - `make linux`
1. Copy the applicable binary to whatever systems that need to run it
   1. Linux: `/tmp/send2teams/send2teams`
   1. Windows: `/tmp/send2teams/send2teams.exe`

## Configuration Options

### Webhook URLs

Valid webhook URLs for Microsoft Teams use one of two (confirmed) FQDNs:

- `outlook.office.com`
- `outlook.office365.com`

Using a webhook URL with either FQDN appears to give identical results.

Here is a complete example webhook URL from Microsoft's documentation using
both FQDNs:

- <https://outlook.office.com/webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>

- <https://outlook.office365.com/webhook/a1269812-6d10-44b1-abc5-b84f93580ba0@9e7b80c7-d1eb-4b52-8582-76f921e416d9/IncomingWebhook/3fdd6767bae44ac58e5995547d66a4e4/f332c8d9-3397-4ac5-957b-b8e3fc465a8c>

### Command-line

Currently `send2teams` only supports command-line configuration flags.
Requests for other configuration sources will be considered.

| Flag            | Required | Default   | Possible                                                  | Description                                                                                                         |
| --------------- | -------- | --------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`     | No       | N/A       | N/A                                                       | Display Help; show available flags.                                                                                 |
| `v`, `version`  | No       | `false`   | `true`, `false`                                           | Whether to display application version and then immediately exit application.                                       |
| `channel`       | Yes      |           | *valid Microsoft Teams channel name*                      | The target channel where we will send a message.                                                                    |
| `color`         | No       | `#832561` | *valid hex color code with leading `#`*                   | The hex color code used to set the desired trim color on submitted messages.                                        |
| `message`       | Yes      |           | *valid message string*                                    | The (optionally) Markdown-formatted message to submit.                                                              |
| `team`          | Yes      |           | *valid Microsoft Teams team name*                         | The name of the Team containing our target channel.                                                                 |
| `title`         | Yes      |           | *valid title string*                                      | The title for the message to submit.                                                                                |
| `url`           | Yes      |           | [*valid Microsoft Office 365 Webhook URL*](#webhook-urls) | The Webhook URL provided by a pre-configured Connector.                                                             |
| `verbose`       | No       | `false`   | `true`, `false`                                           | Whether detailed output should be shown after message submission success or failure                                 |
| `silent`        | No       | `false`   | `true`, `false`                                           | Whether ANY output should be shown after message submission success or failure                                      |
| `convert-eol`   | No       | `false`   | `true`, `false`                                           | Whether messages with Windows, Mac and Linux newlines are updated to use break statements before message submission |
| `retries`       | No       | `2`       | *positive whole number*                                   | The number of attempts that this application will make to deliver messages before giving up.                        |
| `retries-delay` | No       | `2`       | *positive whole number*                                   | The number of seconds that this application will wait before making another delivery attempt.                       |

## Examples

### One-off

This example illustrates the basics of using the application to submit a
single message. This can serve as a starting point for use with Nagios,
scripts or any other tool that calls out to others in order to perform its
tasks.

```ShellSession
./send2teams.exe -silent -channel "Testing" -message "Testing from command-line!" -title "Another test" -color "#832561" -url "https://outlook.office.com/webhook/www@xxx/IncomingWebhook/yyy/zzz"
```

Remove the `-silent` flag in order to see pass or failure output, otherwise
look at the exit code (`$?`) or Microsoft Teams to determine results.

Accidentally typing the wrong flag results in a message like this one:

```ShellSession
flag provided but not defined: -fake-flag
```

## License

From the [LICENSE](LICENSE) file:

```license
MIT License

Copyright (c) 2019 Adam Chalkley

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

- Dependencies
  - <https://github.com/dasrick/go-teams-notify/>

- Webhook / Office 365
  - <https://sankalpit.com/how-to-get-channel-webhook-url/>
  - <https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference>
  - <https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/outgoingwebhook>
  - <https://docs.microsoft.com/en-us/outlook/actionable-messages/send-via-connectors>
  - <https://messagecardplayground.azurewebsites.net/>

- General Golang
  - <https://stackoverflow.com/questions/38807903/how-do-i-handle-plain-text-http-get-response-in-golang>
  - <https://stackoverflow.com/questions/32042989/go-lang-differentiate-n-and-line-break>
    - <https://stackoverflow.com/a/42793954/903870>
