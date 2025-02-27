package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"portfolio/internal/cli"
	"time"
)

// A runtime configuration for the server.
type Config struct {
	// The port to listen to.
	Port uint16
}

// The core server process.
func Run(config Config) (err cli.ErrorCode) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	shutdownOtel, setupErr := setupOtelSDK(ctx)
	if setupErr != nil {
		err = errFailedToSetupOtel
		return
	}
	defer func() {
		if shutdownOtel(context.Background()) != nil {
			err = errFailedToShutdownOtel
		}
	}()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Handler:      newHTTPHandler(),
	}
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- server.ListenAndServe()
	}()

	select {
	case <-serverErr:
		err = errFailedToStartServer
		return
	case <-ctx.Done():
		stop()
	}

	if server.Shutdown(context.Background()) != nil {
		err = errFailedToShutdownServer
	}
	return
}
