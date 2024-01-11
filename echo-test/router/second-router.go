package router

import (
	"echo-test/handlers"

	"github.com/labstack/echo/v4"
)

func AddSecondRoutes(e *echo.Echo) *echo.Group {
	router := e.Group("/kanto")

	router.GET("", handlers.GetAllKanto)

	return router
}
