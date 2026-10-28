package routes

import (
	"net/http"

	"employee-management-microservices/services/employee/controller"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(router *gin.Engine, c *controller.EmployeeController) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "UP", "service": "employee"})
	})

	api := router.Group("/api/employees")
	{
		api.POST("", c.Create)
		api.GET("", c.GetAll)
		api.GET("/:id", c.GetByID)
		api.PUT("/:id", c.Update)
		api.DELETE("/:id", c.Delete)
	}
}
