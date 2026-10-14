package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"
const RequestIDHeader = "X-Request-ID"

// Request ID Middleware
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.NewString()

		c.Set(RequestIDKey, requestID)
		c.Header(RequestIDHeader, requestID)

		c.Next()
	}
}

// Recovery Middleware
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID, _ := c.Get(RequestIDKey)

				log.Printf(
					"PANIC: request_id=%v error=%v\n%s",
					requestID,
					err,
					debug.Stack(),
				)

				// Abort further handler execution
				c.Abort()

				// Return a generic error response
				if !c.Writer.Written() {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error":      "Internal server error",
						"request_id": requestID,
					})
				}
			}
		}()

		c.Next()
	}
}

func main() {
	r := gin.New()

	// Register middleware
	r.Use(RequestIDMiddleware())
	r.Use(RecoveryMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Test recovery middleware
	r.GET("/panic", func(c *gin.Context) {
		panic("something went wrong")
	})

	log.Fatal(r.Run(":8080"))
}
