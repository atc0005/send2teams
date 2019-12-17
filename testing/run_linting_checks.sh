#!/bin/bash

# Copyright 2019 Adam Chalkley
#
# https://github.com/atc0005/send2teams
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.


# Purpose: Run common linters locally to confirm code quality


# Go ahead and append $GOPATH/bin to $PATH in an effort to locate
# the go linters referenced in this script.
export PATH=${PATH}:$(go env GOPATH)/bin

# Assume all is well starting out
final_exit_code=0
failed_app=""

###########################################################
# Run linters
###########################################################


# https://stackoverflow.com/a/42510278/903870
diff -u <(echo -n) <(gofmt -l -e -d .)

status=$?
if [[ $status -ne 0 ]]; then
    final_exit_code=$status
    failed_app="gofmt"
    echo "Non-zero exit code from $failed_app: $status"
fi

go vet ./...

status=$?
if [[ $status -ne 0 ]]; then
    final_exit_code=$status
    failed_app="go vet"
    echo "Non-zero exit code from $failed_app: $status"
fi

if ! which golint > /dev/null; then
cat <<\EOF
Error: Unable to locate "golint"

Install golint with the following command:

make lintinstall

EOF
    exit 1
else
    golint -set_exit_status ./...
fi

# TODO: This might not be needed based on use of "-set_exit_status"
status=$?
if [[ $status -ne 0 ]]; then
    final_exit_code=$status
    failed_app="staticcheck"
    echo "Non-zero exit code from $failed_app: $status"
fi

if ! which golangci-lint > /dev/null; then
cat <<\EOF
Error: Unable to locate "golangci-lint"

Install golangci-lint with the following command:

make lintinstall

EOF
    exit 1
else
    golangci-lint run \
        -E goimports \
        -E gosec \
        -E stylecheck \
        -E goconst \
        -E depguard \
        -E prealloc \
        -E misspell \
        -E maligned \
        -E dupl \
        -E unconvert \
        -E golint \
        -E gocritic
fi

status=$?
if [[ $status -ne 0 ]]; then
    final_exit_code=$status
    failed_app="golangci-lint"
    echo "Non-zero exit code from $failed_app: $status"
fi

if ! which staticcheck > /dev/null; then
cat <<\EOF
Error: Unable to locate "staticcheck"

Install staticcheck with the following command:

make lintinstall

EOF
    exit 1
else
    staticcheck ./...
fi

status=$?
if [[ $status -ne 0 ]]; then
    final_exit_code=$status
    failed_app="staticcheck"
    echo "Non-zero exit code from $failed_app: $status"
fi

# Give feedback on linting failure cause
if [[ $final_exit_code -ne 0 ]]; then
    echo "Linting failed, most recent failure: $failed_app"
fi

exit $final_exit_code
