# send2teams

- [send2teams](#send2teams)
  - [Overview](#overview)
  - [Features](#features)
  - [Requirements](#requirements)
  - [How to install it](#how-to-install-it)
  - [License](#license)
  - [References](#references)

## Overview

Small CLI tool used to submit messages to Microsoft Teams. `send2teams` is
intended for use by Nagios, scripts or other scripted actions that may need to
submit pass/fail results to a MS Teams channel.

Many thanks to the developer/contributors of the `dasrick/go-teams-notify`
package for making this tool possible.

## Features

- single binary, no outside dependencies
- minimal configuration
- very few build dependencies

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
     - `go build`
   - for all supported platforms
      - `make all`
   - for Windows
      - `make windows`
   - for Linux
     - `make linux`
1. Copy the applicable binary to whatever systems needs to run it
   1. Linux: `/tmp/send2teams/send2teams`
   1. Windows: `/tmp/send2teams/send2teams.exe`

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

- <https://github.com/dasrick/go-teams-notify/>
  - the package/library this app depends on

- <https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference>
- <https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference>
- <https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/outgoingwebhook>
- <https://docs.microsoft.com/en-us/outlook/actionable-messages/send-via-connectors>
- <https://messagecardplayground.azurewebsites.net/>

- <https://sankalpit.com/how-to-get-channel-webhook-url/>
