/*
`server` serves a portfolio website.

Usage:

	server <flag> ...
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jtrrll/portfolio/internal/server"
	"github.com/urfave/cli/v3"
)

// main serves a portfolio website.
// Will exit with a non-zero status code upon failure.
func main() {
	cmd := &cli.Command{
		Name:  "portfolio-server",
		Usage: "Serve jtrrll's portfolio website",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			server := server.New(
				server.WithPort(cmd.Uint("port")),
				server.WithHandler(NewRouter(cmd.Bool("trust-proxy"))),
			)
			// TODO: Add terminal output that prints a welcome message
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}
			// TODO: Add a safe shutdown message
			return nil
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
