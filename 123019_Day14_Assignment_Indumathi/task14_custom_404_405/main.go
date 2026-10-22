package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	// Custom 404 Handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"error":   "Resource Not Found",
			"path":    c.Request.URL.Path,
		})
	})

	// Custom 405 Handler
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":   405,
			"error":  "Method Not Allowed",
			"method": c.Request.Method,
		})
	})

	r.GET("/valid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Valid route"})
	})

	r.Run(":8080")
}
