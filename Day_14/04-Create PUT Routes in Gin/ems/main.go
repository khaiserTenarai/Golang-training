package main

import (
	"put-demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router.
	router := gin.Default()

	// Register routes.
	routes.SetupRoutes(router)

	// Start server.
	router.Run(":8080")
}
