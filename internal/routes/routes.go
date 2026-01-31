package routes

import (
	"net/http"
	"rest-api-go-gin/internal/handlers"
	"rest-api-go-gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	authHandler := handlers.NewAuthHandler()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Public auth routes (no auth required)
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}

	// Protected routes (auth required)
	api := router.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		// Get current user
		api.GET("/me", func(c *gin.Context) {
			userID := c.GetUint("userID")
			email := c.GetString("email")
			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"email":   email,
			})
		})

		// Events routes (placeholder - add your event handler)
		api.GET("/events", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List all events"})
		})
		api.POST("/events", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "Create event"})
		})
	}
}
