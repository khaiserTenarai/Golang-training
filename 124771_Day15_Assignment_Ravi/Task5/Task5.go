package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"
const RequestIDHeader = "X-Request-ID"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.NewString()

		c.Set(RequestIDKey, requestID)
		c.Header(RequestIDHeader, requestID)

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Register middleware globally
	r.Use(RequestIDMiddleware())

	r.GET("/", func(c *gin.Context) {
		requestID, _ := c.Get(RequestIDKey)

		c.JSON(http.StatusOK, gin.H{
			"message":    "Hello World",
			"request_id": requestID,
		})
	})

	r.Run(":8080")
}
