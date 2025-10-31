package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectWhenNotFound(url string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if httpErr, ok := err.(*echo.HTTPError); ok && httpErr.Code == http.StatusNotFound {
				return c.Redirect(http.StatusSeeOther, url)
			}
			return err
		}
	}
}
