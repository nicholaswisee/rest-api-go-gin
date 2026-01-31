package main

import (
	"rest-api-go-gin/internal/database"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.Connect()
	database.AutoMigrate()
}

func main() {

}
