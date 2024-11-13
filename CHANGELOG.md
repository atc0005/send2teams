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

## [v0.13.7] - 2024-11-13

### Changed

#### Dependency Updates

- (GH-605) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.14 to go-ci-oldstable-build-v0.21.15 in /dependabot/docker/builds
- (GH-603) Go Runtime: Bump golang from 1.22.8 to 1.22.9 in /dependabot/docker/go

## [v0.13.6] - 2024-10-18

### Changed

#### Dependency Updates

- (GH-599) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.13 to go-ci-oldstable-build-v0.21.14 in /dependabot/docker/builds
- (GH-597) Go Runtime: Bump golang from 1.22.7 to 1.22.8 in /dependabot/docker/go

## [v0.13.5] - 2024-09-25

### Changed

#### Dependency Updates

- (GH-589) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.11 to go-ci-oldstable-build-v0.21.12 in /dependabot/docker/builds
- (GH-592) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.12 to go-ci-oldstable-build-v0.21.13 in /dependabot/docker/builds
- (GH-584) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.9 to go-ci-oldstable-build-v0.21.11 in /dependabot/docker/builds
- (GH-590) Go Dependency: Bump github.com/atc0005/go-teams-notify/v2 from 2.12.0 to 2.13.0
- (GH-587) Go Runtime: Bump golang from 1.22.6 to 1.22.7 in /dependabot/docker/go

## [v0.13.4] - 2024-08-21

### Changed

#### Dependency Updates

- (GH-572) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.8 to go-ci-oldstable-build-v0.21.9 in /dependabot/docker/builds
- (GH-576) Go Dependency: Bump github.com/atc0005/go-teams-notify/v2 from 2.11.0 to 2.12.0
- (GH-575) Go Runtime: Bump golang from 1.21.13 to 1.22.6 in /dependabot/docker/go
- (GH-574) Update project to Go 1.22 series

## [v0.13.3] - 2024-08-13

### Changed

#### Dependency Updates

- (GH-563) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.7 to go-ci-oldstable-build-v0.21.8 in /dependabot/docker/builds
- (GH-560) Go Runtime: Bump golang from 1.21.12 to 1.21.13 in /dependabot/docker/go

## [v0.13.2] - 2024-08-06

### Changed

#### Dependency Updates

- (GH-552) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.6 to go-ci-oldstable-build-v0.21.7 in /dependabot/docker/builds
- (GH-555) Go Dependency: Bump github.com/atc0005/go-teams-notify/v2 from 2.11.0-rc.1 to 2.11.0

#### Other Changes

- (GH-549) Update O365 deprecation timeline details
- (GH-553) Push `REPO_VERSION` var into containers for builds

## [v0.13.1] - 2024-07-25

### Changed

#### Dependency Updates

- (GH-543) Go Dependency: Bump github.com/atc0005/go-teams-notify/v2 from 2.11.0-alpha.1 to 2.11.0-alpha.5
- (GH-544) Go Dependency: Bump github.com/atc0005/go-teams-notify/v2 from 2.11.0-alpha.5 to 2.11.0-rc.1

## [v0.13.0] - 2024-07-23

### Added

- (GH-538) Initial support for Power Automate workflow connectors

### Changed

#### Dependency Updates

- (GH-530) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.4 to go-ci-oldstable-build-v0.21.5 in /dependabot/docker/builds
- (GH-535) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.5 to go-ci-oldstable-build-v0.21.6 in /dependabot/docker/builds

### Other

- (GH-536) Remove automatic display of help text on err

### Fixed

- (GH-525) Fix Teams submission timeout value logic

## [v0.12.13] - 2024-07-10

### Changed

#### Dependency Updates

- (GH-508) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.7 to go-ci-oldstable-build-v0.20.8 in /dependabot/docker/builds
- (GH-511) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.8 to go-ci-oldstable-build-v0.21.2 in /dependabot/docker/builds
- (GH-515) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.2 to go-ci-oldstable-build-v0.21.3 in /dependabot/docker/builds
- (GH-517) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.3 to go-ci-oldstable-build-v0.21.4 in /dependabot/docker/builds
- (GH-513) Go Runtime: Bump golang from 1.21.11 to 1.21.12 in /dependabot/docker/go

## [v0.12.12] - 2024-06-07

### Changed

#### Dependency Updates

- (GH-493) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.4 to go-ci-oldstable-build-v0.20.5 in /dependabot/docker/builds
- (GH-495) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.5 to go-ci-oldstable-build-v0.20.6 in /dependabot/docker/builds
- (GH-503) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.6 to go-ci-oldstable-build-v0.20.7 in /dependabot/docker/builds
- (GH-500) Go Runtime: Bump golang from 1.21.10 to 1.21.11 in /dependabot/docker/go

### Fixed

- (GH-497) Remove inactive maligned linter
- (GH-498) Fix errcheck linting errors

## [v0.12.11] - 2024-05-11

### Changed

#### Dependency Updates

- (GH-482) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.1 to go-ci-oldstable-build-v0.20.2 in /dependabot/docker/builds
- (GH-486) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.2 to go-ci-oldstable-build-v0.20.3 in /dependabot/docker/builds
- (GH-489) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.3 to go-ci-oldstable-build-v0.20.4 in /dependabot/docker/builds
- (GH-484) Go Runtime: Bump golang from 1.21.9 to 1.21.10 in /dependabot/docker/go

## [v0.12.10] - 2024-04-11

### Changed

#### Dependency Updates

- (GH-464) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.4 to go-ci-oldstable-build-v0.16.0 in /dependabot/docker/builds
- (GH-466) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.0 to go-ci-oldstable-build-v0.16.1 in /dependabot/docker/builds
- (GH-467) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.1 to go-ci-oldstable-build-v0.19.0 in /dependabot/docker/builds
- (GH-471) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.19.0 to go-ci-oldstable-build-v0.20.0 in /dependabot/docker/builds
- (GH-476) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.0 to go-ci-oldstable-build-v0.20.1 in /dependabot/docker/builds
- (GH-474) Go Runtime: Bump golang from 1.21.8 to 1.21.9 in /dependabot/docker/go

#### Other

- (GH-473) Minor README tweaks

## [v0.12.9] - 2024-03-08

### Changed

#### Dependency Updates

- (GH-444) Update project to Go 1.21 series
- (GH-445) canary: bump golang from 1.21.6 to 1.21.7 in /dependabot/docker/go
- (GH-443) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.15.0 to go-ci-oldstable-build-v0.15.2 in /dependabot/docker/builds
- (GH-446) Update Dependabot PR prefixes
- (GH-447) Update Dependabot PR prefixes (redux)
- (GH-449) Go Dependency: Bump github.com/atc0005/go-teams-notify/v2 from 2.9.0 to 2.10.0
- (GH-453) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.2 to go-ci-oldstable-build-v0.15.3 in /dependabot/docker/builds
- (GH-458) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.3 to go-ci-oldstable-build-v0.15.4 in /dependabot/docker/builds
- (GH-460) Add todo/release label to "Go Runtime" PRs
- (GH-456) Go Runtime: Bump golang from 1.21.7 to 1.21.8 in /dependabot/docker/go

## [v0.12.8] - 2024-02-16

### Changed

#### Dependency Updates

- (GH-433) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-419) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.3 to go-ci-oldstable-build-v0.14.4 in /dependabot/docker/builds
- (GH-424) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.4 to go-ci-oldstable-build-v0.14.5 in /dependabot/docker/builds
- (GH-426) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.5 to go-ci-oldstable-build-v0.14.6 in /dependabot/docker/builds
- (GH-434) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.6 to go-ci-oldstable-build-v0.14.9 in /dependabot/docker/builds
- (GH-438) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.9 to go-ci-oldstable-build-v0.15.0 in /dependabot/docker/builds
- (GH-423) go.mod: bump github.com/atc0005/go-teams-notify/v2 from 2.8.0 to 2.9.0

## [v0.12.7] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-413) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-414) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.2 to go-ci-oldstable-build-v0.14.3 in /dependabot/docker/builds
- (GH-410) ghaw: bump github/codeql-action from 2 to 3

## [v0.12.6] - 2023-12-09

### Changed

#### Dependency Updates

- (GH-403) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-406) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.1 to go-ci-oldstable-build-v0.14.2 in /dependabot/docker/builds

## [v0.12.5] - 2023-11-15

### Changed

#### Dependency Updates

- (GH-397) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-386) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-388) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.10 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-392) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.0 in /dependabot/docker/builds
- (GH-398) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.0 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds

### Fixed

- (GH-393) Fix goconst linting errors

## [v0.12.4] - 2023-10-06

### Changed

#### Dependency Updates

- (GH-364) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-356) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.4 to go-ci-oldstable-build-v0.13.5 in /dependabot/docker/builds
- (GH-359) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-361) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-366) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-372) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-376) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.10 in /dependabot/docker/builds
- (GH-363) ghaw: bump actions/checkout from 3 to 4

## [v0.12.3] - 2023-08-18

### Added

- (GH-329) Add initial automated release notes config
- (GH-331) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.4` to `go-ci-oldstable-build-v0.13.4`
  - `atc0005/go-teams-notify`
    - `v2.8.0-alpha.1` to `v2.8.0`
- (GH-333) Update Dependabot config to monitor both branches
- (GH-351) Update project to Go 1.20 series

## [v0.12.2] - 2023-07-14

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.10` to `1.19.11`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.6` to `go-ci-oldstable-build-v0.11.4`

### Fixed

- (GH-321) Restore local CodeQL workflow

## [v0.12.1] - 2023-06-09

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.9` to `1.19.10`
  - `atc0005/go-teams-notify`
    - `v2.7.1-rc.3` to `v2.8.0-alpha.1`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.5` to `go-ci-oldstable-build-v0.10.6`
- (GH-317) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-314) Disable depguard linter

## [v0.12.0] - 2023-05-11

### Overview

- Build improvements
- Dependency updates
- built using Go 1.19.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-306) Add rootless container builds via Docker/Podman

### Changed

- Dependencies
  - `Go`
    - `1.19.7` to `1.19.9`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.3` to `go-ci-oldstable-build-v0.10.5`
- (GH-307) Add .dockerignore file for use during image builds

## [v0.11.3] - 2023-03-27

### Overview

- Dependency updates
- built using Go 1.19.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `atc0005/go-teams-notify`
    - `v2.7.1-rc.2` to `v2.7.1-rc.3`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.2` to `go-ci-oldstable-build-v0.10.3`

## [v0.11.2] - 2023-03-23

### Overview

- Dependency updates
- built using Go 1.19.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `atc0005/go-teams-notify`
    - `v2.7.1-rc.1` to `v2.7.1-rc.2`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.0` to `go-ci-oldstable-build-v0.10.2`

## [v0.11.1] - 2023-03-21

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.6` to `1.19.7`
  - `atc0005/go-teams-notify`
    - `v2.7.0` to `v2.7.1-rc.1`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.9.0` to `go-ci-oldstable-build-v0.10.0`

### Fixed

- (GH-284) Fix CHANGELOG summary for v0.11.0 release
- (GH-290) Update vuln analysis GHAW to use on.push hook
- (GH-293) internal/config/config.go:289:10: call could wrap the error with
  error-wrapping directive `%w`

## [v0.11.0] - 2023-03-06

### Overview

- Add support for generating packages
- Generated binary changes
  - filename patterns
  - compression
  - executable metadata
- Build improvements
- built using Go 1.19.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-273) Generate RPM/DEB packages using nFPM

### Changed

- (GH-272) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-274) Add version metadata to Windows executables
- (GH-275) Makefile: Compress binaries and use fixed filenames
- (GH-276) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-277) Build dev/stable releases using go-ci Docker image

## [v0.10.3] - 2023-03-06

### Overview

- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.6`
  - `atc0005/go-teams-notify`
    - `v2.7.0-rc.2` to `v2.7.0`
- (GH-266) Add Go Module Validation, Dependency Updates jobs
- (GH-268) Drop `Push Validation` workflow
- (GH-269) Rework workflow scheduling
- (GH-271) Remove `Push Validation` workflow status badge

## [v0.10.2] - 2022-12-08

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
  - `atc0005/go-teams-notify`
    - `v2.7.0-rc.1` to `v2.7.0-rc.2`
- (GH-257) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-260) Fix Makefile Go module base path detection

## [v0.10.1] - 2022-09-20

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.13` to `1.19.1`
  - `atc0005/go-teams-notify`
    - `v2.7.0-alpha.6` to `v2.7.0-rc.1`
  - `github/codeql-action`
    - `v2.1.21` to `v2.1.24`

- (GH-248) Update project to Go 1.19
- (GH-249) Update Makefile and GitHub Actions Workflows

### Fixed

- (GH-247) Add missing cmd doc file

## [v0.10.0] - 2022-08-23

### Overview

- Add support for disabling branding trailer
- Dependency updates
- built using Go 1.17.13
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-227) Add optional flag to disable branding "trailer" on submitted
  messages

### Changed

- Dependencies
  - `Go`
    - `1.17.12` to `1.17.13`

### Fixed

- (GH-245) Apply Go 1.19 specific doc comments linting fixes

## [v0.9.3] - 2022-07-21

### Overview

- Bugfixes
- Dependency updates
- built using Go 1.17.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.10` to `1.17.12`
  - `atc0005/go-teams-notify`
    - `v2.7.0-alpha.2` to `v2.7.0-alpha.6`

### Fixed

- (GH-238) Update lintinstall Makefile recipe

## [v0.9.2] - 2022-05-13

### Overview

- Dependency updates
- built using Go 1.17.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.10`

## [v0.9.1] - 2022-04-27

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.8` to `1.17.9`

## [v0.9.0] - 2022-04-11

### Overview

- Fixed support for user mentions
- Dependency updates
- Swap from legacy `MessageCard` format to `Adaptive Card`
- Flag changes
- built using Go 1.17.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.8`
  - `atc0005/go-teams-notify`
    - `v2.7.0-alpha.1` to `v2.7.0-alpha.2`
  - `actions/checkout`
    - `v2.4.0` to `v3`

- (GH-225) Microsoft Teams messages are now generated using the `Adaptive
  Card` format instead of the legacy `MessageCard` format
  - this produces some minor visual differences
  - see
    [atc0005/check-vmware#649](https://github.com/atc0005/check-vmware/pull/649/commits/37ef45cf98efdf0faa958578207ff3d0b826cea4)
    for an example of changes made to a Nagios command definition to retain
    comparable visual parity
  - see
    [atc0005/check-vmware#651](https://github.com/atc0005/check-vmware/pull/651/commits/e0f87d08c9e9db5f417e4d6104f94c039d87acea)
    for an example of improvements to the command definition using syntax
    compatible with `Adaptive Card` text formatting support
- (GH-225) The `--target-url` flag no longer enforces a set limit of 4 URL
  "buttons"
- (GH-225) `--color` flag is now a NOOP
  - produces no effect; see README for details

### Fixed

- (GH-224) README doesn't make clear that the `--user-mention` flag can be
  repeated
- (GH-225) Intermittent message submission failure when using `--user-mention`
  flag
- (GH-225) Add missing checks for use of `--silent` flag before emitting
  warning/error output

## [v0.8.0] - 2022-02-25

### Overview

- Add support for user mentions
- Requirement changes
- Dependency updates
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-197) Add initial support for user mentions

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-216) Remove message title requirement

## [v0.7.0] - 2022-02-09

### Overview

- Add flags to optionally disable validation
- CI / linting improvements
- Bugfixes
- Dependency update
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-201) Add flag for disabling webhook URL validation
- (GH-204) Add flag for disabling validation of message submission response

### Changed

- Dependencies
  - `Go`
    - (GH-195) Update go.mod file, canary Dockerfile to reflect Go 1.17
    - `1.16.12` to `1.17.6`

- (GH-203) Override default user agent with project-specific value
- (GH-206) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-210) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-198) Wrong binary name in `TestConfigInitialization()` func
- (GH-202) Update `Config.String()` to expose current boolean config settings
- (GH-208) var-declaration: should omit type string from declaration of var
  version; it will be inferred from the right-hand side (revive)

## [v0.6.6] - 2021-12-28

### Overview

- Dependency update
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.6.5] - 2021-11-08

### Overview

- Dependency update
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

## [v0.6.4] - 2021-09-23

### Overview

- Dependency update
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`

- Update README to list downloading binaries as alternative to building from
  source

## [v0.6.3] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.6.2] - 2021-07-14

### Overview

- Dependency update
- Bugfix
- CI / test changes
- built using Go 1.16.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- Add test/CI coverage for `--version` flag

### Changed

- Dependencies
  - `Go`
    - `1.16.5` to `1.16.6`

### Fixed

- handling of `--version` flag broken in `v0.6.1` release

## [v0.6.1] - 2021-07-09

### Overview

- Dependency updates
- Minor tweaks
- built using Go 1.16.5
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries

### Changed

- Refactor config initialization

- Dependencies
  - `actions/setup-node`
    - updated from `v2.1.5` to `v2.2.0`
  - `atc0005/go-teams-notify`
    - updated from `v2.5.0` to `v2.6.0`

## [v0.6.0] - 2021-06-24

### Overview

- New feature
- built using Go 1.16.5
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- `--target-url` flag (optional)
  - provided as a means of specifying up to 4 `url`, `description`
    (comma-separated) pairs for use with displaying labelled "buttons" in a
    Microsoft Teams message

## [v0.5.0] - 2021-06-18

### Overview

- Flag tweaks
- Doc updates
- built using Go 1.16.5
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- `--sender` flag (optional)
  - provided as a means of noting which application is responsible for
    generating the message that *this one* attempts to deliver to Microsoft
    Teams

### Changed

- `--team` flag is now optional
- `--channel` flag is now optional
- branding text changed from *Message `generated by`* to a more accurate
  *Message `delivered by`*
- README coverage refreshed

## [v0.4.13] - 2021-06-16

### Overview

- built using Go 1.16.5
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.3` to `1.16.5`

## [v0.4.12] - 2021-04-08

### Overview

- Misc fixes
- built using Go 1.16.3

### Changed

- Dependencies
  - Built using Go 1.16.3
    - **Statically linked**
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `actions/setup-node`
    - updated from `v2.1.4` to `v2.1.5`
  - `atc0005/go-teams-notify`
    - updated from `v2.4.2` to `v2.5.0`

### Fixed

- Linting
  - fieldalignment: struct with X pointer bytes could be Y (govet)
  - Replace deprecated linters: maligned, scopelint
  - SA1019: goteamsnotify.IsValidWebhookURL is deprecated: use
    API.ValidateWebhook() method instead. (staticcheck)

## [v0.4.11] - 2021-01-29

### Overview

- Misc fixes
- built using Go 1.15.7

### Changed

- Application timeout changed from `5s` (hard-coded ceiling) to `8s` (default,
  configurable via `retries` and `retries-delay` flags)

### Fixed

- Context cancellation (timeout) does not respect retries and retries-delay
  flag values

## [v0.4.10] - 2021-01-29

### Changed

- Documentation
  - Extend Webhook URL HowTo coverage
  - Replace godoc.org badge with pkg.go.dev badge

- Dependencies
  - Built using Go 1.15.7
    - **Statically linked**
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `actions/setup-node`
    - updated from `v2.1.2` to `v2.1.4`
  - `atc0005/go-teams-notify`
    - updated from `v2.3.0` to `v2.4.2`

### Fixed

- Fix exit code handling
- Fix typo

## [v0.4.9] - 2020-11-17

### Changed

- Dependencies
  - Built using Go 1.15.5
    - **Statically linked**
    - Windows
      - x86
      - x64
    - Linux
      - x86
      - x64
  - `actions/checkout`
    - updated from `v2.3.3` to `v2.3.4`

## [v0.4.8] - 2020-10-11

### Added

- Binary release
  - Built using Go 1.15.2
  - **Statically linked**
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

### Changed

- Dependencies
  - `actions/setup-node`
    - updated from `v2.1.1` to `v2.1.2`
  - `actions/checkout`
    - updated from `v2.3.2` to `v2.3.3`

- Add `-trimpath` build flag

### Fixed

- Makefile build options do not generate static binaries
- Misc linting errors raised by latest `gocritic` release included with
  `golangci-lint` `v1.31.0`
- Makefile generates checksums with qualified path

## [v0.4.7] - 2020-08-30

### Changed

- Dependencies
  - upgrade `atc0005/go-teams-notify`
    - `v2.2.0` to `v2.3.0`

- Exclusive use of `atc0005/go-teams-notify` for functionality previously
  provided by the (since removed) `teams` subpackage

- Documentation updates to reflect these changes

### Removed

- This project no longer provides the `teams` subpackage
  - all functionality previously provided by that package has been moved into
    the separate `atc0005/go-teams-notify` project

### Fixed

- `YYYY-MM-DD` date references in this CHANGELOG file

## [v0.4.6] - 2020-08-21

### Added

- Docker-based GitHub Actions Workflows
  - Replace native GitHub Actions with containers created and managed through
    the `atc0005/go-ci` project.

  - New, primary workflow
    - with parallel linting, testing and building tasks
    - with three Go environments
      - "old stable"
      - "stable"
      - "unstable"
    - Makefile is *not* used in this workflow
    - staticcheck linting using latest stable version provided by the
      `atc0005/go-ci` containers

  - Separate Makefile-based linting and building workflow
    - intended to help ensure that local Makefile-based builds that are
      referenced in project README files continue to work as advertised until
      a better local tool can be discovered/explored further
    - use `golang:latest` container to allow for Makefile-based linting
      tooling installation testing since the `atc0005/go-ci` project provides
      containers with those tools already pre-installed
      - linting tasks use container-provided `golangci-lint` config file
        *except* for the Makefile-driven linting task which continues to use
        the repo-provided copy of the `golangci-lint` configuration file

  - Add Quick Validation workflow
    - run on every push, everything else on pull request updates
    - linting via `golangci-lint` only
    - testing
    - no builds

### Changed

- Disable `golangci-lint` default exclusions

- dependencies
  - `go.mod` Go version
    - updated from `1.13` to `1.14`
  - `actions/setup-go`
    - updated from `v2.1.0` to `v2.1.2`
      - since replaced with Docker containers
  - `actions/setup-node`
    - updated from `v2.1.0` to `v2.1.1`
  - `actions/checkout`
    - updated from `v2.3.1` to `v2.3.2`

- README
  - Link badges to applicable GitHub Actions workflows results

- Linting
  - Local
    - `Makefile`
      - install latest stable `golangci-lint` binary instead of using a fixed
          version
  - CI
    - remove repo-provided copy of `golangci-lint` config file at start of
      linting task in order to force use of Docker container-provided config
      file

## [v0.4.5] - 2020-07-17

### Added

- Enable Dependabot version updates

### Fixed

- Context error is unintentionally masked by early return
- CHANGELOG
  - wrong section name
- README
  - incorrect path for generated binaries

### Changed

- dependencies
  - `go-yaml/yaml`
    - updated from `v2.2.8` to `v2.3.0`
  - `actions/setup-go`
    - updated from `v1` to `v2.1.0`
  - `actions/setup-node`
    - updated from `v1` to `v2.1.0`
  - `actions/checkout`
    - updated from `v1` to `v2.3.1`

## [v0.4.4] - 2020-04-28

### Fixed

- CHANGELOG formatting

## [v0.4.3] - 2020-04-28

### Fixed

- Remove bash shebang from GitHub Actions Workflow files
- Update README to list accurate build/deploy steps based
  on recent restructuring work

### Changed

- Update golangci-lint to v1.25.1
- Remove gofmt and golint as separate checks, enable
  these linters in golangci-lint config

## [v0.4.2] - 2020-04-25

### Changed

- Install specific binary version of golangci-lint instead of building from
  `master`

### Fixed

- Makefile: Linting commands do not exclude vendor subfolder
- Makefile: go vet doesn't explicitly include -mod=vendor

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

[Unreleased]: https://github.com/atc0005/send2teams/compare/v0.13.7...HEAD
[v0.13.7]: https://github.com/atc0005/send2teams/releases/tag/v0.13.7
[v0.13.6]: https://github.com/atc0005/send2teams/releases/tag/v0.13.6
[v0.13.5]: https://github.com/atc0005/send2teams/releases/tag/v0.13.5
[v0.13.4]: https://github.com/atc0005/send2teams/releases/tag/v0.13.4
[v0.13.3]: https://github.com/atc0005/send2teams/releases/tag/v0.13.3
[v0.13.2]: https://github.com/atc0005/send2teams/releases/tag/v0.13.2
[v0.13.1]: https://github.com/atc0005/send2teams/releases/tag/v0.13.1
[v0.13.0]: https://github.com/atc0005/send2teams/releases/tag/v0.13.0
[v0.12.13]: https://github.com/atc0005/send2teams/releases/tag/v0.12.13
[v0.12.12]: https://github.com/atc0005/send2teams/releases/tag/v0.12.12
[v0.12.11]: https://github.com/atc0005/send2teams/releases/tag/v0.12.11
[v0.12.10]: https://github.com/atc0005/send2teams/releases/tag/v0.12.10
[v0.12.9]: https://github.com/atc0005/send2teams/releases/tag/v0.12.9
[v0.12.8]: https://github.com/atc0005/send2teams/releases/tag/v0.12.8
[v0.12.7]: https://github.com/atc0005/send2teams/releases/tag/v0.12.7
[v0.12.6]: https://github.com/atc0005/send2teams/releases/tag/v0.12.6
[v0.12.5]: https://github.com/atc0005/send2teams/releases/tag/v0.12.5
[v0.12.4]: https://github.com/atc0005/send2teams/releases/tag/v0.12.4
[v0.12.3]: https://github.com/atc0005/send2teams/releases/tag/v0.12.3
[v0.12.2]: https://github.com/atc0005/send2teams/releases/tag/v0.12.2
[v0.12.1]: https://github.com/atc0005/send2teams/releases/tag/v0.12.1
[v0.12.0]: https://github.com/atc0005/send2teams/releases/tag/v0.12.0
[v0.11.3]: https://github.com/atc0005/send2teams/releases/tag/v0.11.3
[v0.11.2]: https://github.com/atc0005/send2teams/releases/tag/v0.11.2
[v0.11.1]: https://github.com/atc0005/send2teams/releases/tag/v0.11.1
[v0.11.0]: https://github.com/atc0005/send2teams/releases/tag/v0.11.0
[v0.10.3]: https://github.com/atc0005/send2teams/releases/tag/v0.10.3
[v0.10.2]: https://github.com/atc0005/send2teams/releases/tag/v0.10.2
[v0.10.1]: https://github.com/atc0005/send2teams/releases/tag/v0.10.1
[v0.10.0]: https://github.com/atc0005/send2teams/releases/tag/v0.10.0
[v0.9.3]: https://github.com/atc0005/send2teams/releases/tag/v0.9.3
[v0.9.2]: https://github.com/atc0005/send2teams/releases/tag/v0.9.2
[v0.9.1]: https://github.com/atc0005/send2teams/releases/tag/v0.9.1
[v0.9.0]: https://github.com/atc0005/send2teams/releases/tag/v0.9.0
[v0.8.0]: https://github.com/atc0005/send2teams/releases/tag/v0.8.0
[v0.7.0]: https://github.com/atc0005/send2teams/releases/tag/v0.7.0
[v0.6.6]: https://github.com/atc0005/send2teams/releases/tag/v0.6.6
[v0.6.5]: https://github.com/atc0005/send2teams/releases/tag/v0.6.5
[v0.6.4]: https://github.com/atc0005/send2teams/releases/tag/v0.6.4
[v0.6.3]: https://github.com/atc0005/send2teams/releases/tag/v0.6.3
[v0.6.2]: https://github.com/atc0005/send2teams/releases/tag/v0.6.2
[v0.6.1]: https://github.com/atc0005/send2teams/releases/tag/v0.6.1
[v0.6.0]: https://github.com/atc0005/send2teams/releases/tag/v0.6.0
[v0.5.0]: https://github.com/atc0005/send2teams/releases/tag/v0.5.0
[v0.4.13]: https://github.com/atc0005/send2teams/releases/tag/v0.4.13
[v0.4.12]: https://github.com/atc0005/send2teams/releases/tag/v0.4.12
[v0.4.11]: https://github.com/atc0005/send2teams/releases/tag/v0.4.11
[v0.4.10]: https://github.com/atc0005/send2teams/releases/tag/v0.4.10
[v0.4.9]: https://github.com/atc0005/send2teams/releases/tag/v0.4.9
[v0.4.8]: https://github.com/atc0005/send2teams/releases/tag/v0.4.8
[v0.4.7]: https://github.com/atc0005/send2teams/releases/tag/v0.4.7
[v0.4.6]: https://github.com/atc0005/send2teams/releases/tag/v0.4.6
[v0.4.5]: https://github.com/atc0005/send2teams/releases/tag/v0.4.5
[v0.4.4]: https://github.com/atc0005/send2teams/releases/tag/v0.4.4
[v0.4.3]: https://github.com/atc0005/send2teams/releases/tag/v0.4.3
[v0.4.2]: https://github.com/atc0005/send2teams/releases/tag/v0.4.2
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
