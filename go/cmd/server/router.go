package main

import (
	"net/http"
	"time"

	"github.com/jtrrll/portfolio/internal/middleware"
	"github.com/jtrrll/portfolio/internal/pages"

	"embed"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

//go:embed static
var staticAssets embed.FS

// NewRouter creates an HTTP request handler.
func NewRouter() http.Handler {
	globalRouter := echo.New()
	globalRouter.Use(
		echoMiddleware.Recover(),
		echoMiddleware.Secure(),
		echoMiddleware.RateLimiter(echoMiddleware.NewRateLimiterMemoryStore(20)),
		echoMiddleware.BodyLimit("1MB"),
		echoMiddleware.Decompress(),
		echoMiddleware.GzipWithConfig(echoMiddleware.GzipConfig{
			Level: 5,
		}),
		echoMiddleware.ContextTimeout(60*time.Second),
	)

	globalRouter.Group("/api")

	pagesRouter := globalRouter.Group("",
		middleware.CacheControl(PAGE_MAX_AGE),
		middleware.RedirectWhenNotFound("/"))
	pagesRouter.GET("/", templHandler(pages.Index()))
	pagesRouter.GET("/audio", templHandler(pages.Audio()))
	pagesRouter.GET("/interactive", templHandler(pages.Interactive()))
	pagesRouter.GET("/software", templHandler(pages.Software()))
	pagesRouter.GET("/software/:name", func(c echo.Context) error {
		return templHandler(pages.SoftwareProject(c.Param("name")))(c)
	})
	pagesRouter.GET("/visual", templHandler(pages.Visual()))

	globalRouter.Group("/static",
		middleware.CacheControl(STATIC_ASSET_MAX_AGE),
		echoMiddleware.StaticWithConfig(echoMiddleware.StaticConfig{
			Browse:     true,
			Filesystem: http.FS(staticAssets),
			Root:       "static",
		}))

	return globalRouter
}

func templHandler(component templ.Component) echo.HandlerFunc {
	templHandler := templ.Handler(component, templ.WithStreaming())
	return echo.WrapHandler(templHandler)
}
