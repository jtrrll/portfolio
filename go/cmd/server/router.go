package main

import (
	"net/http"
	"time"

	"github.com/jtrrll/portfolio/internal/handlers"
	"github.com/jtrrll/portfolio/internal/middleware"
	"github.com/jtrrll/portfolio/internal/pages"

	"embed"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

//go:embed static
var staticAssets embed.FS

// NewRouter creates an HTTP request handler.
func NewRouter(trustProxy bool) http.Handler {
	globalRouter := echo.New()

	if trustProxy {
		globalRouter.IPExtractor = echo.ExtractIPFromXFFHeader()
	}

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

	apiRouter := globalRouter.Group("/api")
	apiRouter.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	pagesRouter := globalRouter.Group("",
		middleware.CacheControl(PAGE_MAX_AGE),
		middleware.RedirectWhenNotFound("/"))
	pagesRouter.GET("/", handlers.TemplHandler(pages.Index()))
	pagesRouter.GET("/audio", handlers.TemplHandler(pages.Audio()))
	pagesRouter.GET("/interactive", handlers.TemplHandler(pages.Interactive()))
	pagesRouter.GET("/software", handlers.TemplHandler(pages.Software()))
	pagesRouter.GET("/software/:name", func(c echo.Context) error {
		return handlers.TemplHandler(pages.SoftwareProject(c.Param("name")))(c)
	})
	pagesRouter.GET("/visual", handlers.TemplHandler(pages.Visual()))

	globalRouter.Group("/static",
		middleware.CacheControl(STATIC_ASSET_MAX_AGE),
		echoMiddleware.StaticWithConfig(echoMiddleware.StaticConfig{
			Browse:     ENABLE_STATIC_ASSET_BROWSING,
			Filesystem: http.FS(staticAssets),
			Root:       "static",
		}))

	return globalRouter
}
