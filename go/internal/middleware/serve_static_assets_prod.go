//go:build !dev

package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServeStaticAssets(fs http.FileSystem, root string) echo.MiddlewareFunc {
	serveFs := middleware.StaticWithConfig(middleware.StaticConfig{
		Browse:     true,
		Filesystem: fs,
		Root:       root,
	})
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "public, max-age=86400")
			return serveFs(next)(c)
		}
	}
}
