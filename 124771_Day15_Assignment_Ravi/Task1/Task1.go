package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Before the handler executes
		log.Printf("Request started: %s %s",
			c.Request.Method,
			c.Request.URL.Path,
		)

		// Execute the next middleware or handler
		c.Next()

		// After the handler executes
		duration := time.Since(start)

		log.Printf("Request completed: %s %s | Status: %d | Duration: %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}

func main() {
	r := gin.New()

	// Register custom middleware globally
	r.Use(LoggerMiddleware())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}
