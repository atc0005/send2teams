#!/bin/bash

# Copyright 2019 Adam Chalkley
#
# https://github.com/atc0005/send2teams
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.


# Purpose: Helper script for installing linting tools used by this project

export PATH=${PATH}:$(go env GOPATH)/bin

# Temporarily disable module-aware mode so that we can install linting tools
# without modifying this project's go.mod and go.sum files
export GO111MODULE="off"
go get -u golang.org/x/lint/golint
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
go get -u honnef.co/go/tools/cmd/staticcheck


# Reset GO111MODULE back to the default
export GO111MODULE=""
