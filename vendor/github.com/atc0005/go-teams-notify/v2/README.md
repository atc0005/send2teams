<!-- omit in toc -->
# go-teams-notify

A package to send messages to Microsoft Teams (channels)

[![Latest release][githubtag-image]][githubtag-url]
[![Go Reference][goref-image]][goref-url]
[![License][license-image]][license-url]
[![Validate Codebase](https://github.com/atc0005/go-teams-notify/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/go-teams-notify/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/go-teams-notify/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/go-teams-notify/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Quick+Validation%22)

<!-- omit in toc -->
## Table of contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
- [Project Status](#project-status)
  - [Now](#now)
  - [Past](#past)
  - [Future](#future)
- [Changelog](#changelog)
- [Usage](#usage)
  - [Add this project as a dependency](#add-this-project-as-a-dependency)
  - [Example: Basic](#example-basic)
  - [Example: Disable webhook URL prefix validation](#example-disable-webhook-url-prefix-validation)
- [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/go-teams-notify) for the
latest code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Overview

The `goteamsnotify` package (aka, `go-teams-notify`) allows sending simple or
complex messages to a Microsoft Teams channel.

Simple messages can be composed of only a title and a text body with complex
messages supporting multiple sections, key/value pairs (aka, `Facts`) and/or
externally hosted images.

## Features

- Generate simple or complex messages
  - simple messages consist of only a title and a text body (one or more
    strings)
  - complex messages consist of one or more sections, key/value pairs (aka,
    `Facts`) and/or externally hosted images. or images (hosted externally)
- Submit messages to Microsoft Teams

## Project Status

### Now

This fork is now a standalone project.

This project should be considered to be in "maintenance" mode. Further
contributions and but fixes are welcome, but the overall cadence and priority
is likely to be lower in comparison to other projects that I maintain. I plan
to apply bugfixes, maintain dependencies and make improvements as warranted to
meet the needs of dependent projects.

With work having stalled on the upstream project, others have also taken an
interest in [maintaining their own
forks](https://github.com/atc0005/go-teams-notify/network/members) of the
parent project codebase. See those forks for other ideas/changes that you may
find useful.

### Past

1. Initial release up to and including `v2.1.0`
   - The last [upstream project
     release](https://github.com/dasrick/go-teams-notify/releases).
1. [v2.1.1](https://github.com/atc0005/go-teams-notify/releases/tag/v2.1.1)
   - Further attempts to reach the upstream project maintainer failed.
   - I promoted my PR-only fork into a standalone project.
   - The first release from this project since diverging from upstream.
1. [v2.2.0](https://github.com/atc0005/go-teams-notify/releases/tag/v2.2.0)
   onward
   - I merged vendored local changes from another project that I maintain,
     [atc0005/send2teams](https://github.com/atc0005/send2teams).

### Future

I hope to eventually collapse this project and merge all changes back into the
upstream project. As of early 2021 however, I've still not heard back from the
upstream project maintainer, so this does not look to be the case any time
soon.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Usage

### Add this project as a dependency

Assuming that you're using [Go
Modules](https://blog.golang.org/using-go-modules), add this line to your
imports like so:

```golang
import "github.com/atc0005/go-teams-notify/v2"
```

Your editor will likely resolve the import and update your `go.mod` and
`go.sum` files accordingly. If not, read the official [Go
Modules](https://blog.golang.org/using-go-modules) blog post on the topic for
further information.

### Example: Basic

Here is an example of a simple client application which uses this library:

```golang
import (
  "github.com/atc0005/go-teams-notify/v2"
)

func main() {
  _ = sendTheMessage()
}

func sendTheMessage() error {
  // init the client
  mstClient := goteamsnotify.NewClient()

  // setup webhook url
  webhookUrl := "https://outlook.office.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"

  // setup message card
  msgCard := goteamsnotify.NewMessageCard()
  msgCard.Title = "Hello world"
  msgCard.Text = "Here are some examples of formatted stuff like "+
      "<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
  msgCard.ThemeColor = "#DF813D"

  // send
  return mstClient.Send(webhookUrl, msgCard)
}
```

Of note:

- default timeout
- package-level logging is disabled by default
- known webhook URL prefix validation is *enabled*
- simple message submitted to Microsoft Teams consisting of formatted body and
  title

### Example: Disable webhook URL prefix validation

This example disables the validation of known webhook URL prefixes so that
custom/private webhook URL endpoints can be used.

```golang
import (
  "github.com/atc0005/go-teams-notify/v2"
)

func main() {
  _ = sendTheMessage()
}

func sendTheMessage() error {
  // init the client
  mstClient := goteamsnotify.NewClient()

  // setup webhook url
  webhookUrl := "https://example.webhook.office.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"

  // Disable webhook URL prefix validation
  mstClient.SkipWebhookURLValidationOnSend(true)

  // setup message card
  msgCard := goteamsnotify.NewMessageCard()
  msgCard.Title = "Hello world"
  msgCard.Text = "Here are some examples of formatted stuff like "+
      "<br> * this list itself  <br> * **bold** <br> * *italic* <br> * ***bolditalic***"
  msgCard.ThemeColor = "#DF813D"

  // send
  return mstClient.Send(webhookUrl, msgCard)
}
```

Of note:

- known webhook URL prefix validation is **disabled**
  - allows use of custom/private webhook URL endpoints
- other settings are the same as the basic example previously listed

## References

- [Original project](https://github.com/dasrick/go-teams-notify)
- [Forks of original project](https://github.com/atc0005/go-teams-notify/network/members)

- Microsoft Teams
  - MS Teams - adaptive cards
  ([de-de](https://docs.microsoft.com/de-de/outlook/actionable-messages/adaptive-card),
  [en-us](https://docs.microsoft.com/en-us/outlook/actionable-messages/adaptive-card))
  - MS Teams - send via connectors
  ([de-de](https://docs.microsoft.com/de-de/outlook/actionable-messages/send-via-connectors),
  [en-us](https://docs.microsoft.com/en-us/outlook/actionable-messages/send-via-connectors))
  - [adaptivecards.io](https://adaptivecards.io/designer)

- Related projects
  - <https://github.com/atc0005/send2teams>
  - <https://github.com/atc0005/bounce>
  - <https://github.com/atc0005/brick>

[githubtag-image]: https://img.shields.io/github/release/atc0005/go-teams-notify.svg?style=flat
[githubtag-url]: https://github.com/atc0005/go-teams-notify

[goref-image]: https://pkg.go.dev/badge/github.com/atc0005/go-teams-notify/v2.svg
[goref-url]: https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2

[license-image]: https://img.shields.io/github/license/atc0005/go-teams-notify.svg?style=flat
[license-url]: https://github.com/atc0005/go-teams-notify/blob/master/LICENSE
