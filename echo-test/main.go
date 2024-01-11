package main

import (
	"echo-test/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.SetupMiddleware(e)
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
