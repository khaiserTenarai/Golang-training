package middleware

import (
	"fmt"
	"time"

	"employee-management/logging"

	"github.com/gin-gonic/gin"
)

func Logging(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Info(fmt.Sprintf(
			"requestID=%s method=%s path=%s status=%d duration=%s",
			c.GetString("requestID"),
			c.Request.Method,
			c.Request.URL.RequestURI(),
			c.Writer.Status(),
			time.Since(start),
		))
	}
}
