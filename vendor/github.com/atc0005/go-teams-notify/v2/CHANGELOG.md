# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/go-teams-notify/issues) for any
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

## [v2.2.0] - 2020-8-28

### Added

- Add package-level logger
- Extend API to allow request cancellation via context
- Add formatting functions useful for text conversion
  - Convert Windows/Mac/Linux EOL to Markdown break statements
    - used to provide equivalent Teams-compatible formatting
  - Format text as code snippet
    - this inserts leading and trailing ` character to provide Markdown string
      formatting
  - Format text as code block
    - this inserts three leading and trailing ` characters to provide Markdown
      code block formatting
  - *`Try`* variants of code formatting functions
    - return formatted string if no errors, otherwise return the original
      string

### Changed

- Expose API response strings containing potential error messages
- README
  - Explicitly note that this fork is now standalone until such time that the
    upstream project resumes development/maintenance efforts

### Fixed

- CHANGELOG section link in previous release
- Invalid `RoundTripper` implementation used in `TestTeamsClientSend` test
  function
  - see `GH-46` and `GH-47`; thank you `@davecheney` for the fix!

## [v2.1.1] - 2020-08-25

### Added

- README
  - Add badges for GitHub Actions Workflows
  - Add release badge for latest project release
- Add CHANGELOG file
- Add GoDoc package-level documentation
- Extend webhook validation error handling
- Add Docker-based GitHub Actions Workflows
- Enable Dependabot updates
- Add Markdownlint config file

### Changed

- README
  - Replace badge for latest tag with latest release
  - Update GoDoc badge to reference this fork
  - Update license badge to reference this fork
  - Add new sections common to other projects that I maintain
    - table of contents
    - overview
    - changelog
    - references
    - features
- Vendor dependencies
- Update license to add @atc0005 (new) in addition to @dasrick (existing)
- Update go.mod to replace upstream with this fork
- Rename golangci-lint config file to match officially supported name
- Remove files no longer used by this fork
  - Travis CI configuration
  - editorconfig file (and settings)
- Add license header to source files
  - combined copyright statement for existing files
  - single copyright statement for new files

### Fixed

- Add missing Facts assignment in MessageCardSection
- scopelint: Fix improper range loop var reference
- Fix misc linting issues with README
- Test failure from previous upstream pull request submissions
  - `Object expected to be of type *url.Error, but was *errors.errorString`
- Misc linting issues with primary and test files

## [v2.1.0] - 2020-04-08

### Added

- `MessageCard` type includes additional fields
  - `Type` and `Context` fields provide required JSON payload
    fields
    - preset to required static values via updated
      `NewMessageCard()` constructor
  - `Summary`
    - required if `Text` field is not set, optional otherwise
  - `Sections` slice
    - `MessageCardSection` type

- Additional nested types
  - `MessageCardSection`
  - `MessageCardSectionFact`
  - `MessageCardSectionImage`

- Additional methods for `MessageCard` and nested types
  - `MessageCard.AddSection()`
  - `MessageCardSection.AddFact()`
  - `MessageCardSection.AddFactFromKeyValue()`
  - `MessageCardSection.AddImage()`
  - `MessageCardSection.AddHeroImageStr()`
  - `MessageCardSection.AddHeroImage()`

- Additional factory functions
  - `NewMessageCardSection()`
  - `NewMessageCardSectionFact()`
  - `NewMessageCardSectionImage()`

- `IsValidMessageCard()` added to check for minimum required
    field values.
  - This function has the potential to be extended
    later with additional validation steps.

- Wrapper `IsValidInput()` added to handle all validation
  needs from one location.
  - the intent was to both solve a CI erro and provide
    a location to easily extend validation checks in
    the future (if needed)

### Changed

- `MessageCard` type includes additional fields
- `NewMessageCard` factory function sets fields needed for
   required JSON payload fields
  - `Type`
  - `Context`

- `teamsClient.Send()` method updated to apply `MessageCard` struct
  validation alongside existing webhook URL validation

- `isValidWebhookURL()` exported as `IsValidWebhookURL()` so that client
  code can use the validation functionality instead of repeating the
  code
  - e.g., flag value validation for "fail early" behavior

### Known Issues

- No support in this set of changes for `potentialAction` types
  - `ViewAction`
  - `OpenUri`
  - `HttpPOST`
  - `ActionCard`
  - `InvokeAddInCommand`
    - Outlook specific based on what I read; likely not included
      in a future release due to non-Teams specific usage

## [v2.0.0] - 2020-03-29

### Breaking

- `NewClient()` will NOT return multiple values
- remove provided mock

### Changed

- switch dependency/package management tool to from `dep` to `go mod`
- switch from `golint` to `golangci-lint`
- add more golang versions to pass via travis-ci

## [v1.3.1] - 2020-3-29

### Fixed

- fix redundant error logging
- fix redundant comment

## [v1.3.0] - 2020-03-26

### Changed

- feature: allow multiple valid webhook URL FQDNs (thx @atc0005)

## [v1.2.0] - 2019-11-08

### Added

- add mock

### Changed

- update deps
- `gosimple` (shorten two conditions)

## [v1.1.1] - 2019-05-02

### Changed

- rename client interface into API
- update deps

### Fixed

- fix typo in README

## [v1.1.0] - 2019-04-30

### Added

- add missing tests
- append documentation

### Changed

- add/change to client/interface

## [v1.0.0] - 2019-04-29

### Added

- add initial functionality of sending messages to MS Teams channel

[Unreleased]: https://github.com/atc0005/go-teams-notify/compare/v2.2.0...HEAD
[v2.2.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v2.2.0
[v2.1.1]: https://github.com/atc0005/go-teams-notify/releases/tag/v2.1.1
[v2.1.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v2.1.0
[v2.0.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v2.0.0
[v1.3.1]: https://github.com/atc0005/go-teams-notify/releases/tag/v1.3.1
[v1.3.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v1.3.0
[v1.2.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v1.2.0
[v1.1.1]: https://github.com/atc0005/go-teams-notify/releases/tag/v1.1.1
[v1.1.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v1.1.0
[v1.0.0]: https://github.com/atc0005/go-teams-notify/releases/tag/v1.0.0
