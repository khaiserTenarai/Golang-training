package utility

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const CorrelationHeader = "X-Correlation-ID"

func CorrelationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(CorrelationHeader)
		if id == "" {
			id = newCorrelationID()
		}
		c.Set("correlation_id", id)
		c.Header(CorrelationHeader, id)
		c.Next()
	}
}

func CorrelationID(c *gin.Context) string {
	value, _ := c.Get("correlation_id")
	id, _ := value.(string)
	return id
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("correlation_id=%s method=%s path=%s status=%d duration=%s",
			CorrelationID(c), c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	}
}

func ErrorJSON(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message, "correlation_id": CorrelationID(c)})
}

func ForwardCorrelation(req *http.Request, id string) {
	if id != "" {
		req.Header.Set(CorrelationHeader, id)
	}
}

func newCorrelationID() string {
	// Unique enough for this teaching example; use a UUID library in production.
	return time.Now().UTC().Format("20060102T150405.000000000Z")
}
