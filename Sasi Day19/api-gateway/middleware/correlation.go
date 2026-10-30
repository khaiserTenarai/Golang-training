package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const CorrelationIDHeader = "X-Correlation-ID"

func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		corrID := c.GetHeader(CorrelationIDHeader)
		if corrID == "" {
			corrID = fmt.Sprintf("corr-gw-%d", time.Now().UnixNano())
		}
		c.Set(CorrelationIDHeader, corrID)
		c.Header(CorrelationIDHeader, corrID)
		c.Next()
	}
}
