package main

import (
	"employee-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin router.
	router := gin.Default()

	// Register employee routes.
	routes.EmployeeRoutes(router)

	// Start the server on port 8080.
	router.Run(":8081")
}
