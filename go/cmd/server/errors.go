package main

import "portfolio/internal/cli"

// Common errors.
var (
	errFailedToSetupOtel      = cli.CommandLineError{StatusCode: 1, Message: "failed to set up OpenTelemetry"}
	errFailedToShutdownOtel   = cli.CommandLineError{StatusCode: 2, Message: "failed to shut down OpenTelemetry"}
	errFailedToStartServer    = cli.CommandLineError{StatusCode: 3, Message: "failed to start server"}
	errFailedToShutdownServer = cli.CommandLineError{StatusCode: 4, Message: "failed to shut down server"}
)
