package router

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	ApiRouter(e)
	WebRouter(e)
}
