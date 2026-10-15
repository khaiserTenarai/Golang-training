package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve single static file
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// Serve static directory assets
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Static files server is ready.")
	})

	r.Run(":8080")
}
