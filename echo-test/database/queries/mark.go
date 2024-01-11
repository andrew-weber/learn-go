package database

import (
	"database/sql"
	models "echo-test/models/db"
	"log"
)

func GetMarks(db *sql.DB) (*[]models.Mark, error) {
	rows, err := db.Query("SELECT * FROM marks;")

	if err != nil {
		log.Printf("failed to execute query: %s", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.Mark

	for rows.Next() {
		var mark = models.Mark{}
		err := rows.Scan(&mark.ID, &mark.Date, &mark.Active)
		if err != nil {
			log.Printf("failed to scan row: %s", err)
			return nil, err
		}
		results = append(results, mark)
	}

	return &results, nil
}
