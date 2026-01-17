package middleware

import (
	"fmt"
	"log/slog"
	"runtime"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("portfolio")

// DoNotPanic returns a middleware that recovers from panics anywhere in the chain,
// logs the error with slog, records it in a span, and returns a 500 response.
func DoNotPanic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, span := tracer.Start(c.Request().Context(), "http.recoverable_request",
				trace.WithSpanKind(trace.SpanKindServer),
			)
			defer span.End()

			// Update the context in the echo.Context
			c.SetRequest(c.Request().WithContext(ctx))

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					// Record error in span
					span.RecordError(err, trace.WithStackTrace(true))
					span.SetStatus(codes.Error, err.Error())

					// Capture stack trace
					stack := make([]byte, 4096)
					length := runtime.Stack(stack, false)
					stackTrace := string(stack[:length])

					// Log the error with slog
					slog.ErrorContext(ctx, "recovered from panic",
						"error", err,
						"stack", stackTrace,
					)

					// Return error through Echo's error handler
					c.Error(err)
				}
			}()

			return next(c)
		}
	}
}
