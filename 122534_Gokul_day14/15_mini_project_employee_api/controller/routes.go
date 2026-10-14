package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, employees *EmployeeController) {
	router.HandleMethodNotAllowed = true

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found", "path": c.Request.URL.Path})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed", "method": c.Request.Method})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api/v1")
	{
		api.POST("/employees", employees.Create)
		api.GET("/employees", employees.GetAll)
		api.GET("/employees/:id", employees.GetOne)
		api.PUT("/employees/:id", employees.Replace)
		api.PATCH("/employees/:id", employees.Patch)
		api.DELETE("/employees/:id", employees.Delete)
	}
}
