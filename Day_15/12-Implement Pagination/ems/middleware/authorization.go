package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(RoleKey)

		// Reject requests that do not have the required role.
		if !exists || role != allowedRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "insufficient permissions",
			})
			return
		}

		c.Next()
	}
}
