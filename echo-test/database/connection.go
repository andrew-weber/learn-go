package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dbHost := os.Getenv("DATABASE_URL")

	db, err := sql.Open("libsql", dbHost)
	if err != nil {
		return nil, err
	}

	return db, nil
}
