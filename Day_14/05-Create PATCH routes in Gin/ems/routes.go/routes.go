package routes

import (
	"patch-demo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.PATCH("/employees/:id", controller.PatchEmployee)
}
