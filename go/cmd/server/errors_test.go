package main_test

import (
	"errors"
	"fmt"
	main "portfolio/cmd/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
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
				assert.Equal(t, input, main.Error{Message: input}.Error())
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
				assert.Equal(t, input, fmt.Sprintf("%v", main.Error{Message: input}))
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
				baseError := main.Error{Message: input}
				assert.Equal(t, baseError, errors.Unwrap(fmt.Errorf("wrapped error: %w", baseError)))
			}
		})
	})
	t.Run("ExitCode()", func(t *testing.T) {
		t.Parallel()
		t.Run("produces the correct exit code", func(t *testing.T) {
			t.Parallel()
			testCases := []uint8{
				0,
				1,
				2,
			}
			for _, input := range testCases {
				assert.Equal(t, input, main.Error{Code: input}.ExitCode())
			}
		})
	})
}
