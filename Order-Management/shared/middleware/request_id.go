package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestIDContextKey struct{}

// RequestIDMiddleware assigns a unique ID to every incoming request and propagates it.
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("request_id", requestID)
		c.Request.Header.Set("X-Request-ID", requestID)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), requestIDContextKey{}, requestID))
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// RequestIDFromContext returns the request ID, if middleware assigned one.
func RequestIDFromContext(ctx context.Context) string {
	value, _ := ctx.Value(requestIDContextKey{}).(string)
	return value
}
