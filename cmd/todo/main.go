package main

import (
	"go-todo/internal/database"
	"go-todo/internal/handlers"
	"go-todo/internal/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Println("No .env file found")
	}

	//connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto migration
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate todo table : %v", err)
	}

	// initialize the gin router
	router := gin.Default()

	// setup routes
	handlers.RegisterRoutes(router, db)

	//Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}