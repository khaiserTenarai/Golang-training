package routes

import (
	"put-demo/controller"

	"github.com/gin-gonic/gin"
)

// SetupRoutes registers employee routes.
func SetupRoutes(router *gin.Engine) {
	router.PUT("/employees/:id", controller.UpdateEmployee)
}
