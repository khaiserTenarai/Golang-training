package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func m1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("M1 before")
		c.Next()
		fmt.Println("M1 after")
	}
}
func m2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("M2 before")
		c.Next()
		fmt.Println("M2 after")
	}
}
func main() {
	r := gin.New()
	r.Use(m1(), m2())
	r.GET("/order", func(c *gin.Context) { fmt.Println("Handler"); c.JSON(200, gin.H{"message": "check terminal"}) })
	r.Run(":8080")
}
