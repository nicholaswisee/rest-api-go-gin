package main

import (
	"log"
	"os"
	"rest-api-go-gin/internal/database"
	"rest-api-go-gin/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.Connect()
	database.AutoMigrate()
}

func main() {
	router := gin.Default()

	// Setup all routes
	routes.Setup(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
