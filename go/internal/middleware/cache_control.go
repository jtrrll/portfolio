package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func CacheControl(maxAge time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", uint32(maxAge.Seconds())))
			return next(c)
		}
	}
}
