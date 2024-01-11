package router

import (
	"database/sql"
	"echo-test/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

func SetupMiddleware(e *echo.Echo) {
	dbx, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db = dbx
	e.Use(middleware.Logger())
	e.Use(injectDbMiddleware)
}

func injectDbMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("db", db)
		return next(c)
	}
}
