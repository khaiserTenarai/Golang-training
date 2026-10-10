// Package routes registers the API endpoints.
package routes

import (
	"employee-gin-api/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all employee routes.
func RegisterRoutes(router *gin.Engine) {
	// GET: retrieve all employees.
	router.GET("/employees", controller.GetAllEmployees)

	// GET: retrieve one employee by ID.
	router.GET("/employees/:id", controller.GetEmployeeByID)

	// POST: create a new employee.
	router.POST("/employees", controller.CreateEmployee)
}
