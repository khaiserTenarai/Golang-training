package routes

import (
	"employee-management/controller"

	"github.com/gin-gonic/gin"
)

// EmployeeRoutes registers all employee routes.
func EmployeeRoutes(router *gin.Engine) {

	// All employee endpoints start with /api/employees.
	employee := router.Group("/api/employees")

	employee.GET("", controller.GetEmployees)
	employee.GET("/:id", controller.GetEmployee)
	employee.POST("", controller.CreateEmployee)
	employee.PUT("/:id", controller.UpdateEmployee)
	employee.DELETE("/:id", controller.DeleteEmployee)
}
