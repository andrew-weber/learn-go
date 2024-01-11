package main

import (
	"database/sql"
	dbc "echo-test/database"
	"echo-test/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

func injectDbMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("db", db)
		return next(c)
	}
}

func main() {
	e := echo.New()

	dbx, err := dbc.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db = dbx

	e.Use(middleware.Logger())
	e.Use(injectDbMiddleware)

	e.GET("/kanto", handlers.GetAllKanto)
	e.GET("/turso", handlers.Query)

	e.Logger.Fatal(e.Start(":3000"))
}
