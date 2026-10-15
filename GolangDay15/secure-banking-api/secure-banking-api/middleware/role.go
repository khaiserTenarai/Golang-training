package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(CtxRole)
		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"success": false, "message": "insufficient permissions"})
	}
}
