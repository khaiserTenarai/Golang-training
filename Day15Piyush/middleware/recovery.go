// Task 6: Recovery Middleware — catches panics

package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("  [PANIC RECOVERED] %v\n", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Internal server error",
				})
			}
		}()
		c.Next()
	}
}
