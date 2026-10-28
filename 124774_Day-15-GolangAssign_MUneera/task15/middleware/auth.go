package middleware

import (
	"net/http"

	"employee-management/auth"

	"github.com/gin-gonic/gin"
)

func RequireAuth(a *auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, role, err := a.Parse(
			c.GetHeader("Authorization"),
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized: " + err.Error(),
			})
			return
		}

		c.Set("username", username)
		c.Set("role", role)

		c.Next()
	}
}
