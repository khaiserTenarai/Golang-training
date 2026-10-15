package middleware

import (
	"employee-management/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method

		if rawQuery != "" {
			path = path + "?" + rawQuery
		}

		logger.Info(method + " | " + path + " | Status: " + string(rune(status)) + " | Latency: " + latency.String())
	}
}