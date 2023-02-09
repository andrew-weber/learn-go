package main

import (
	"log"
	"os"
	"webserver/controllers"
	"webserver/global"
	"webserver/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	global.SetupDb()

	// Migrate the schema
	global.DB.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()

	r.POST("/create", controllers.CreateUser)

	r.Run(":" + os.Getenv("PORT"))
}
