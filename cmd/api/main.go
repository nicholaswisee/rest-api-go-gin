package main

import (
	"log"
	"os"
	"rest-api-go-gin/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database and run migrations
	database.Connect()
	database.AutoMigrate()
}

func main() {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 routes group
	root := router.Group("/api")
	{
		// Users routes
		root.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "List all users"})
		})

		// Events routes
		root.GET("/events", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "List all events"})
		})
	}

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
