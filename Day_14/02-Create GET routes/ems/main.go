// Package main starts the Gin application.
package main

import (
	"employee-gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with logging and recovery middleware.
	router := gin.Default()

	// Register the GET routes.
	routes.RegisterRoutes(router)

	// Start the server on port 8080.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
