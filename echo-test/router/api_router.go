package router

import (
	"echo-test/handlers"

	"github.com/labstack/echo/v4"
)

func ApiRouter(e *echo.Echo) *echo.Group {
	router := e.Group("/api")

	router.GET("/pokemon", handlers.GetAllKanto)
	router.GET("/db", handlers.Query)

	return router
}
