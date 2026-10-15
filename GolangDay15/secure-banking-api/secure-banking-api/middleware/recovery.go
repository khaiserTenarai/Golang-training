package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)


func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[%s] panic: %v\n%s", c.GetString(CtxRequestID), r, debug.Stack())
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false, "message": "internal server error",
				})
			}
		}()
		c.Next()
	}
}
