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
  - [History](#history)
  - [Future](#future)
- [Changelog](#changelog)
- [Usage](#usage)
  - [Add this project as a dependency](#add-this-project-as-a-dependency)
  - [Webhook URLs](#webhook-urls)
    - [Expected format](#expected-format)
    - [How to create a webhook URL (Connector)](#how-to-create-a-webhook-url-connector)
  - [Example: Basic](#example-basic)
  - [Example: Disable webhook URL prefix validation](#example-disable-webhook-url-prefix-validation)
  - [Example: Enable custom patterns' validation](#example-enable-custom-patterns-validation)
- [Used by](#used-by)
- [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/go-teams-notify) for the
latest code, to file an issue or submit improvements for review and potential
inclusion into the project.

## Overview

The `goteamsnotify` package (aka, `go-teams-notify`) allows sending simple or
complex messages to a Microsoft Teams channel.

Simple messages can be composed of only a title and a text body. More complex
messages can be composed of multiple sections, key/value pairs (aka, `Facts`)
and/or externally hosted images.

## Features

- Generate simple or complex messages
  - simple messages consist of only a title and a text body (one or more
    strings)
  - complex messages consist of one or more sections, key/value pairs (aka,
    `Facts`) and/or externally hosted images. or images (hosted externally)
- Configurable validation of webhook URLs
  - enabled by default, attempts to match most common known webhook URL
    patterns
  - option to disable validation entirely
  - option to use custom validation patterns
- Configurable validation of `MessageCard` type
  - default assertion that bare-minimum required fields are present
  - support for providing a custom validation function to override default
    validation behavior
- Configurable timeouts
- Configurable retry support
- Submit messages to Microsoft Teams

## Project Status

### Now

This fork is now a standalone project.

This project should be considered in "maintenance" mode. Further contributions
and bug fixes are welcome, but the overall cadence and priority is likely to
be lower in comparison to other projects that I maintain. That said, I plan to
apply bugfixes, maintain dependencies and make improvements as warranted to
meet the needs of dependent projects.

With work having stalled on the upstream project, others have also taken an
interest in [maintaining their own
forks](https://github.com/atc0005/go-teams-notify/network/members) of the
parent project codebase. See those forks for other ideas/changes that you may
find useful.

### History

1. Initial release up to and including `v2.1.0`
   - The last [upstream project
     release](https://github.com/dasrick/go-teams-notify/releases).
1. [v2.1.1](https://github.com/atc0005/go-teams-notify/releases/tag/v2.1.1)
   - Further attempts to reach the upstream project maintainer failed.
   - I promoted my PR-only fork into a standalone project.
   - The first release from this project since diverging from upstream.
1. [v2.2.0](https://github.com/atc0005/go-teams-notify/releases/tag/v2.2.0)
   - I merged vendored local changes from another project that I maintain,
     [atc0005/send2teams](https://github.com/atc0005/send2teams).

For more recent changes, see the
[Releases](https://github.com/atc0005/go-teams-notify/releases) section or our
[Changelog](https://github.com/atc0005/go-teams-notify/blob/master/CHANGELOG.md).

### Future

I hope to eventually collapse this project and merge all changes back into the
upstream project. As of early 2021 however, I've still not heard back from the
upstream project maintainer, so this does not look to be the case any time
soon. In the meantime, I plan to continue maintaining this fork and making
changes as needed to support dependent projects.

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

All of these patterns when provided to this library should pass the default
validation applied. See the example further down for the option of disabling
webhook URL validation entirely.

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
- validation of known webhook URL prefixes is *enabled*
- simple message submitted to Microsoft Teams consisting of formatted body and
  title

### Example: Disable webhook URL prefix validation

This example disables the validation webhook URLs, including the validation of
known prefixes so that custom/private webhook URL endpoints can be used (e.g.,
testing purposes).

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

  // Disable webhook URL validation
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

- webhook URL validation is **disabled**
  - allows use of custom/private webhook URL endpoints
- other settings are the same as the basic example previously listed

### Example: Enable custom patterns' validation

This example demonstrates how to enable custom validation patterns for webhook URLs.

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
  webhookUrl := "https://my.domain.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"

  // Add a custom pattern for webhook URL validation
  mstClient.AddWebhookURLValidationPatterns(`^https://.*\.domain\.com/.*$`)
  // It's also possible to use multiple patterns with one call
  // mstClient.AddWebhookURLValidationPatterns(`^https://arbitrary\.example\.com/webhook/.*$`, `^https://.*\.domain\.com/.*$`)
  // To keep the default behavior and add a custom one, use something like the following:
  // mstClient.AddWebhookURLValidationPatterns(DefaultWebhookURLValidationPattern, `^https://.*\.domain\.com/.*$`)

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

- webhook URL validation uses custom pattern
  - allows use of custom/private webhook URL endpoints
- other settings are the same as the basic example previously listed

## Used by

This library is used by the following projects.

- <https://github.com/tilmorproducts/gohelpers>
- <https://github.com/nikoksr/notify/service/msteams>
- <https://github.com/tomekwlod/go-teams>
- <https://github.com/atc0005/bounce>
- <https://github.com/atc0005/brick>
- <https://github.com/atc0005/send2teams>

See the Known importers lists below for a dynamically updated list of projects
using either this library or the original project.

- original project
  - [v1](https://pkg.go.dev/github.com/dasrick/go-teams-notify?tab=importedby)
  - [v2](https://pkg.go.dev/github.com/dasrick/go-teams-notify/v2?tab=importedby)
- this fork
  - [v1](https://pkg.go.dev/github.com/atc0005/go-teams-notify?tab=importedby)
  - [v2](https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2?tab=importedby)

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

[githubtag-image]: https://img.shields.io/github/release/atc0005/go-teams-notify.svg?style=flat
[githubtag-url]: https://github.com/atc0005/go-teams-notify

[goref-image]: https://pkg.go.dev/badge/github.com/atc0005/go-teams-notify/v2.svg
[goref-url]: https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2

[license-image]: https://img.shields.io/github/license/atc0005/go-teams-notify.svg?style=flat
[license-url]: https://github.com/atc0005/go-teams-notify/blob/master/LICENSE
