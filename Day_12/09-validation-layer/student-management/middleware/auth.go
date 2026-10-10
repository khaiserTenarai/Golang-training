package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authentication validates the token and identifies the role.
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing bearer token",
			})
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		var role string

		switch token {
		case "student-token":
			role = "student"
		case "admin-token":
			role = "admin"
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		// Store the authenticated user's role.
		c.Set("role", role)

		c.Next()
	}
}
