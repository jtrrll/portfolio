package main

import (
	"net/http"
	"portfolio/internal/middleware"
	"portfolio/internal/pages"
	"portfolio/internal/static"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/go-pkgz/routegroup"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Initializes the server's HTTP handler.
func newHTTPHandler() http.Handler {
	router := routegroup.New(http.NewServeMux())
	router.Use(middleware.DoNotPanic())
	router.NotFoundHandler(func(w http.ResponseWriter, r *http.Request) {
		routegroup := strings.Split(r.URL.Path, "/")[1]
		if !(routegroup == "api" || routegroup == "static") {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})

	apiGroup := router.Mount("/api")
	apiGroup.Use(otelhttp.NewMiddleware("api"))
	// TODO: Serve API docs at /api
	apiGroup.HandleFunc("/healthcheck", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	pagesGroup := router.Group()
	pagesGroup.Use(otelhttp.NewMiddleware("pages"))
	pagesGroup.HandleFunc("/", templ.Handler(pages.Index()).ServeHTTP)
	pagesGroup.Handle("/code", templ.Handler(pages.Code()))
	pagesGroup.Handle("/games", templ.Handler(pages.Games()))
	pagesGroup.Handle("/music", templ.Handler(pages.Music()))
	pagesGroup.Handle("/art", templ.Handler(pages.Art()))

	staticGroup := router.Mount("/static")
	staticGroup.Use(otelhttp.NewMiddleware("static"), middleware.CacheControl(time.Hour))
	staticGroup.HandleFiles("/", http.FS(static.StaticFs))

	return router
}
