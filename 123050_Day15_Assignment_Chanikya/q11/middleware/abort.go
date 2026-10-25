package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortDemo() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Abort Middleware: Before Abort()")

		c.Abort()

		log.Println("Abort Middleware: After Abort()")

		c.JSON(http.StatusForbidden, gin.H{
			"message": "Request aborted by middleware",
		})
	}
}