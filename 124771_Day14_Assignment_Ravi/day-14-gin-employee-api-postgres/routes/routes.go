package routes

import (
	"day-14-gin-employee-api/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, ctl *controller.EmployeeController) {
	ctl.RegisterRoutes(r)
}
