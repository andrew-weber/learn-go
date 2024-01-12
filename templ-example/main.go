package main

import (
	web "templ-exampl/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderHtml(cmp templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return cmp.Render(c.Request().Context(), c.Response().Writer)
	}
}

func main() {
	e := echo.New()

	hello := web.Hello("John")
	test := web.Test("asdfasdf")

	e.GET("/", renderHtml(hello))
	e.GET("/1", renderHtml(test))

	// e.GET("/")
	e.Logger.Fatal(e.Start(":3000"))
}
