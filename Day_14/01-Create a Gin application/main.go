// Package main starts the Gin application.
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with logging and recovery middleware.
	router := gin.Default()

	// Register a simple GET endpoint.
	router.GET("/", func(c *gin.Context) {
		// Send a JSON response with HTTP status 200.
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Employee Management API",
		})
	})

	// Start the server on port 8080.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

/*
go mod init employee-gin-api
go get github.com/gin-gonic/gin
go run .


Open this URL in your browser:

http://localhost:8080/
*/
