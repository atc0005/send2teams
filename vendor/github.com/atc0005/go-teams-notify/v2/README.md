<!-- omit in toc -->
# go-teams-notify

A package to send messages to Microsoft Teams (channels)

[![Latest release][githubtag-image]][githubtag-url]
[![GoDoc][godoc-image]][godoc-url]
[![License][license-image]][license-url]
[![Validate Codebase](https://github.com/atc0005/go-teams-notify/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/go-teams-notify/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/go-teams-notify/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/go-teams-notify/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/go-teams-notify/actions?query=workflow%3A%22Quick+Validation%22)

<!-- omit in toc -->
## Table of contents

- [Project home](#project-home)
- [Project status](#project-status)
- [Project goals](#project-goals)
- [Overview](#overview)
- [Features](#features)
- [Changelog](#changelog)
- [Usage](#usage)
- [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/go-teams-notify) for the
latest code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Project status

This fork is now a standalone project.

The last [upstream project
release](https://github.com/dasrick/go-teams-notify/releases) was `v2.1.0`. I
have tried multiple times to reach the project maintainer (through a variety
of means) and having been unsuccessful, I have opted to take all
pending/proposed changes for that project and implement them here.

With work having stalled on the upstream project, others have also taken an
interest in [maintaining their own
forks](https://github.com/atc0005/go-teams-notify/network/members) of the
parent project codebase. See those forks for other ideas/changes that you may
find useful.

## Project goals

[v2.1.1](https://github.com/atc0005/go-teams-notify/releases/tag/v2.1.1) was
the first release from this project since diverging from upstream. As I write
this, `v2.2.0` is nearly ready for release. I have further changes that I plan
to incorporate into this fork from another project that I maintain,
[atc0005/send2teams](https://github.com/atc0005/send2teams).

After that work is complete, this project is likely to enter a "maintenance"
mode in favor of other projects that I maintain. I plan to apply bugfixes,
maintain dependencies and make improvements as warranted to meet the needs of
dependent projects.

All of that said, I hope to eventually collapse this project and merge all
changes back into the upstream project.

While forking a project allows bypassing hangups with local development, it
also fragments the user base. Unless absolutely necessary (as may be the case
here based on lack of respone from upstream), that carries a potentially high
cost of maintenance across all forks that wish to remain up to date/relevant.

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

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Usage

To get the package, execute:

```console
go get https://github.com/atc0005/go-teams-notify/v2
```

To import this package, add the following line to your code:

```golang
import "github.com/atc0005/go-teams-notify/v2"
```

And this is an example of a simple implementation ...

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

[godoc-image]: https://godoc.org/github.com/atc0005/go-teams-notify?status.svg
[godoc-url]: https://godoc.org/github.com/atc0005/go-teams-notify

[license-image]: https://img.shields.io/github/license/atc0005/go-teams-notify.svg?style=flat
[license-url]: https://github.com/atc0005/go-teams-notify/blob/master/LICENSE
