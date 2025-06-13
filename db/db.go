package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("POSTGRES_DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	err = DB.AutoMigrate(&User{}, &Event{})
	if err != nil {
		log.Fatal("Failed migration: ", err)
	}
}
