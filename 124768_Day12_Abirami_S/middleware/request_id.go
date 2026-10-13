package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID adds a unique ID to each request and returns it in the response header.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" { id = uuid.NewString() }
		c.Set("request_id", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}
