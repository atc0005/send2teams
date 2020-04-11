# Copyright 2020 Adam Chalkley
#
# https://github.com/atc0005/send2teams
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

# References:
#
# https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
# https://github.com/mapnik/sphinx-docs/blob/master/Makefile
# https://stackoverflow.com/questions/23843106/how-to-set-child-process-environment-variable-in-makefile
# https://stackoverflow.com/questions/3267145/makefile-execute-another-target
# https://unix.stackexchange.com/questions/124386/using-a-make-rule-to-call-another
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
# https://www.gnu.org/software/make/manual/html_node/Recipe-Syntax.html#Recipe-Syntax
# https://www.gnu.org/software/make/manual/html_node/Special-Variables.html#Special-Variables
# https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html
# https://gist.github.com/subfuzion/0bd969d08fe0d8b5cc4b23c795854a13
# https://stackoverflow.com/questions/10858261/abort-makefile-if-variable-not-set
# https://stackoverflow.com/questions/38801796/makefile-set-if-variable-is-empty

SHELL = /bin/bash

APPNAME					= send2teams

# What package holds the "version" variable used in branding/version output?
VERSION_VAR_PKG			= $(shell go list .)/internal/config

OUTPUTDIR 				= release_assets

# https://gist.github.com/TheHippo/7e4d9ec4b7ed4c0d7a39839e6800cc16
VERSION 				= $(shell git describe --always --long --dirty)

# The default `go build` process embeds debugging information. Building
# without that debugging information reduces the binary size by around 28%.
BUILDCMD				=	go build -mod=vendor -a -ldflags="-s -w -X $(VERSION_VAR_PKG).version=$(VERSION)"
GOCLEANCMD				=	go clean -mod=vendor ./...
GITCLEANCMD				= 	git clean -xfd
CHECKSUMCMD				=	sha256sum -b

.DEFAULT_GOAL := help

  ##########################################################################
  # Targets will not work properly if a file with the same name is ever
  # created in this directory. We explicitly declare our targets to be phony
  # by making them a prerequisite of the special target .PHONY
  ##########################################################################

.PHONY: help
## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: lintinstall
## lintinstall: install common linting tools
# https://github.com/golang/go/issues/30515#issuecomment-582044819
lintinstall:
	@echo "Installing linting tools"

	@export PATH="${PATH}:$(go env GOPATH)/bin"

	@echo "Explicitly enabling Go modules mode per command"
	(cd; GO111MODULE="on" go get golang.org/x/lint/golint)
	(cd; GO111MODULE="on" go get github.com/golangci/golangci-lint/cmd/golangci-lint)
	(cd; GO111MODULE="on" go get honnef.co/go/tools/cmd/staticcheck)
	@echo "Finished updating linting tools"

.PHONY: linting
## linting: runs common linting checks
# https://stackoverflow.com/a/42510278/903870
linting:
	@echo "Running linting tools ..."

	@echo "Running gofmt ..."

	@test -z $(shell gofmt -l -e .) || (echo "WARNING: gofmt linting errors found" \
		&& gofmt -l -e -d . \
		&& exit 1 )

	@echo "Running go vet ..."
	@go vet ./...

	@echo "Running golint ..."
	@golint -set_exit_status ./...

	@echo "Running golangci-lint using settings from .golangci.yml ..."
	@golangci-lint run

	@echo "Running staticcheck ..."
	@staticcheck ./...

	@echo "Finished running linting checks"

.PHONY: gotests
## gotests: runs go test recursively, verbosely
gotests:
	@echo "Running go tests ..."
	@go test ./...
	@echo "Finished running go tests"

.PHONY: goclean
## goclean: removes local build artifacts, temporary files, etc
goclean:
	@echo "Removing object files and cached files ..."
	@$(GOCLEANCMD)
	@echo "Removing any existing release assets"
	@mkdir -p "$(OUTPUTDIR)"
	@rm -vf $(wildcard ${OUTPUTDIR}/*/*-linux-*)
	@rm -vf $(wildcard ${OUTPUTDIR}/*/*-windows-*)

.PHONY: clean
## clean: alias for goclean
clean: goclean

.PHONY: gitclean
## gitclean: WARNING - recursively cleans working tree by removing non-versioned files
gitclean:
	@echo "Removing non-versioned files ..."
	@$(GITCLEANCMD)

.PHONY: pristine
## pristine: run goclean and gitclean to remove local changes
pristine: goclean gitclean

.PHONY: all
# https://stackoverflow.com/questions/3267145/makefile-execute-another-target
## all: generates assets for Linux distros and Windows
all: clean windows linux
	@echo "Completed all cross-platform builds ..."

.PHONY: windows
## windows: generates assets for Windows systems
windows:
	@echo "Building release assets for windows ..."

	@mkdir -p $(OUTPUTDIR)/$(APPNAME)

	@echo "Building 386 binaries"
	@env GOOS=windows GOARCH=386 $(BUILDCMD) -o $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-windows-386.exe ${PWD}/cmd/$(APPNAME)

	@echo "Building amd64 binaries"
	@env GOOS=windows GOARCH=amd64 $(BUILDCMD) -o $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-windows-amd64.exe ${PWD}/cmd/$(APPNAME)

	@echo "Generating checksum files"
	@$(CHECKSUMCMD) $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-windows-386.exe > $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-windows-386.exe.sha256
	@$(CHECKSUMCMD) $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-windows-amd64.exe > $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-windows-amd64.exe.sha256

	@echo "Completed build tasks for windows"

.PHONY: linux
## linux: generates assets for Linux distros
linux:
	@echo "Building release assets for linux ..."

	@mkdir -p $(OUTPUTDIR)/$(APPNAME)

	@echo "Building 386 binaries"
	@env GOOS=linux GOARCH=386 $(BUILDCMD) -o $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-linux-386 ${PWD}/cmd/$(APPNAME)

	@echo "Building amd64 binaries"
	@env GOOS=linux GOARCH=amd64 $(BUILDCMD) -o $(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-linux-amd64 ${PWD}/cmd/$(APPNAME)

	@echo "Generating checksum files"
	@$(CHECKSUMCMD) "$(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-linux-386" > "$(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-linux-386.sha256"
	@$(CHECKSUMCMD) "$(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-linux-amd64" > "$(OUTPUTDIR)/$(APPNAME)/$(APPNAME)-$(VERSION)-linux-amd64.sha256"

	@echo "Completed build tasks for linux"
