package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func TemplHandler(component templ.Component) echo.HandlerFunc {
	templHandler := templ.Handler(component, templ.WithStreaming())
	return echo.WrapHandler(templHandler)
}
