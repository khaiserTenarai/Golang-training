package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := "invalid-token"

		if token != "my-secret-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.GET("/private", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the private route",
		})
	})

	r.Run(":8080")
}
