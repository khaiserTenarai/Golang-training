package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() 

		log.Printf("[%s] %s %s -> %d (%v)",
			c.GetString(CtxRequestID), c.Request.Method, c.Request.URL.Path,
			c.Writer.Status(), time.Since(start))
	}
}
