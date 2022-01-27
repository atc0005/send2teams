// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"os"
	"testing"

	"github.com/atc0005/send2teams/internal/config"
)

// Setup basic tests to ensure that config initialization works as expected.

func TestConfigInitialization(t *testing.T) {

	// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang

	// Save old command-line arguments so that we can restore them later
	oldArgs := os.Args

	// Defer restoring original command-line arguments
	defer func() { os.Args = oldArgs }()

	// Note to self: Don't add/escape double-quotes here. The shell strips
	// them away and the application never sees them.
	os.Args = []string{
		"/usr/local/bin/send2teams", "--version",
	}

	// based on given CLI args, we *should* have a sentinel error here.
	_, cfgErr := config.NewConfig()

	if !errors.Is(cfgErr, config.ErrVersionRequested) {
		t.Fatalf("got %v; expected error %q", cfgErr, config.ErrVersionRequested)
	} else {
		config.Branding()
		t.Log("--version flag properly recognized")
	}

}
