package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func first() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("First before")
		c.Next()
		fmt.Println("First after")
	}
}
func second() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Second before")
		c.Next()
		fmt.Println("Second after")
	}
}
func main() {
	r := gin.New()
	r.Use(first(), second())
	r.GET("/next", func(c *gin.Context) { fmt.Println("Handler"); c.JSON(200, gin.H{"message": "Next continues chain"}) })
	r.Run(":8080")
}
