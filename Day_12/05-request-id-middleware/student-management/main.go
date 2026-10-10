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
	router := gin.Default()

	// Register request-ID middleware.
	router.Use(middleware.RequestID())

	// Register routes.
	router.GET("/students", studentHandler.GetAll)
	router.POST("/students", studentHandler.Create)

	// Start the server.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
