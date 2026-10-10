// Package main starts the Gin application.
package main

import (
	"employee-gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin router with logging and recovery middleware.
	router := gin.Default()

	// Register GET and POST routes.
	routes.RegisterRoutes(router)

	// Start the server on port 8080.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
