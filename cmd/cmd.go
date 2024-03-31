package cmd

import "errors"

// ErrNoSubcommand just means we expect subcommands in this command
var ErrNoSubcommand = errors.New("no subcommand supplied")
