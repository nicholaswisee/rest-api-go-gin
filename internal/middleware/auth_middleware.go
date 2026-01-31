package middleware

import (
	"net/http"
	"rest-api-go-gin/internal/repositories"
	"rest-api-go-gin/internal/services"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	authService := &services.AuthService{}
	sessionRepo := &repositories.SessionRepository{}

	return func(c *gin.Context) {
		// Get token from header: "Authorization: Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}

		// Extract token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		tokenString := parts[1]

		// Validate JWT
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check session exists and is not revoked
		session, err := sessionRepo.FindByToken(tokenString)
		if err != nil || session.Revoked {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired or revoked"})
			return
		}

		// Check if session has expired
		if session.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
