package main

import (
	"database/sql"
	"echo-test/database"
	"echo-test/router"
	"log"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

func main() {
	dbx, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db = dbx

	e := echo.New()
	router.SetupMiddleware(e, db)
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
