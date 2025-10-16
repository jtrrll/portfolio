package pages

import (
	"portfolio/internal/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Software() echo.HandlerFunc {
	return templPage(
		metadata{
			title:       "Software - Jackson Terrill",
			description: "Jackson Terrill's software projects",
		},
		components.Header("Software", "By Jackson Terrill"),
		templ.Raw("software"), // TODO: Fill in content
	)
}
