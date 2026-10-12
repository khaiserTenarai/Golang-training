// Package middleware contains Gin middleware.
package middleware

import (
	"employee-management/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Logging logs every HTTP request after it completes.
func Logging(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info(fmt.Sprintf("%s %s %d %s", c.Request.Method, c.Request.URL.RequestURI(), c.Writer.Status(), time.Since(start)))
	}
}
