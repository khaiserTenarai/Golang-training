package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Set("request_id", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}
func main() {
	r := gin.New()
	r.Use(requestID())
	r.GET("/request-id", func(c *gin.Context) { id, _ := c.Get("request_id"); c.JSON(200, gin.H{"request_id": id}) })
	r.Run(":8080")
}
