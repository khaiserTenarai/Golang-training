package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging records details about each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the request start time.
		start := time.Now()

		// Process the request.
		c.Next()

		// Calculate the processing time.
		duration := time.Since(start)

		// Get the request ID.
		requestID, _ := c.Get("requestID")

		// Log request details.
		log.Printf(
			"RequestID: %v | Method: %s | Path: %s | Status: %d | Duration: %s",
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
