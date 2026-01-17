/*
`server` serves a portfolio website.

Usage:

	server <flag> ...
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jtrrll/portfolio/internal/server"
	"github.com/urfave/cli/v3"
	"go.opentelemetry.io/contrib/bridges/otelslog"
)

// main serves a portfolio website.
// Will exit with a non-zero status code upon failure.
func main() {
	cmd := &cli.Command{
		Name:  "portfolio-server",
		Usage: "Serve jtrrll's portfolio website",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			defer stop()

			// Set up logger.
			logger := slog.New(otelslog.NewHandler("portfolio"))
			slog.SetDefault(logger)

			// Set up OpenTelemetry.
			shutdownOpenTelemetry, err := setupOpenTelemetry(ctx)
			if err != nil {
				return err
			}
			defer func() {
				err = errors.Join(err, shutdownOpenTelemetry(context.Background()))
			}()

			// Start HTTP server.
			srv := server.New(
				server.WithPort(cmd.Uint("port")),
				server.WithHandler(NewRouter(logger, cmd.Bool("trust-proxy"))),
			)
			srvErr := make(chan error, 1)
			go func() {
				fmt.Printf("Starting server on port %d...\n", cmd.Uint("port"))
				srvErr <- srv.ListenAndServe()
			}()

			// Wait for interruption.
			select {
			case err = <-srvErr:
				// Error when starting HTTP server.
				return err
			case <-ctx.Done():
				// Wait for first CTRL+C.
				// Stop receiving signal notifications as soon as possible.
				stop()
			}

			fmt.Println("Shutting down gracefully...")
			err = srv.Shutdown(context.Background())
			return err
		},
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:    "port",
				Usage:   "Port to listen on",
				Value:   8080,
				Aliases: []string{"p"},
				Action: func(ctx context.Context, cmd *cli.Command, v uint) error {
					if v < 1024 || 65535 < v {
						return fmt.Errorf("invalid port %d (must be between 1024 and 65535)", v)
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:  "trust-proxy",
				Usage: "Trust X-Forwarded-For header from reverse proxy",
				Value: false,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
