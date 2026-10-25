package middleware

import (
	"net/http"
	"strings"

	"ems/auth"

	"github.com/gin-gonic/gin"
)

const (
	UserKey = "username"
	RoleKey = "role"
)

func RequireAuth(a *auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the Authorization header.
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") ||
			strings.TrimSpace(parts[1]) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Bearer token required",
			})
			return
		}

		// Verify signature and expiration.
		claims, err := a.ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Save authenticated user information for later middleware.
		c.Set(UserKey, claims.Subject)
		c.Set(RoleKey, claims.Role)

		c.Next()
	}
}
