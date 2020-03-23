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

[Unreleased]: https://github.com/atc0005/send2teams/compare/v0.2.2...HEAD
[v0.2.2]: https://github.com/atc0005/send2teams/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/send2teams/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/send2teams/releases/tag/v0.2.0
[v0.1.0]: https://github.com/atc0005/send2teams/releases/tag/v0.1.0
