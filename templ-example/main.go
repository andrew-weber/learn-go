package main

import (
	web "web/hello"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	component := web.Hello("John")

	e.GET("/", templ.Handler(component))
	e.Logger.Fatal(e.Start(":3000"))
}
