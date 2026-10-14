package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("Recovered:", e)
				c.AbortWithStatusJSON(500, gin.H{"error": "Internal server error"})
			}
		}()
		c.Next()
	}
}
func main() {
	r := gin.New()
	r.Use(recovery())
	r.GET("/panic", func(c *gin.Context) { panic("something went wrong") })
	r.Run(":8080")
}
