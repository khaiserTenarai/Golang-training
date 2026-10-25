package correlation

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const Header = "X-Correlation-ID"

// Middleware creates or forwards a correlation ID.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(Header)

		if id == "" {
			id = uuid.NewString()
		}

		c.Set("correlation_id", id)
		c.Header(Header, id)
		c.Next()
	}
}

// ID returns the correlation ID from the request.
func ID(c *gin.Context) string {
	value, exists := c.Get("correlation_id")
	if !exists {
		return ""
	}

	id, _ := value.(string)
	return id
}
