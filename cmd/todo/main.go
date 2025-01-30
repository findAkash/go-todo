package main

import (
	"go-todo/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Println("No .env file found")
	}

	//connect to the database
	database.Connect()
}