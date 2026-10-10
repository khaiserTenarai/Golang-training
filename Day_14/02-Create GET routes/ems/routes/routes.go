// Package routes registers API endpoints.
package routes

import (
	"employee-gin-api/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all GET routes.
func RegisterRoutes(router *gin.Engine) {
	// GET route to retrieve all employees.
	router.GET("/employees", controller.GetAllEmployees)

	// GET route to retrieve one employee by ID.
	router.GET("/employees/:id", controller.GetEmployeeByID)
}
