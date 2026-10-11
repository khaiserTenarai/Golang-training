package main

import (
	"os"

	"advanced-gin-api/handlers"
	"advanced-gin-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	_ = os.MkdirAll("uploads", os.ModePerm)

	r := gin.New()

	// Apply Middlewares globally (Topics 1, 2, 5, 6, 7)
	r.Use(middleware.ExecutionOrderLogger("M1-Global"))
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.CustomRecovery())
	r.Use(middleware.CORS())

	// Public Routes
	r.POST("/api/login", handlers.Login)

	// Protected Routes (Topic 9: JWT Authentication & Topic 10: RBAC)
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// Accessible to both User and Admin
		protected.GET("/employees", middleware.RequireRole("Admin", "User"), handlers.GetEmployeesPaginated)
		protected.POST("/upload", middleware.RequireRole("Admin", "User"), handlers.UploadDocument)

		// Restrict creation strictly to Admin role (Topic 10: RBAC)
		protected.POST("/employees", middleware.RequireRole("Admin"), handlers.CreateEmployee)
	}

	r.Run(":8080")
}