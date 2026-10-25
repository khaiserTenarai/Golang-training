package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging creates custom HTTP request logging middleware.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the request start time.
		start := time.Now()

		// Process the next middleware and route handler.
		c.Next()

		// Calculate the time taken to process the request.
		duration := time.Since(start)

		// Log the request details.
		log.Printf(
			"Method: %s | Path: %s | Status: %d | Duration: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}