package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		requestID, _ := c.Get("RequestID")

		log.Printf(
			"Request ID: %v | Method: %s | Path: %s | Status: %d | Duration: %s",
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}