package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const (
	CtxRequestID = "requestID"
	CtxUserID    = "userID"
	CtxRole      = "role"
)


func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			b := make([]byte, 8)
			_, _ = rand.Read(b)
			id = hex.EncodeToString(b)
		}
		c.Set(CtxRequestID, id)
		c.Writer.Header().Set("X-Request-ID", id)
		c.Next()
	}
}
