package gateway

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// RequestID adds a correlation ID. This demo does not implement JWT auth
// or rate limiting; those can be added as separate middleware later.
func RequestID() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.GetHeader("X-Request-ID")
        if id == "" { id = uuid.NewString() }
        c.Writer.Header().Set("X-Request-ID", id)
        c.Set("request_id", id)
        c.Next()
    }
}
