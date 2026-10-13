package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		requestID, _ := c.Get("request_id")
		log.Printf("request_id=%v method=%s path=%s status=%d duration=%s",
			requestID, c.Request.Method, c.Request.URL.RequestURI(), c.Writer.Status(), time.Since(start))
	}
}
