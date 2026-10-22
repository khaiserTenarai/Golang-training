package main

import (
	"gin-layered-api/controllers"
	"gin-layered-api/middlewares"
	"gin-layered-api/repositories"
	"gin-layered-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepo := repositories.NewUserRepository()
	productRepo := repositories.NewProductRepository()

	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)

	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)

	r := gin.Default()

	r.POST("/login", authController.Login)

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/products", productController.GetProducts)
		api.POST("/products", middlewares.RBACMiddleware("admin"), productController.CreateProduct)
	}

	r.Run(":8080")
}
