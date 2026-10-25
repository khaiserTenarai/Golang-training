package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// NextDemo demonstrates execution before and after c.Next().
func NextDemo() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Next Middleware: Before Next()")

		c.Next()

		log.Println("Next Middleware: After Next()")
	}
}