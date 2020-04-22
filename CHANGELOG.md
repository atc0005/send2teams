# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/send2teams/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.4.1] - 2020-04-22

### Changed

- message trailer now includes RFC3339 formatted datestamp for troubleshooting
  purposes

## [v0.4.0] - 2020-04-19

### Added

- Pin `atc0005/go-teams-notify` at commit
  atc0005/go-teams-notify@55cca556e7267ec69dc41180591bf666b12321f5
  - provides new `API.SendWithContext()` method

- `teams` subpackage `SendMessage()` now accepts a context and uses it to
  instrument the new `API.SendWithContext()` method

- Add default `TeamsSubmissionTimeout` to mirror original
  `dasrick/go-teams-notify` v1 http client timeout

### Changed

- `teams.SendMessage()`
  - now requires a context
  - Tweak log messages to note the current and total number of attempts allowed

## [v0.3.1] - 2020-04-18

### Fixed

- Remove internal validation func merged upstream
- Update bundled `atc0005/go-teams-notify` fork to reflect inclusion of commit
  943cdeb90f3e53d1ead03bcc1f86cb5de9b4f264

## [v0.3.0] - 2020-04-10

### Added

- Add configurable message submission retry and retry delay flag with default
  setting of two retries, two seconds apart
- `golangci-lint` config file created with current linters + `scopelint`
  linter enabled

### Changed

- `config` subpackage moved into `internal` subdirectory to make it private to
  this project
- `send2teams` app moved into `cmd` subdirectory structure

### Fixed

- Restore version embedding broken in v0.2.5
- Bump copyright year

## [v0.2.5] - 2020-04-09

### Added

- `teams` subpackage
  - intentionally exported for external use
  - the goal is to have as much of the code accepted into the
    `dasrick/go-teams-notify` project as is feasible, and maintain the
    remaining content and anything new related to Microsoft Teams for shared
    use in other projects I work with.
- `config` subpackage
  - will probably move it into an `internal` package structure at some point
    once I read more about it as it is intended only for this project to use
- `README`
  - add brief coverage of new `teams` package
  - add brief coverage of known valid webhook URL FQDNs and provided examples
    of complete webhook URLs using each of the known FQDNs

### Changed

- Using [vendoring](https://golang.org/cmd/go/#hdr-Vendor_Directories)
  - created top-level `vendor` directory using `go mod vendor`
  - locked-in specific commit from the prototype
    `test-extended-messagecard-type` branch from the `atc0005/go-teams-notify`
    fork in order to provide the required functionality used by recent changes
    to this project
  - updated GitHub Actions Workflow to specify `-mod=vendor` build flag for
    all `go` commands that I know of that respect the flag
  - updated GitHub Actions Workflow to exclude `vendor` directory from
    Markdown file linting to prevent potential linting issues in vendored
    dependencies from affecting our CI checks
  - updated `Makefile` to use `-mod=vendor` where applicable
  - updated `go vet` linting check to use `-mod=vendor`

- Updated dependencies
  - `gopkg.in/yaml.v2`
    - `v2.2.4` to `v2.28`
  - `atc0005/go-teams-notify`
    - see note above

### Fixed

- `ConvertEOLToBreak()` function updated to properly handle literal embedded
  newlines as well as the matching escape sequence

## [v0.2.4] - 2020-03-26

### Changed

- As with the `atc0005/send2teams` v0.2.3 release, this release still
  references our fork for now
  - further changes are being developed on our fork for potential inclusion
    upstream

### Fixed

- Update go.mod to use v1.3.0 of `dasrick/go-teams-notify` package
  - changes temporarily provided by our fork as noted in the
    v0.2.3 release notes have been merged upstream

## [v0.2.3] - 2020-03-25

### Changed

- Switch from upstream `dasrick/go-teams-notify` package to our fork,
  `atc0005/go-teams-notify` (intended to be temporary) in order to allow both
  valid webhook URL FQDNs
  - upstream currently only allows the (apparently) more common
    outlook.office.com FQDN
  - an issue has been filed with upstream to extend the `isValidWebhookURL()`
    validation function so that a fork is not necessary

### Fixed

- Update webhook URL validation checks
  - allow either of the known valid webhook URL FQDNs
    - outlook.office.com
    - outlook.office365.com
  - webhook URL length check to fail early with (hopefully) a useful error
    message
  - full regex pattern check in an effort to help catch poorly formatted
    webhook URLs

## [v0.2.2] - 2020-03-23

### Added

- GitHub Actions Workflow
  - print Go version used
    - intended as a future troubleshooting aid

### Fixed

- README
  - formatting for flags table
- Code
  - "slice bounds out of range" panic due to incorrect validity check against
    webhook URL pattern
- GitHub Actions Workflow
  - use current Go versions
    - remove Go v1.12
    - add Go v1.14

## [v0.2.1] - 2019-12-19

### Fixed

- Add missing flag in help output
- Remove forced line break/wrapping since GoDoc interprets
  this as a code block instead of continuing the line

## [v0.2.0] - 2019-12-19

### Added

- Optional conversion of messages with Windows, Mac or Linux newlines to
  `<br>` to increase compatibility with Teams formatting

## [v0.1.0] - 2019-12-18

### Added

This initial prototype supports/provides:

- Command-line flags support via `flag` standard library package
- Go modules (vs classic `GOPATH` setup)
- GitHub Actions linting and build checks
- Makefile for general use cases

[Unreleased]: https://github.com/atc0005/send2teams/compare/v0.4.1...HEAD
[v0.4.1]: https://github.com/atc0005/send2teams/releases/tag/v0.4.1
[v0.4.0]: https://github.com/atc0005/send2teams/releases/tag/v0.4.0
[v0.3.1]: https://github.com/atc0005/send2teams/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/send2teams/releases/tag/v0.3.0
[v0.2.5]: https://github.com/atc0005/send2teams/releases/tag/v0.2.5
[v0.2.4]: https://github.com/atc0005/send2teams/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/send2teams/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/send2teams/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/send2teams/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/send2teams/releases/tag/v0.2.0
[v0.1.0]: https://github.com/atc0005/send2teams/releases/tag/v0.1.0
