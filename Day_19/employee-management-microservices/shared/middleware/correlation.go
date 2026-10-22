package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CorrelationIDHeader = "X-Correlation-ID"

func CorrelationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := strings.TrimSpace(c.GetHeader(CorrelationIDHeader))

		if id == "" || len(id) > 128 {
			id = uuid.NewString()
		}

		c.Set("correlation_id", id)
		c.Request.Header.Set(CorrelationIDHeader, id)
		c.Header(CorrelationIDHeader, id)

		log.Printf("correlation_id=%s method=%s path=%s",
			id, c.Request.Method, c.Request.URL.Path)

		c.Next()
	}
}

func PropagateCorrelationID(req *http.Request, id string) {
	if id != "" {
		req.Header.Set(CorrelationIDHeader, id)
	}
}
