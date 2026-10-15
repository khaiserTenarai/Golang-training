package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/contact", func(c *gin.Context) {
		name := c.PostForm("name")
		message := c.DefaultPostForm("message", "No message provided")

		c.JSON(http.StatusOK, gin.H{
			"status":  "submitted",
			"name":    name,
			"message": message,
		})
	})

	r.Run(":8080")
}
