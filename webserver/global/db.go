package global

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDb() {
	// Connect to the database
	db, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	DB = db

	if err != nil {
		panic("failed to connect database")
	}
}
