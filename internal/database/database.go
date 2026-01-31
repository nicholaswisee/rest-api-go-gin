package database

import (
	"fmt"
	"log"
	"os"
	"rest-api-go-gin/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // e.g., "root"
		os.Getenv("DB_PASSWORD"), // e.g., "password"
		os.Getenv("DB_HOST"),     // e.g., "localhost"
		os.Getenv("DB_PORT"),     // e.g., "3306"
		os.Getenv("DB_NAME"),     // e.g., "myapp"
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Attendee{},
		&models.Session{},
	)
}
