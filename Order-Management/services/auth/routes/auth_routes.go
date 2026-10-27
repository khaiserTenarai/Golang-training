package routes

import (
	"net/http"

	"order-management/services/auth/controller"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, authController *controller.AuthController) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "auth-service",
		})
	})

	router.POST("/auth/login", authController.Login)
}
