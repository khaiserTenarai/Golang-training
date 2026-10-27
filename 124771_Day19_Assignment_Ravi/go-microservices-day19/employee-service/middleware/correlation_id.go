package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

const CorrelationHeader = "X-Correlation-ID"

// CorrelationID reuses a caller's ID or generates one, then returns it in the response.
func CorrelationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(CorrelationHeader)
		if id == "" {
			id = uuid.NewString()
		}
		c.Set("correlation_id", id)
		c.Header(CorrelationHeader, id)
		start := time.Now()
		c.Next()
		log.Printf("correlation_id=%s method=%s path=%s status=%d duration=%s", id, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	}
}
