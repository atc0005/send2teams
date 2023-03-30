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
# https://stackoverflow.com/questions/1909188/define-make-variable-at-rule-execution-time

SHELL := /bin/bash

# Space-separated list of cmd/BINARY_NAME directories to build
WHAT 					:= send2teams

PROJECT_NAME			:= send2teams

# What package holds the "version" variable used in branding/version output?
# VERSION_VAR_PKG			= $(shell go list -m)
VERSION_VAR_PKG			:= $(shell go list -m)/internal/config
# VERSION_VAR_PKG			= main

OUTPUTDIR 				:= release_assets

ASSETS_PATH				:= $(CURDIR)/$(OUTPUTDIR)

PROJECT_DIR				:= $(CURDIR)

# https://gist.github.com/TheHippo/7e4d9ec4b7ed4c0d7a39839e6800cc16
# VERSION 				:= $(shell git describe --always --long --dirty)

# Use https://github.com/choffmeister/git-describe-semver to generate
# semantic version compatible tag values for use as image suffix.
#
# Attempt to use environment variable. This is set within GitHub Actions
# Workflows, but not via local Makefile use. If environment variable is not
# set, attempt to use local installation of choffmeister/git-describe-semver
# tool.
#
# The build image used by this project for release builds includes this tool,
# so we are covered there.
#
# If this tool is not already installed locally and someone runs this Makefile
# we use a "YOU_NEED_TO_*" placeholder string as a breadcrumb indicating what
# needs to be done to resolve the issue.
#
# This is a VERY expensive operation as it is expanded on every use (due to
# the ?= assignment operator)
# REPO_VERSION              ?= $(shell git-describe-semver --fallback 'v0.0.0' 2>/dev/null || echo YOU_NEED_TO_RUN_depsinstall_makefile_recipe)
#
ifeq ($(origin REPO_VERSION), undefined)
# This is an expensive operation on systems where tools like Cisco AMP are
# present/active.
REPO_VERSION 				:= $(shell git-describe-semver --fallback 'v0.0.0' 2>/dev/null || echo YOU_NEED_TO_RUN_depsinstall_makefile_recipe)

# https://make.mad-scientist.net/deferred-simple-variable-expansion/
# https://stackoverflow.com/questions/50771834/how-to-get-at-most-once-semantics-in-variable-assignments
# REPO_VERSION = $(eval REPO_VERSION := $$(shell git-describe-semver --fallback 'v0.0.0' 2>/dev/null || echo YOU_NEED_TO_RUN_depsinstall_makefile_recipe))$(REPO_VERSION)

endif

DEB_X64_STABLE_PKG_FILE	:= $(PROJECT_NAME)-$(REPO_VERSION)_amd64.deb
RPM_X64_STABLE_PKG_FILE	:= $(PROJECT_NAME)-$(REPO_VERSION).x86_64.rpm

DEB_X64_DEV_PKG_FILE	:= $(PROJECT_NAME)-dev-$(REPO_VERSION)_amd64.deb
RPM_X64_DEV_PKG_FILE	:= $(PROJECT_NAME)-dev-$(REPO_VERSION).x86_64.rpm

# Used when generating download URLs when building assets for public release.
# If the current commit doesn't match an existing tag an error is emitted. We
# toss that error and use a placeholder value.
RELEASE_TAG 			:= $(shell git describe --exact-match --tags 2>/dev/null || echo PLACEHOLDER)

# TESTING purposes
#RELEASE_TAG 			:= 0.11.0

BASE_URL				:= https://github.com/atc0005/$(PROJECT_NAME)/releases/download

ALL_DOWNLOAD_LINKS_FILE	:= $(ASSETS_PATH)/$(PROJECT_NAME)-$(RELEASE_TAG)-all-links.txt
PKG_DOWNLOAD_LINKS_FILE	:= $(ASSETS_PATH)/$(PROJECT_NAME)-$(RELEASE_TAG)-pkgs-links.txt

# Exported so that nFPM can reference it when generating packages.
export SEMVER  := $(REPO_VERSION)

# The default `go build` process embeds debugging information. Building
# without that debugging information reduces the binary size by around 28%.
#
# We also include additional flags in an effort to generate static binaries
# that do not have external dependencies. As of Go 1.15 this still appears to
# be a mixed bag, so YMMV.
#
# See https://github.com/golang/go/issues/26492 for more information.
#
# -s
#	Omit the symbol table and debug information.
#
# -w
#	Omit the DWARF symbol table.
#
# -tags 'osusergo,netgo'
#	Use pure Go implementation of user and group id/name resolution.
#	Use pure Go implementation of DNS resolver.
#
# -extldflags '-static'
#	Pass 'static' flag to external linker.
#
# -linkmode=external
#	https://golang.org/src/cmd/cgo/doc.go
#
#   NOTE: Using external linker requires installation of `gcc-multilib`
#   package when building 32-bit binaries on a Debian/Ubuntu system. It also
#   seems to result in an unstable build that crashes on startup. This *might*
#   be specific to the WSL environment used for builds, but since this is a
#   new issue and and I do not yet know much about this option, I am leaving
#   it out.
#
# CGO_ENABLED=0
#	https://golang.org/cmd/cgo/
#	explicitly disable use of cgo
#	removes potential need for linkage against local c library (e.g., glibc)
BUILDCMD				:=	CGO_ENABLED=0 go build -mod=vendor -trimpath -a -ldflags "-s -w -X $(VERSION_VAR_PKG).version=$(REPO_VERSION)"
QUICK_BUILDCMD			:=	go build -mod=vendor
GOCLEANCMD				:=	go clean -mod=vendor ./...
GITCLEANCMD				:= 	git clean -xfd
CHECKSUMCMD				:=	sha256sum -b
COMPRESSCMD				:= xz --compress --threads=0 --stdout

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

	@echo "Installing latest stable staticcheck version via go install command ..."
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck --version

	@echo Installing latest stable golangci-lint version per official installation script ...
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin
	golangci-lint --version

	@echo "Finished updating linting tools"

.PHONY: linting
## linting: runs common linting checks
linting:
	@echo "Running linting tools ..."

	@echo "Running go vet ..."
	@go vet -mod=vendor $(shell go list -mod=vendor ./... | grep -v /vendor/)

	@echo "Running golangci-lint ..."
	@golangci-lint --version
	@golangci-lint run

	@echo "Running staticcheck ..."
	@staticcheck --version
	@staticcheck $(shell go list -mod=vendor ./... | grep -v /vendor/)

	@echo "Finished running linting checks"

.PHONY: gotests
## gotests: runs go test recursively, verbosely
gotests:
	@echo "Running go tests ..."
	@go test -mod=vendor ./...
	@echo "Finished running go tests"

.PHONY: goclean
## goclean: removes local build artifacts, temporary files, etc
goclean:
	@echo "Removing object files and cached files ..."
	@$(GOCLEANCMD)
	@echo "Removing any existing release assets"
	@mkdir -p "$(ASSETS_PATH)"
	@rm -vf $(wildcard $(ASSETS_PATH)/*/*-linux-*)
	@rm -vf $(wildcard $(ASSETS_PATH)/*/*-windows-*)
	@rm -vf $(wildcard $(ASSETS_PATH)/packages/*/*.rpm)
	@rm -vf $(wildcard $(ASSETS_PATH)/packages/*/*.rpm.sha256)
	@rm -vf $(wildcard $(ASSETS_PATH)/packages/*/*.deb)
	@rm -vf $(wildcard $(ASSETS_PATH)/packages/*/*.deb.sha256)
	@rm -vf $(wildcard $(ASSETS_PATH)/*-links.txt)
	@rm -vf $(wildcard $(PROJECT_DIR)/cmd/*/*.syso)

	@echo "Removing any existing quick build release assets"
	@for target in $(WHAT); do \
		rm -vf $(ASSETS_PATH)/$${target}/$${target}; \
	done

	@echo "Removing any empty asset build paths"
	@find ${ASSETS_PATH} -mindepth 1 -type d -empty -delete

.PHONY: clean
## clean: alias for goclean
clean: goclean

.PHONY: clean-linux-x64-dev
## clean-linux-x64-dev: removes dev assets for Linux x64 distros
clean-linux-x64-dev:
	@echo "Removing dev build assets used to generate dev packages"
	@for target in $(WHAT); do \
		rm -vf $(ASSETS_PATH)/$${target}/$$target-linux-amd64-dev; \
	done

.PHONY: gitclean
## gitclean: WARNING - recursively cleans working tree by removing non-versioned files
gitclean:
	@echo "Removing non-versioned files ..."
	@$(GITCLEANCMD)

.PHONY: pristine
## pristine: run goclean and gitclean to remove local changes
pristine: goclean gitclean

.PHONY: depsinstall
## depsinstall: install or update common build dependencies
depsinstall:
	@echo "Installing current version of build dependencies"

	@export PATH="${PATH}:$(go env GOPATH)/bin"

	@echo "Installing latest go-winres version ..."
	go install github.com/tc-hib/go-winres@latest

	@echo "Installing latest nFPM version ..."
	go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
	nfpm --version

	@echo "Installing latest git-describe-semver version ..."
	go install github.com/choffmeister/git-describe-semver@latest

	@echo "Finished installing or updating build dependencies"

.PHONY: all
# https://stackoverflow.com/questions/3267145/makefile-execute-another-target
## all: generates assets for Linux distros and Windows
all: clean windows linux
	@echo "Completed all cross-platform builds ..."

.PHONY: quick
## quick: generates non-release binaries for current platform, arch
quick:
	@echo "Building non-release assets for current platform, arch ..."

	@set -e; for target in $(WHAT); do \
		mkdir -p $(ASSETS_PATH)/$${target} && \
		echo "  building $${target} binary" && \
		$(QUICK_BUILDCMD) -o $(ASSETS_PATH)/$${target}/$${target} $(PROJECT_DIR)/cmd/$${target}; \
	done

	@echo "Completed tasks for quick build"

.PHONY: windows-x86-build
## windows-x86-build: builds assets for Windows x86 systems
windows-x86-build:
	@echo "Building release assets for windows x86 ..."

	@set -e; for target in $(WHAT); do \
		mkdir -p $(ASSETS_PATH)/$$target && \
		echo "  running go generate for $$target 386 binary ..." && \
		cd $(PROJECT_DIR)/cmd/$$target && \
		env GOOS=windows GOARCH=386 go generate && \
		cd $(PROJECT_DIR) && \
		echo "  building $$target 386 binary" && \
		env GOOS=windows GOARCH=386 $(BUILDCMD) -o $(ASSETS_PATH)/$$target/$$target-windows-386.exe $(PROJECT_DIR)/cmd/$$target; \
	done

	@echo "Completed build tasks for windows x86"

.PHONY: windows-x86-compress
## windows-x86-compress: compresses generated Windows x86 assets
windows-x86-compress:
	@echo "Compressing release assets for windows x86 ..."

	@set -e; for target in $(WHAT); do \
		echo "  compressing $$target 386 binary" && \
		$(COMPRESSCMD) $(ASSETS_PATH)/$$target/$$target-windows-386.exe > \
			$(ASSETS_PATH)/$$target/$$target-windows-386.exe.xz && \
		rm -f $(ASSETS_PATH)/$$target/$$target-windows-386.exe; \
	done

	@echo "Completed compress tasks for windows x86"

.PHONY: windows-x86-checksums
## windows-x86-checksums: generates checksum files for Windows x86 assets
windows-x86-checksums:
	@echo "Generating checksum files for windows x86 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target checksum file" && \
		cd $(ASSETS_PATH)/$$target && \
		$(CHECKSUMCMD) $$target-windows-386.exe.xz > $$target-windows-386.exe.xz.sha256 && \
		cd $$OLDPWD; \
	done

	@echo "Completed generation of checksum files for windows x86"

.PHONY: windows-x86-links
## windows-x86-links: generates download URLs for Windows x86 assets
windows-x86-links:
	@echo "Generating download links for windows x86 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target download links" && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-windows-386.exe.xz" >> $(ALL_DOWNLOAD_LINKS_FILE) && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-windows-386.exe.xz.sha256" >> $(ALL_DOWNLOAD_LINKS_FILE); \
	done

	@echo "Completed generating download links for windows x86 assets"

.PHONY: windows-x64-build
## windows-x64-build: builds assets for Windows x64 systems
windows-x64-build:
	@echo "Building release assets for windows x64 ..."

	@set -e; for target in $(WHAT); do \
		mkdir -p $(ASSETS_PATH)/$$target && \
		echo "  running go generate for $$target amd64 binary ..." && \
		cd $(PROJECT_DIR)/cmd/$$target && \
		env GOOS=windows GOARCH=amd64 go generate && \
		cd $(PROJECT_DIR) && \
		echo "  building $$target amd64 binary" && \
		env GOOS=windows GOARCH=amd64 $(BUILDCMD) -o $(ASSETS_PATH)/$$target/$$target-windows-amd64.exe $(PROJECT_DIR)/cmd/$$target; \
	done

	@echo "Completed build tasks for windows x64"

.PHONY: windows-x64-compress
## windows-x64-compress: compresses generated Windows x64 assets
windows-x64-compress:
	@echo "Compressing release assets for windows x64 ..."

	@set -e; for target in $(WHAT); do \
		echo "  compressing $$target amd64 binary" && \
		$(COMPRESSCMD) $(ASSETS_PATH)/$$target/$$target-windows-amd64.exe > \
			$(ASSETS_PATH)/$$target/$$target-windows-amd64.exe.xz && \
		rm -f $(ASSETS_PATH)/$$target/$$target-windows-amd64.exe; \
	done

	@echo "Completed compress tasks for windows x64"

.PHONY: windows-x64-checksums
## windows-x64-checksums: generates checksum files for Windows x64 assets
windows-x64-checksums:
	@echo "Generating checksum files for windows x64 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target checksum file" && \
		cd $(ASSETS_PATH)/$$target && \
		$(CHECKSUMCMD) $$target-windows-amd64.exe.xz > $$target-windows-amd64.exe.xz.sha256 && \
		cd $$OLDPWD; \
	done

	@echo "Completed generation of checksum files for windows x64"

.PHONY: windows-x64-links
## windows-x64-links: generates download URLs for Windows x64 assets
windows-x64-links:
	@echo "Generating download links for windows x64 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target download links" && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-windows-amd64.exe.xz" >> $(ALL_DOWNLOAD_LINKS_FILE) && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-windows-amd64.exe.xz.sha256" >> $(ALL_DOWNLOAD_LINKS_FILE); \
	done

	@echo "Completed generating download links for windows x64 assets"

.PHONY: windows-x86
## windows-x86: generates assets for Windows x86
windows-x86: windows-x86-build windows-x86-compress windows-x86-checksums
	@echo "Completed all tasks for windows x86"

.PHONY: windows-x64
## windows-x64: generates assets for Windows x64
windows-x64: windows-x64-build windows-x64-compress windows-x64-checksums
	@echo "Completed all tasks for windows x64"

.PHONY: windows
## windows: generates assets for Windows x86 and x64 systems
windows: windows-x86 windows-x64
	@echo "Completed all tasks for windows"

.PHONY: windows-links
## windows-links: generates download URLs for Windows x86 and x64 assets
windows-links: windows-x86-links windows-x64-links
	@echo "Completed generating download links for windows x86 and x64 assets"

.PHONY: linux-x86-build
## linux-x86-build: builds assets for Linux x86 distros
linux-x86-build:
	@echo "Building release assets for linux x86 ..."

	@set -e; for target in $(WHAT); do \
		mkdir -p $(ASSETS_PATH)/$$target && \
		echo "  building $$target 386 binary" && \
		env GOOS=linux GOARCH=386 $(BUILDCMD) -o $(ASSETS_PATH)/$$target/$$target-linux-386 $(PROJECT_DIR)/cmd/$$target; \
	done

	@echo "Completed build tasks for linux x86"

.PHONY: linux-x86-compress
## linux-x86-compress: compresses generated Linux x86 assets
linux-x86-compress:
	@echo "Compressing release assets for linux x86 ..."

	@set -e; for target in $(WHAT); do \
		echo "  compressing $$target 386 binary" && \
		$(COMPRESSCMD) $(ASSETS_PATH)/$$target/$$target-linux-386 > \
			$(ASSETS_PATH)/$$target/$$target-linux-386.xz && \
		rm -f $(ASSETS_PATH)/$$target/$$target-linux-386; \
	done

	@echo "Completed compress tasks for linux x86"

.PHONY: linux-x86-checksums
## linux-x86-checksums: generates checksum files for Linux x86 assets
linux-x86-checksums:
	@echo "Generating checksum files for linux x86 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target checksum file" && \
		cd $(ASSETS_PATH)/$$target && \
		$(CHECKSUMCMD) $$target-linux-386.xz > $$target-linux-386.xz.sha256 && \
		cd $$OLDPWD; \
	done

	@echo "Completed generation of checksum files for linux x86"

.PHONY: linux-x86-links
## linux-x86-links: generates download URLs for Linux x86 assets
linux-x86-links:
	@echo "Generating download links for linux x86 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target download links" && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-linux-386.xz" >> $(ALL_DOWNLOAD_LINKS_FILE) && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-linux-386.xz.sha256" >> $(ALL_DOWNLOAD_LINKS_FILE); \
	done

	@echo "Completed generating download links for linux x86 assets"

.PHONY: linux-x64-build
## linux-x64-build: builds assets for Linux x64 distros
linux-x64-build:
	@echo "Building release assets for linux x64 ..."

	@set -e; for target in $(WHAT); do \
		mkdir -p $(ASSETS_PATH)/$$target && \
		echo "  building $$target amd64 binary" && \
		env GOOS=linux GOARCH=amd64 $(BUILDCMD) -o $(ASSETS_PATH)/$$target/$$target-linux-amd64 $(PROJECT_DIR)/cmd/$$target; \
	done

	@echo "Completed build tasks for linux x64"

.PHONY: linux-x64-compress
## linux-x64-compress: compresses generated Linux x64 assets
linux-x64-compress:
	@echo "Compressing release assets for linux x64 ..."

	@set -e; for target in $(WHAT); do \
		echo "  compressing $$target amd64 binary" && \
		$(COMPRESSCMD) $(ASSETS_PATH)/$$target/$$target-linux-amd64 > \
			$(ASSETS_PATH)/$$target/$$target-linux-amd64.xz && \
		rm -f $(ASSETS_PATH)/$$target/$$target-linux-amd64; \
	done

	@echo "Completed compress tasks for linux x64"

.PHONY: linux-x64-checksums
## linux-x64-checksums: generates checksum files for Linux x64 assets
linux-x64-checksums:
	@echo "Generating checksum files for linux x64 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target checksum file" && \
		cd $(ASSETS_PATH)/$$target && \
		$(CHECKSUMCMD) $$target-linux-amd64.xz > $$target-linux-amd64.xz.sha256 && \
		cd $$OLDPWD; \
	done

.PHONY: linux-x64-links
## linux-x64-links: generates download URLs for Linux x64 assets
linux-x64-links:
	@echo "Generating download links for linux x64 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  Generating $$target download links" && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-linux-amd64.xz" >> $(ALL_DOWNLOAD_LINKS_FILE) && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-linux-amd64.xz.sha256" >> $(ALL_DOWNLOAD_LINKS_FILE); \
	done

	@echo "Completed generating download links for linux x64 assets"

.PHONY: linux-x64-dev-build
## linux-x64-dev-build: builds dev assets for Linux x64 distros
linux-x64-dev-build:
	@echo "Building dev assets for linux x64 ..."

	@set -e; for target in $(WHAT); do \
		mkdir -p $(ASSETS_PATH)/$$target && \
		echo "  building $$target amd64 binary" && \
		env GOOS=linux GOARCH=amd64 $(BUILDCMD) -o $(ASSETS_PATH)/$$target/$$target-linux-amd64-dev $(PROJECT_DIR)/cmd/$$target; \
	done

	@echo "Completed dev build tasks for linux x64"

.PHONY: linux-x64-dev-compress
## linux-x64-dev-compress: compresses generated dev Linux x64 assets
linux-x64-dev-compress:
	@echo "Compressing dev assets for linux x64 ..."

	@set -e; for target in $(WHAT); do \
		echo "  compressing $$target amd64 binary" && \
		$(COMPRESSCMD) $(ASSETS_PATH)/$$target/$$target-linux-amd64-dev > \
			$(ASSETS_PATH)/$$target/$$target-linux-amd64-dev.xz && \
		rm -f $(ASSETS_PATH)/$$target/$$target-linux-amd64-dev; \
	done

	@echo "Completed dev compress tasks for linux x64"

.PHONY: linux-x64-dev-checksums
## linux-x64-dev-checksums: generates checksum files for dev Linux x64 assets
linux-x64-dev-checksums:
	@echo "Generating checksum files for dev linux x64 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  generating $$target checksum file" && \
		cd $(ASSETS_PATH)/$$target && \
		$(CHECKSUMCMD) $$target-linux-amd64-dev.xz > $$target-linux-amd64-dev.xz.sha256 && \
		cd $$OLDPWD; \
	done

.PHONY: linux-x64-dev-links
## linux-x64-dev-links: generates download URLs for dev Linux x64 assets
linux-x64-dev-links:
	@echo "Generating download links for dev linux x64 assets ..."

	@set -e; for target in $(WHAT); do \
		echo "  Generating $$target download links" && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-linux-amd64-dev.xz" >> $(ALL_DOWNLOAD_LINKS_FILE) && \
		echo "$(BASE_URL)/$(RELEASE_TAG)/$$target-linux-amd64-dev.xz.sha256" >> $(ALL_DOWNLOAD_LINKS_FILE); \
	done

	@echo "Completed generating download links for dev linux x64 assets"

.PHONY: linux-x86
## linux-x86: generates assets for Linux x86
linux-x86: linux-x86-build linux-x86-compress linux-x86-checksums
	@echo "Completed all tasks for linux x86"

.PHONY: linux-x64
## linux-x64: generates assets for Linux x64
linux-x64: linux-x64-build linux-x64-compress linux-x64-checksums
	@echo "Completed all tasks for linux x64"

.PHONY: linux
## linux: generates assets for Linux x86 and x64 distros
linux: linux-x86 linux-x64
	@echo "Completed all tasks for linux"

.PHONY: linux-links
## linux-links: generates download URLs for Linux x86 and x64 assets
linux-links: linux-x86-links linux-x64-links
	@echo "Completed generating download links for linux x86 and x64 assets"

.PHONY: packages-stable
## packages-stable: generates "stable" release series DEB and RPM packages
packages-stable: linux-x64-build

	@echo
	@echo Generating stable release series packages ...

	@mkdir -p $(ASSETS_PATH)/packages/stable

	@echo
	@echo "  - stable DEB package ..."
	@cd $(PROJECT_DIR)/packages/stable && \
		nfpm package --config nfpm.yaml --packager deb --target $(ASSETS_PATH)/packages/stable/$(DEB_X64_STABLE_PKG_FILE)

	@echo
	@echo "  - stable RPM package ..."
	@cd $(PROJECT_DIR)/packages/stable && \
		nfpm package --config nfpm.yaml --packager rpm --target $(ASSETS_PATH)/packages/stable/$(RPM_X64_STABLE_PKG_FILE)

	@echo
	@echo "Generating checksum files ..."

	@echo "  - stable DEB package checksum file"
	@set -e ;\
		cd $(ASSETS_PATH)/packages/stable && \
		for file in $$(find . -name "*.deb" -printf '%P\n'); do \
			$(CHECKSUMCMD) $${file} > $${file}.sha256 ; \
		done

	@echo "  - stable RPM package checksum file"
	@set -e ;\
		cd $(ASSETS_PATH)/packages/stable && \
		for file in $$(find . -name "*.rpm" -printf '%P\n'); do \
			$(CHECKSUMCMD) $${file} > $${file}.sha256 ; \
		done

	@echo
	@echo "Completed package build tasks"

.PHONY: packages-dev
## packages-dev: generates "dev" release series DEB and RPM packages
packages-dev: linux-x64-dev-build

	@echo
	@echo Generating dev release series packages ...

	@mkdir -p $(ASSETS_PATH)/packages/dev

	@echo
	@echo "  - dev DEB package ..."
	@cd $(PROJECT_DIR)/packages/dev && \
		nfpm package --config nfpm.yaml --packager deb --target $(ASSETS_PATH)/packages/dev/$(DEB_X64_DEV_PKG_FILE)

	@echo
	@echo "  - dev RPM package ..."
	@cd $(PROJECT_DIR)/packages/dev && \
		nfpm package --config nfpm.yaml --packager rpm --target $(ASSETS_PATH)/packages/dev/$(RPM_X64_DEV_PKG_FILE)

	@echo
	@echo "Generating checksum files ..."

	@echo "  - dev DEB package checksum file"
	@set -e ;\
		cd $(ASSETS_PATH)/packages/dev && \
		for file in $$(find . -name "*.deb" -printf '%P\n'); do \
			$(CHECKSUMCMD) $${file} > $${file}.sha256 ; \
		done

	@echo "  - dev RPM package checksum file"
	@set -e ;\
		cd $(ASSETS_PATH)/packages/dev && \
		for file in $$(find . -name "*.rpm" -printf '%P\n'); do \
			$(CHECKSUMCMD) $${file} > $${file}.sha256 ; \
		done

	@echo
	@echo "Completed dev release package build tasks"

.PHONY: packages
## packages: generates "dev" and "stable" release series DEB and RPM packages
packages: packages-dev packages-stable
	@echo "Completed all package build tasks"

.PHONY: package-links
## package-links: generates download URLs for package assets
package-links:

	@echo "Generating download links for package assets ..."

	@echo "  - DEB package download links"
	@if [ -d $(ASSETS_PATH)/packages/dev ]; then \
		cd $(ASSETS_PATH)/packages/dev && \
		for file in $$(find . -name "*.deb" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi
	@if [ -d $(ASSETS_PATH)/packages/dev ]; then \
		cd $(ASSETS_PATH)/packages/dev && \
		for file in $$(find . -name "*.deb.sha256" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi

	@if [ -d $(ASSETS_PATH)/packages/stable ]; then \
		cd $(ASSETS_PATH)/packages/stable && \
		for file in $$(find . -name "*.deb" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi

	@if [ -d $(ASSETS_PATH)/packages/stable ]; then \
		cd $(ASSETS_PATH)/packages/stable && \
		for file in $$(find . -name "*.deb.sha256" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi

	@echo "  - RPM package download links"
	@if [ -d $(ASSETS_PATH)/packages/dev ]; then \
		cd $(ASSETS_PATH)/packages/dev && \
		for file in $$(find . -name "*.rpm" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi
	@if [ -d $(ASSETS_PATH)/packages/dev ]; then \
		cd $(ASSETS_PATH)/packages/dev && \
		for file in $$(find . -name "*.rpm.sha256" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi

	@if [ -d $(ASSETS_PATH)/packages/stable ]; then \
		cd $(ASSETS_PATH)/packages/stable && \
		for file in $$(find . -name "*.rpm" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi

	@if [ -d $(ASSETS_PATH)/packages/stable ]; then \
		cd $(ASSETS_PATH)/packages/stable && \
		for file in $$(find . -name "*.rpm.sha256" -printf '%P\n'); do \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(PKG_DOWNLOAD_LINKS_FILE) && \
			echo "$(BASE_URL)/$(RELEASE_TAG)/$${file}" >> $(ALL_DOWNLOAD_LINKS_FILE); \
		done; \
	fi

	@echo "Completed generating download links for package assets"

.PHONY: links
## links: generates download URLs for release assets
links: windows-x86-links windows-x64-links linux-x86-links linux-x64-links package-links
	@echo "Completed generating download links for all release assets"

.PHONY: dev-build
## dev-build: generates dev build assets for public release
dev-build: clean linux-x64-dev-build packages-dev package-links linux-x64-dev-compress linux-x64-dev-checksums linux-x64-dev-links
	@echo "Completed all tasks for dev release build"

.PHONY: release-build
## release-build: generates stable build assets for public release
release-build: clean windows linux-x86 packages-dev clean-linux-x64-dev packages-stable linux-x64-compress linux-x64-checksums links
	@echo "Completed all tasks for stable release build"

.PHONY: helper-builder-setup
helper-builder-setup:

	@echo "Beginning regeneration of builder image using $(CONTAINER_COMMAND)"
	@echo "Removing any previous build image"
	$(CONTAINER_COMMAND) image prune --all --force --filter "label=atc0005_projects_builder_image"

	@echo "Gathering $(CONTAINER_COMMAND) build environment details"
	@$(CONTAINER_COMMAND) version

	@echo
	@echo "Generating release builder image"
	$(CONTAINER_COMMAND) image build \
		--pull \
		--no-cache \
		--force-rm \
		. \
		-f dependabot/docker/builds/Dockerfile \
		-t builder_image \
		--label="atc0005_projects_builder_image"
	@echo "Completed generation of release builder image"

	@echo "Listing current container images managed by $(CONTAINER_COMMAND)"
	$(CONTAINER_COMMAND) image ls

	@echo
	@echo "Inspecting release builder image environment using $(CONTAINER_COMMAND)"
	@$(CONTAINER_COMMAND) inspect --format "{{range .Config.Env}}{{println .}}{{end}}" builder_image

	@echo "Completed regeneration of builder image using $(CONTAINER_COMMAND)"

	@echo "Prepare output path for generated assets"
	@mkdir -p $(ASSETS_PATH)

.PHONY: docker-release-build
## docker-release-build: generates stable build assets for public release using docker container
docker-release-build: CONTAINER_COMMAND := docker
docker-release-build: clean helper-builder-setup

	@echo "Beginning release build using $(CONTAINER_COMMAND)"

	@echo
	@echo "Using release builder image to generate project release assets"
	$(CONTAINER_COMMAND) container run \
		--user builduser:builduser \
		--rm \
		-i \
		-v $$PWD/$(OUTPUTDIR):/builds/$(OUTPUTDIR):rw \
		-w /builds \
		builder_image \
		make release-build

	@echo "Completed release build using $(CONTAINER_COMMAND)"

.PHONY: podman-release-build
## podman-release-build: generates stable build assets for public release using podman container
podman-release-build: CONTAINER_COMMAND := podman
podman-release-build: clean helper-builder-setup

	@echo "Beginning release build using $(CONTAINER_COMMAND)"

	@echo
	@echo "Using release builder image to generate project release assets"
	$(CONTAINER_COMMAND) container run \
		--rm \
		-i \
		-v $$PWD/$(OUTPUTDIR):/builds/$(OUTPUTDIR):rw \
		-w /builds \
		builder_image \
		make release-build

	@echo "Completed release build using $(CONTAINER_COMMAND)"

.PHONY: docker-dev-build
## docker-dev-build: generates dev build assets for public release using docker container
docker-dev-build: CONTAINER_COMMAND := docker
docker-dev-build: clean helper-builder-setup

	@echo "Beginning dev build using $(CONTAINER_COMMAND)"

	@echo
	@echo "Using release builder image to generate project release assets"
	$(CONTAINER_COMMAND) container run \
		--user builduser:builduser \
		--rm \
		-i \
		-v $$PWD/$(OUTPUTDIR):/builds/$(OUTPUTDIR):rw \
		-w /builds \
		builder_image \
		make dev-build

	@echo "Completed dev build using $(CONTAINER_COMMAND)"

.PHONY: podman-dev-build
## podman-dev-build: generates dev build assets for public release using podman container
podman-dev-build: CONTAINER_COMMAND := podman
podman-dev-build: clean helper-builder-setup

	@echo "Beginning dev build using $(CONTAINER_COMMAND)"

	@echo
	@echo "Using release builder image to generate project release assets"
	$(CONTAINER_COMMAND) container run \
		--rm \
		-i \
		-v $$PWD/$(OUTPUTDIR):/builds/$(OUTPUTDIR):rw \
		-w /builds \
		builder_image \
		make dev-build

	@echo "Completed dev build using $(CONTAINER_COMMAND)"

.PHONY: docker-packages
## docker-packages: generates dev and stable packages using builder image
docker-packages: CONTAINER_COMMAND := docker
docker-packages: helper-builder-setup

	@echo "Beginning package generation using $(CONTAINER_COMMAND)"

	@echo
	@echo "Using release builder image to generate packages"

	@echo "Building with $(CONTAINER_COMMAND)"
	$(CONTAINER_COMMAND) container run \
		--rm \
		--user builduser:builduser \
		-i \
		-v $$PWD/$(OUTPUTDIR):/builds/$(OUTPUTDIR):rw \
		-w /builds \
		builder_image \
		make packages

	@echo "Completed package generation using $(CONTAINER_COMMAND)"

.PHONY: podman-packages
## podman-packages: generates dev and stable packages using podman container
podman-packages: CONTAINER_COMMAND := podman
podman-packages: helper-builder-setup

	@echo "Beginning package generation using $(CONTAINER_COMMAND)"

	@echo
	@echo "Using release builder image to generate packages"

	@echo "Building with $(CONTAINER_COMMAND)"
	$(CONTAINER_COMMAND) container run \
		--rm \
		-i \
		-v $$PWD/$(OUTPUTDIR):/builds/$(OUTPUTDIR):rw \
		-w /builds \
		builder_image \
		make packages

	@echo "Completed package generation using $(CONTAINER_COMMAND)"
