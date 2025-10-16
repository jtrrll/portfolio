package pages

import (
	"portfolio/internal/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Interactive() echo.HandlerFunc {
	return templPage(
		metadata{
			title:       "Interactive Media - Jackson Terrill",
			description: "Jackson Terrill's interactive media",
		},
		components.Header("Interactive Media", "By Jackson Terrill"),
		templ.Raw("interactive media"), // TODO: Fill in content
	)
}
