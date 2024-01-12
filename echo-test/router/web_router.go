package router

import (
	"echo-test/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderHtml(cmp templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return cmp.Render(c.Request().Context(), c.Response().Writer)
	}
}

func WebRouter(e *echo.Echo) *echo.Group {
	e.GET("/", renderHtml(web.Test("test")))
	router := e.Group("/web")
	router.GET("/", renderHtml(web.Test("test")))

	return router
}
