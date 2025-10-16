package pages

import (
	"portfolio/internal/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Visual() echo.HandlerFunc {
	return templPage(
		metadata{
			title:       "Visual Media - Jackson Terrill",
			description: "Jackson Terrill's visual media",
		},
		components.Header("Visual Media", "By Jackson Terrill"),
		templ.Raw("visual media"), // TODO: Fill in content
	)
}
