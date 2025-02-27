// A collection of middleware functions.
package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Adds cache control to the response.
func CacheControl(duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d, public", int(duration.Seconds())))
			next.ServeHTTP(w, r)
		})
	}
}
