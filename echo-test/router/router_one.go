package router

import (
	"echo-test/handlers"

	"github.com/labstack/echo/v4"
)

func AddFirstRoutes(e *echo.Echo) *echo.Group {
	router := e.Group("/db")

	router.GET("", handlers.Query)

	return router
}
