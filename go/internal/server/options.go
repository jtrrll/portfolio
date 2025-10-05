package server

import (
	"fmt"
	"net/http"
)

// WithHandler configures the server's HTTP request handler.
func WithHandler(handler http.Handler) func(*Server) {
	return func(s *Server) {
		s.Handler = handler
	}
}

// WithPort configures the server to listen on the given port.
func WithPort(port uint) func(*Server) {
	return func(s *Server) {
		s.Addr = fmt.Sprintf(":%d", port)
	}
}
