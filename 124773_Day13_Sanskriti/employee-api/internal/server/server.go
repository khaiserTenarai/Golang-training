package server

import (
	"employee-api/internal/controller"
	"employee-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctrl *controller.EmployeeController) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/employees", ctrl.Create)
		v1.GET("/employees/:id", ctrl.GetByID)
	}

	return r
}
