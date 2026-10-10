package main

import (
	"net/http"
	"student-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Enable custom 405 handling.
	router.HandleMethodNotAllowed = true

	// Register student API routes.
	routes.StudentRoutes(router)

	// Serve static files.
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

	// Handle unknown URLs.
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Page not found",
			"message": "The requested URL does not exist",
		})
	})

	// Handle unsupported HTTP methods.
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"error":   "Method not allowed",
			"message": "This HTTP method is not supported for this URL",
		})
	})

	// Start the server.
	router.Run(":8080")
}
