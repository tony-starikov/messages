package models

import (
	"github.com/joho/godotenv"
  "os"

	"gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load()

  if err != nil {
		panic("Error loading .env file")
  }

	dbName := os.Getenv("DB_FILE")

	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Message{})

	DB = database
}
