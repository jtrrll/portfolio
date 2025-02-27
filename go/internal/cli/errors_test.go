package cli_test

import (
	"errors"
	"fmt"
	"portfolio/internal/cli"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandLineError(t *testing.T) {
	t.Parallel()
	t.Run("Error()", func(t *testing.T) {
		t.Parallel()
		t.Run("returns the correct error message", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"an error occurred",
				"yikes!",
				"uh oh",
			}
			for _, input := range testCases {
				assert.Equal(t, input, cli.CommandLineError{Message: input}.Error())
			}
		})
		t.Run("formats correctly as an error", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"an error occurred",
				"yikes!",
				"uh oh",
			}
			for _, input := range testCases {
				assert.Equal(t, input, fmt.Sprintf("%v", cli.CommandLineError{Message: input}))
			}
		})
		t.Run("can be wrapped and unwrapped", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"an error occurred",
				"yikes!",
				"uh oh",
			}
			for _, input := range testCases {
				baseError := cli.CommandLineError{Message: input}
				assert.Equal(t, baseError, errors.Unwrap(fmt.Errorf("wrapped error: %w", baseError)))
			}
		})
	})
	t.Run("Code()", func(t *testing.T) {
		t.Parallel()
		t.Run("produces the correct exit code", func(t *testing.T) {
			t.Parallel()
			testCases := []uint8{
				0,
				1,
				2,
			}
			for _, input := range testCases {
				assert.Equal(t, input, cli.CommandLineError{StatusCode: input}.Code())
			}
		})
	})
}
