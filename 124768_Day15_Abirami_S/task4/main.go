package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Aborting")
		c.AbortWithStatusJSON(401, gin.H{"error": "Access denied"})
	}
}
func main() {
	r := gin.New()
	r.Use(auth())
	r.GET("/protected", func(c *gin.Context) { fmt.Println("Will not execute") })
	r.Run(":8080")
}
