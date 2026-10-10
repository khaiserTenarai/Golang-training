package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestID creates a unique ID for each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a request ID.
		requestID := fmt.Sprintf("req-%d", time.Now().UnixNano())

		// Store the ID in the request context.
		c.Set("requestID", requestID)

		// Send the ID in the response header.
		c.Header("X-Request-ID", requestID)

		// Continue to the next handler.
		c.Next()
	}
}
