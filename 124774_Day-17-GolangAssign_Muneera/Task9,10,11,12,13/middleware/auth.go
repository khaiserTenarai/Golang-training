package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type contextKey string

const roleKey contextKey = "userRole"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") ||
			strings.TrimSpace(parts[1]) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		token := strings.TrimSpace(parts[1])
		role := ""

		adminToken := os.Getenv("ADMIN_TOKEN")
		userToken := os.Getenv("USER_TOKEN")

		if adminToken != "" && token == adminToken {
			role = "ADMIN"
		} else if userToken != "" && token == userToken {
			role = "USER"
		}

		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		ctx := context.WithValue(c.Request.Context(), roleKey, role)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func RoleFromContext(ctx context.Context) string {
	role, _ := ctx.Value(roleKey).(string)
	return role
}
