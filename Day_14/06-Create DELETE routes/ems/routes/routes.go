package routes

import (
	"employee-management/controller"

	"github.com/gin-gonic/gin"
)

// EmployeeRoutes registers employee routes.
func EmployeeRoutes(router *gin.Engine) {

	// DELETE route to delete an employee by ID.
	router.DELETE("/employees/:id", controller.DeleteEmployee)
}
