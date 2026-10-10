package main

import (
	"student-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register student routes.
	routes.StudentRoutes(router)

	// Start the server.
	router.Run(":8080")
}
