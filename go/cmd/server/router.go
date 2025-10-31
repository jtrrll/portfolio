package main

import (
	"net/http"

	"github.com/jtrrll/portfolio/internal/middleware"
	"github.com/jtrrll/portfolio/internal/pages"

	"embed"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
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
	pagesRouter := globalRouter.Group("", middleware.RedirectWhenNotFound("/"))
	pagesRouter.GET("/", templPage(pages.Index()))
	pagesRouter.GET("/audio", templPage(pages.Audio()))
	pagesRouter.GET("/interactive", templPage(pages.Interactive()))
	pagesRouter.GET("/software", templPage(pages.Software()))
	pagesRouter.GET("/software/:name", func(c echo.Context) error {
		return templPage(pages.SoftwareProject(c.Param("name")))(c)
	})
	pagesRouter.GET("/visual", templPage(pages.Visual()))

	globalRouter.Group("/static", middleware.ServeStaticAssets(http.FS(staticAssets), "static"))

	return globalRouter
}

func templPage(page templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		return page.Render(c.Request().Context(), c.Response())
	}
}
