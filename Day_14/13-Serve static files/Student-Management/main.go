package main

import (
	"student-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin router.
	router := gin.Default()

	// Register student API routes.
	routes.StudentRoutes(router)

	// Serve static files from the static folder.
	router.Static("/static", "./static")

	// Serve index.html at the home URL.
	router.StaticFile("/", "./static/index.html")

	// Start the server on port 8080.
	router.Run(":8080")
}
