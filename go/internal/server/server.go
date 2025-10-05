// Package server is an abstraction over a basic HTTP server.
package server

import (
	"net/http"
	"time"
)

type Server struct {
	http.Server
}

// New creates a new HTTP server by sequentially applying the given options.
func New(options ...func(*Server)) *Server {
	server := &Server{
		http.Server{
			DisableGeneralOptionsHandler: true,

			// TODO: Use telemetry to determine safe timeout values.
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       30 * time.Second,
		},
	}
	for _, opt := range options {
		opt(server)
	}
	return server
}
