package main

import (
	"echo-test/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/kanto", router.GetAllKanto)

	e.Logger.Fatal(e.Start(":3000"))
}
