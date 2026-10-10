package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authentication checks the bearer token.
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the Authorization header.
		authHeader := c.GetHeader("Authorization")

		// Check the bearer token format.
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing bearer token",
			})
			return
		}

		// Extract the token.
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token for this learning example.
		if token != "student-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		// Continue if authentication succeeds.
		c.Next()
	}
}
