// A collection of utilities for building command-line interfaces.
package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// An error with a code.
type ErrorCode interface {
	error
	Code() uint8
}

// Terminates the current program.
// Also prints the given error to stderr if present.
func Exit(err ErrorCode) {
	if err == nil {
		os.Exit(0)
	}

	code := err.Code()
	if code == 0 {
		os.Exit(0)
	}
	if code > 125 {
		panic("invalid exit code")
	}

	fmt.Fprint(os.Stderr, color.RedString("error: %s\n", err.Error()))
	os.Exit(int(code))
}

// A command-line error.
type CommandLineError struct {
	Message    string
	StatusCode uint8
}

// Returns the error's status code.
func (err CommandLineError) Code() uint8 {
	return err.StatusCode
}

// Returns the error's message.
func (err CommandLineError) Error() string {
	return err.Message
}
