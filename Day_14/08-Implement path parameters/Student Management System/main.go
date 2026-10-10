package main

import (
	"student-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin router.
	router := gin.Default()

	// Register student route groups.
	routes.StudentRoutes(router)

	// Start the server on port 8080.
	router.Run(":8080")
}
