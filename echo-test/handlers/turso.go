package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Query(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		c.Error(err)
	}
	dbHost := os.Getenv("DATABASE_URL")

	db, err := sql.Open("libsql", dbHost)
	if err != nil {
		log.Fatalf("failed to open db %s: %s", dbHost, err)
		c.Error(err)
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM marks;")
	if err != nil {
		log.Printf("failed to execute query: %s", err)
		c.Error(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var temp, date string // Adjust these types based on your actual table columns

		// Scan the values from the current row into variables
		err := rows.Scan(&temp, &date, &temp)
		if err != nil {
			log.Printf("failed to scan row: %s", err)
			c.Error(err)
			return err
		}

		// Print or process the values as needed
		fmt.Printf("Column1: %s \n", date)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbHost, err)
		c.Error(err)
	}

	return c.NoContent(http.StatusNoContent)
}
