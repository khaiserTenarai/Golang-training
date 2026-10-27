package routes

import (
	"order-management/services/gateway/controller"
	gatewaymiddleware "order-management/services/gateway/middleware"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, c *controller.GatewayController) {
	router.GET("/health", c.Health)

	// Login is public. The Auth Service validates credentials.
	router.POST("/auth/login", c.Auth)

	auth := gatewaymiddleware.RequireAuth()
	admin := gatewaymiddleware.RequireRole("admin")

	// Product reads: admin and user.
	router.GET("/products", auth, c.Product)
	router.GET("/products/*path", auth, c.Product)

	// Product modifications: admin only.
	router.POST("/products", auth, admin, c.Product)
	router.POST("/products/*path", auth, admin, c.Product)
	router.PUT("/products", auth, admin, c.Product)
	router.PUT("/products/*path", auth, admin, c.Product)
	router.PATCH("/products", auth, admin, c.Product)
	router.PATCH("/products/*path", auth, admin, c.Product)
	router.DELETE("/products", auth, admin, c.Product)
	router.DELETE("/products/*path", auth, admin, c.Product)

	// Orders: authenticated users can create and view orders.
	router.POST("/orders", auth, c.Order)
	router.GET("/orders", auth, c.Order)
	router.GET("/orders/*path", auth, c.Order)

	// Only admins can confirm or cancel an order.
	router.PATCH("/orders/:id/status", auth, admin, c.Order)
}
