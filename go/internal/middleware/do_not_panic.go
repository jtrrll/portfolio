package middleware

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// DoNotPanic returns a middleware that recovers from panics anywhere in the chain,
// logs the error with slog, records it in a span, and returns a 500 response.
func DoNotPanic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					ctx := c.Request().Context()

					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					span := trace.SpanFromContext(ctx)
					span.RecordError(err, trace.WithStackTrace(true))
					span.SetStatus(codes.Error, err.Error())

					slog.ErrorContext(ctx, "recovered from panic",
						"error", err,
					)

					c.Error(err)
				}
			}()

			return next(c)
		}
	}
}
