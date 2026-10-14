package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlewareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("A: Before")
		c.Next()
		fmt.Println("A: After")
	}
}

func MiddlewareB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("B: Before")
		c.Next()
		fmt.Println("B: After")
	}
}

func MiddlewareC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("C: Before")
		c.Next()
		fmt.Println("C: After")
	}
}

func main() {
	r := gin.New()

	// Register middleware in order
	r.Use(MiddlewareA())
	r.Use(MiddlewareB())

	r.GET("/hello", MiddlewareC(), func(c *gin.Context) {
		fmt.Println("Handler: Executing")
		c.JSON(200, gin.H{"message": "Hello"})
	})

	r.Run(":8080")
}
