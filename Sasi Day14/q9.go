package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")

		c.JSON(http.StatusOK, gin.H{
			"query": query,
			"page":  page,
			"limit": limit,
		})
	})

	r.Run(":8080")
}
