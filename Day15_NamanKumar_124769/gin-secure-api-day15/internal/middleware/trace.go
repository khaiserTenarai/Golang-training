package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func TraceOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("middleware before: path=%s", c.Request.URL.Path)
		c.Next()
		log.Printf("middleware after: path=%s status=%d", c.Request.URL.Path, c.Writer.Status())
	}
}

func Timing() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		c.Header("X-Response-Time", time.Since(start).String())
	}
}
