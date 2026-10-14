package main

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(LoggerMiddleware())

	// Public routes
	r.POST("/register", Register)
	r.POST("/login", Login)

	// Routes for authenticated users
	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/tasks", CreateTask)
		protected.GET("/tasks", ListTasks)
		protected.DELETE("/tasks/:id", DeleteTask)
	}

	// Admin-only routes
	admin := r.Group("/admin")
	admin.Use(AuthMiddleware(), RoleMiddleware("admin"))
	{
		admin.GET("/users", ListUsers)
	}

	return r
}
