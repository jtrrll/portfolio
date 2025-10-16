package pages

import (
	"portfolio/internal/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Audio() echo.HandlerFunc {
	return templPage(
		metadata{
			title:       "Audio Media - Jackson Terrill",
			description: "Jackson Terrill's audio media",
		},
		components.Header("Audio Media", "By Jackson Terrill"),
		templ.Raw("audio media"), // TODO: Fill in content
	)
}
