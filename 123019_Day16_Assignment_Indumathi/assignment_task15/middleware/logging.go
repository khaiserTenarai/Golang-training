// Package middleware contains Gin HTTP middleware components.
package middleware

import (
	"fmt"
	"time"

	"employee-management/logging"

	"github.com/gin-gonic/gin"
)

// Logging returns a Gin middleware function that logs HTTP request details.
func Logging(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		logger.Info(fmt.Sprintf("%s %s -> %d (%s)", 
			c.Request.Method, 
			c.Request.URL.Path, 
			c.Writer.Status(), 
			duration,
		))
	}
}