// Tasks 1, 2, 3, 4: Custom Middleware Concepts
// Task 1 — Create custom Gin middleware
// Task 2 — Demonstrate middleware execution order
// Task 3 — Demonstrate Next()
// Task 4 — Demonstrate Abort()

package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Task 1: Custom Gin middleware — adds timing header
func TimingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Task 3: c.Next() passes control to the next handler
		// Code BEFORE Next() runs before the handler
		fmt.Printf("  [Timing] BEFORE handler: %s %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next() // Task 3: Execute the rest of the chain

		// Code AFTER Next() runs after the handler returns
		duration := time.Since(start)
		c.Header("X-Response-Time", duration.String())
		fmt.Printf("  [Timing] AFTER handler: %v (status: %d)\n", duration.Round(time.Microsecond), c.Writer.Status())
	}
}

// Task 2: Middleware execution order — First, Second, Third
func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("  [Order] First middleware — BEFORE")
		c.Next()
		fmt.Println("  [Order] First middleware — AFTER")
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("  [Order] Second middleware — BEFORE")
		c.Next()
		fmt.Println("  [Order] Second middleware — AFTER")
	}
}

// Task 4: Demonstrate Abort() — stops the middleware chain
func ApiKeyMiddleware(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key == "" || key != validKey {
			// Task 4: Abort stops remaining handlers from executing
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Invalid or missing API key",
			})
			return // Handler and subsequent middleware are NEVER called
		}
		c.Set("api_key_validated", true)
		c.Next()
	}
}

// Task 11: Request validation middleware
func ValidateContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			ct := c.GetHeader("Content-Type")
			if !strings.Contains(ct, "application/json") && !strings.Contains(ct, "multipart/form-data") {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
					"success": false,
					"message": "Content-Type must be application/json or multipart/form-data",
				})
				return
			}
		}
		c.Next()
	}
}
