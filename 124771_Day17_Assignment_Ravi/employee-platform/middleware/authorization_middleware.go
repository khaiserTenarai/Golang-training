package middleware

import (
	"employee-platform/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Task 7: role-based authorization for REST routes.
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get(ClaimsKey)
		claims, valid := value.(*utils.Claims)
		if !ok || !valid || claims.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}
		c.Next()
	}
}
