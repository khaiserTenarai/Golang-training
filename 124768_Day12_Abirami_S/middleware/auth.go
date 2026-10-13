package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authenticate accepts the configured demo tokens. Use a real identity provider/JWT
// and secrets management for production deployments.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := strings.TrimSpace(c.GetHeader("Authorization"))
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"missing bearer token"})
			c.Abort(); return
		}
		token := parts[1]
		if admin := os.Getenv("ADMIN_TOKEN"); admin != "" && token == admin {
			c.Set("role", "admin"); c.Next(); return
		}
		if user := os.Getenv("USER_TOKEN"); user != "" && token == user {
			c.Set("role", "user"); c.Next(); return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
		c.Abort()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error":"admin role required"})
			c.Abort(); return
		}
		c.Next()
	}
}
