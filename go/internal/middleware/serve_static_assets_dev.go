//go:build dev

package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServeStaticAssets(fs http.FileSystem, root string) echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Browse:     true,
		Filesystem: fs,
		Root:       root,
	})
}
