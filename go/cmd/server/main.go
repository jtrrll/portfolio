/*
`server` runs a portfolio website.

Usage:

	server <flag> ...
*/
package main

import (
	"fmt"
	"os"
	"portfolio/internal/cli"

	"github.com/spf13/pflag"
)

// Serves a portfolio website.
// Will exit with a non-zero exit code upon failure.
func main() {
	flagSet := pflag.NewFlagSet("flags", pflag.ExitOnError)
	if flagSet == nil {
		panic("failed to initialize flag set")
	}
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  server <flag> ...\n\n%s", flagSet.FlagUsages())
	}
	help := flagSet.BoolP("help", "h", false, "Print usage help")
	port := flagSet.Uint16P("port", "p", 8080, "Which port to listen to")
	if flagSet.Parse(os.Args) != nil {
		panic("failed to parse command line flags and arguments")
	}

	if *help {
		flagSet.Usage()
		cli.Exit(nil)
	}

	err := Run(Config{
		Port: *port,
	})
	cli.Exit(err)
}
