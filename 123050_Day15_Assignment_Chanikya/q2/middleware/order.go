package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// First demonstrates the first middleware in the chain.
func First() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Middleware 1: Before Next()")

		c.Next()

		log.Println("Middleware 1: After Next()")
	}
}

// Second demonstrates the second middleware in the chain.
func Second() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Middleware 2: Before Next()")

		c.Next()

		log.Println("Middleware 2: After Next()")
	}
}