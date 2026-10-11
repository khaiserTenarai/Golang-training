package main

import (
	"net/http"

	"gin-employee-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Topic 1: Create a Gin application
	r := gin.Default()

	// Topic 13: Serve static files
	r.Static("/static", "./static")

	// Topic 14: Custom 404 & 405 Handlers
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Custom 404: Endpoint Route Not Found",
		})
	})

	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    405,
			"message": "Custom 405: HTTP Method Not Allowed on this Route",
		})
	})

	// Topic 12: Form Parameter submission endpoint
	r.POST("/submit-form", handlers.ProcessForm)

	// Topic 7 & 15: Route Groups & Employee CRUD REST API
	v1 := r.Group("/api/v1")
	{
		v1.GET("/employees", handlers.GetEmployees)               // GET all or filter by ?department=IT
		v1.GET("/employees/:id", handlers.GetEmployeeByID)         // GET by Path Param
		v1.POST("/employees", handlers.CreateEmployee)             // POST create
		v1.PUT("/employees/:id", handlers.UpdateEmployee)          // PUT replace
		v1.PATCH("/employees/:id", handlers.PartialUpdateEmployee) // PATCH update field
		v1.DELETE("/employees/:id", handlers.DeleteEmployee)       // DELETE remove
	}

	// Start server on port 8080
	r.Run(":8080")
}