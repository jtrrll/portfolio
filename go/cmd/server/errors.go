package main

// Common errors.
var (
	errFailedToStart = Error{Code: 1, Message: "failed to start server"}
)

// A command-line error.
type Error struct {
	Message string
	Code    uint8
}

// Returns the error's status code.
func (err Error) ExitCode() uint8 {
	return err.Code
}

// Returns the error's message.
func (err Error) Error() string {
	return err.Message
}
