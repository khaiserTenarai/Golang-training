package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func customMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before")
		c.Next()
		fmt.Println("After")
	}
}
func main() {
	r := gin.New()
	r.Use(customMiddleware())
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	r.Run(":8080")
}
