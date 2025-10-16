// Pages to be served to users.
package pages

import (
	"portfolio/internal/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Index() echo.HandlerFunc {
	return templPage(
		metadata{
			title:       "Jackson Terrill",
			description: "Jackson Terrill's personal portfolio",
		},
		components.Header("Jackson\nTerrill", "Developer + Designer + Creator"),
		templ.Raw("index"), // TODO: Fill in content
	)
}

type metadata struct {
	title       string
	description string
}

func templPage(metadata metadata, children ...templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := templ.WithChildren(c.Request().Context(), templ.Join(children...))
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		return components.Layout(metadata.title, metadata.description).Render(ctx, c.Response())
	}
}
