package main

import (
	"net/http"
	"strings"

	"embed"

	"portfolio/cmd/server/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static
var staticAssets embed.FS

// NewRouter creates an HTTP request handler.
func NewRouter() http.Handler {
	globalRouter := echo.New()

	// TODO: Add middleware
	globalRouter.Use()

	// TODO: Add middleware and routes
	globalRouter.Group("/api")

	// TODO: Add middleware and routes
	pagesRouter := globalRouter.Group("",
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				if err, ok := next(c).(*echo.HTTPError); ok && err.Code == http.StatusNotFound {
					c.Redirect(http.StatusSeeOther, "/")
				}
				return nil
			}
		},
	)
	pagesRouter.GET("/", pages.Index())
	pagesRouter.GET("/audio", pages.Audio())
	pagesRouter.GET("/interactive", pages.Interactive())
	pagesRouter.GET("/software", pages.Software())
	pagesRouter.GET("/visual", pages.Visual())

	globalRouter.Group("/static",
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				if strings.HasPrefix(c.Request().URL.Path, "/static/fonts/") {
					c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
				}
				return next(c)
			}
		},
		middleware.StaticWithConfig(middleware.StaticConfig{
			Browse:     true,
			Filesystem: http.FS(staticAssets),
			Root:       "static",
		}),
	)

	return globalRouter
}
