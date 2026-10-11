package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom logging middleware (Gin has built-in Logger, this adds request ID)
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()

		if query != "" {
			path = path + "?" + query
		}

		log.Printf("[GIN] %3d | %12v | %15s | %-7s %s",
			status, latency.Round(time.Microsecond), clientIP, method, path)
	}
}
