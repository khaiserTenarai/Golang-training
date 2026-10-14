package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/products", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v1", "products": []string{"laptop", "phone"}})
		})
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/products", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v2", "items": []string{"laptop v2", "tablet"}})
		})
	}

	r.Run(":8080")
}
