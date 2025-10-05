package main

import (
	"net/http"
	"strings"

	"embed"

	"portfolio/internal/components"

	"github.com/a-h/templ"
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
	pagesRouter.GET("/", templPage(
		"Jackson Terrill",
		"Jackson Terrill's personal portfolio",
		components.Header("Jackson\nTerrill", "Developer + Designer + Creator"),
		templ.Raw("index"), // TODO: Fill in content
	))
	pagesRouter.GET("/software", templPage(
		"Software - Jackson Terrill",
		"Jackson Terrill's software projects",
		components.Header("Software", "By Jackson Terrill"),
		templ.Raw("software"), // TODO: Fill in content
	))
	pagesRouter.GET("/interactive", templPage(
		"Interactive Media - Jackson Terrill",
		"Jackson Terrill's interactive media",
		components.Header("Interactive Media", "By Jackson Terrill"),
		templ.Raw("interactive media"), // TODO: Fill in content
	))
	pagesRouter.GET("/visual", templPage(
		"Visual Media - Jackson Terrill",
		"Jackson Terrill's visual media",
		components.Header("Visual Media", "By Jackson Terrill"),
		templ.Raw("visual media"), // TODO: Fill in content
	))
	pagesRouter.GET("/audio", templPage(
		"Audio Media - Jackson Terrill",
		"Jackson Terrill's audio media",
		components.Header("Audio Media", "By Jackson Terrill"),
		templ.Raw("audio media"), // TODO: Fill in content
	))

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

func templPage(title string, description string, children ...templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := templ.WithChildren(c.Request().Context(), templ.Join(children...))
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		return components.Layout(title, description).Render(ctx, c.Response())
	}
}
