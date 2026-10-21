package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a unique ID for this request.
		requestID, err := generateRequestID()
		if err != nil {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}

		// Store the ID in Gin context.
		c.Set("RequestID", requestID)

		// Add the ID to the response header.
		c.Header("X-Request-ID", requestID)

		// Continue executing the remaining handlers.
		c.Next()
	}
}

func generateRequestID() (string, error) {
	bytes := make([]byte, 16)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}