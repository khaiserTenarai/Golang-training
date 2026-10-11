// Package middleware contains custom Gin middlewares.
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"employee-management/logging"
)

// Logging returns a Gin HandlerFunc for structured request logging.
func Logging(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)

		logger.Info(fmt.Sprintf("%s %s -> Status %d (%s)", c.Request.Method, path, status, duration))
	}
}