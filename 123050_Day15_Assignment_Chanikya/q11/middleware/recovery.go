package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Printf("Panic recovered: %v", recovered)

		c.AbortWithStatusJSON(500, gin.H{
			"error": "Internal Server Error",
		})
	})
}