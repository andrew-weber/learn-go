package main

import (
	"echo-test/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/kanto", handlers.GetAllKanto)
	e.GET("/turso", handlers.Query)

	e.Logger.Fatal(e.Start(":3000"))
}
