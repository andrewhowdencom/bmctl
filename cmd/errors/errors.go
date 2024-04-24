package errors

import "errors"

var (
	ErrNoSubcommand = errors.New("no subcommand supplied")
	ErrNotANumber   = errors.New("not able to convert to a number")
)
