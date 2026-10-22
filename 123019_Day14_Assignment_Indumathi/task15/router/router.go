package router

import (
	"net/http"

	"employee-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(empCtrl *controllers.EmployeeController) *gin.Engine {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/employees", empCtrl.GetAll)
		api.GET("/employees/:id", empCtrl.GetByID)
		api.POST("/employees", empCtrl.Create)
		api.PUT("/employees/:id", empCtrl.Update)
		api.PATCH("/employees/:id", empCtrl.Patch)
		api.DELETE("/employees/:id", empCtrl.Delete)
	}

	return r
}
