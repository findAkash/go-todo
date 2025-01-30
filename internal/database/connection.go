package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func Connect()(*gorm.DB, error){
	dbUser := os.Getenv("DB_USER")
	dbPassword :=os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// construct the data scource name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
dbUser,dbPassword,dbHost,dbPort,dbName)
fmt.Println(dsn)
var err error
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Silent),
})

if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
}
return db, nil

}