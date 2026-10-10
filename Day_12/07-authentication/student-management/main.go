package main

import (
	"student-management/handler"
	"student-management/middleware"
	"student-management/repository"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect the layers.
	studentRepo := repository.NewMemoryRepository()
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// Create the router.
	router := gin.New()

	// Register middleware in execution order.
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Logging())

	// Register routes.
	router.GET("/students", studentHandler.GetAll)
	router.POST("/students", studentHandler.Create)

	// Start the server.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
