package routes

import (
	"net/http"

	"employee-management-microservices/services/department/controller"

	"github.com/gin-gonic/gin"
)

func RegisterDepartmentRoutes(router *gin.Engine, c *controller.DepartmentController) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "UP", "service": "department"})
	})

	api := router.Group("/api/departments")
	{
		api.POST("", c.Create)
		api.GET("", c.GetAll)
		api.GET("/:id", c.GetByID)
		api.PUT("/:id", c.Update)
		api.DELETE("/:id", c.Delete)
	}
}
