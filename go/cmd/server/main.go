/*
`server` serves a portfolio website.

Usage:

	server <flag> ...
*/
package main

import (
	"log"
	"net/http"
	"portfolio/internal/server"
)

// main serves a portfolio website.
// Will exit with a non-zero status code upon failure.
func main() {
	server := server.New(
		server.WithPort(8080), // TODO: Read port from CLI flag
		server.WithHandler(NewRouter()),
	)

	// TODO: HTTP -> HTTPS
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
