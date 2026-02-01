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
		// GetMe godoc
		// @Summary      Get current user
		// @Description  Returns the authenticated user's info
		// @Tags         user
		// @Produce      json
		// @Security     BearerAuth
		// @Success      200 {object} map[string]interface{} "User info"
		// @Failure      401 {object} map[string]interface{} "Unauthorized"
		// @Router       /api/me [get]
		api.GET("/me", func(c *gin.Context) {
			userID := c.GetUint("userID")
			email := c.GetString("email")
			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"email":   email,
			})
		})

		// GetEvents godoc
		// @Summary      List all events
		// @Description  Returns a list of all events
		// @Tags         events
		// @Produce      json
		// @Security     BearerAuth
		// @Success      200 {object} map[string]interface{} "List of events"
		// @Failure      401 {object} map[string]interface{} "Unauthorized"
		// @Router       /api/events [get]
		api.GET("/events", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List all events"})
		})

		// CreateEvent godoc
		// @Summary      Create an event
		// @Description  Create a new event
		// @Tags         events
		// @Accept       json
		// @Produce      json
		// @Security     BearerAuth
		// @Success      201 {object} map[string]interface{} "Event created"
		// @Failure      401 {object} map[string]interface{} "Unauthorized"
		// @Router       /api/events [post]
		api.POST("/events", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "Create event"})
		})
	}
}
