package db

import (
	"authtoken/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB : Database
var DB = Database()

// Database: Database Connection
func Database() *gorm.DB {
	dsn := "user=postgres password=postgres dbname=auth_token host=localhost sslmode=disable TimeZone=Asia/Kolkata"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.Logger.LogMode(1)

	Migration(DB)

	return DB
}

// Migrations ...
func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
