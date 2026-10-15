package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Trace(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s: before Next()", name)
		c.Next()
		log.Printf("%s: after Next()", name)
	}
}
