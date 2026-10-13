// Task 5: Request ID Middleware

package middleware

import (
	"crypto/rand"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader("X-Request-ID")
		if rid == "" {
			b := make([]byte, 8)
			rand.Read(b)
			rid = fmt.Sprintf("req-%x", b)
		}
		c.Set("request_id", rid)
		c.Header("X-Request-ID", rid)
		c.Next()
	}
}
