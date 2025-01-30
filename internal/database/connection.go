package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var DB *gorm.DB


func Connect(){
	dbUser := os.Getenv("DB_USER")
	dbPassword :=os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// construct the data scource name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb&paseTime=True&loc=Local",
dbUser,dbPassword,dbName,dbHost,dbPort)

var err error
DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Silent),
})

if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
}
}

// Optionally, auto-migrate your models
    // err = DB.AutoMigrate(&models.Todo{})
    // if err != nil {
    //     log.Fatalf("Failed to auto-migrate models: %v", err)
    // }