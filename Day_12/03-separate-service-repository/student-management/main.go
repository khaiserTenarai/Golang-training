package main

import (
	"student-management/handler"
	"student-management/repository"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the repository.
	studentRepo := repository.NewMemoryRepository()

	// Connect the service to the repository.
	studentService := service.NewStudentService(studentRepo)

	// Connect the handler to the service.
	studentHandler := handler.NewStudentHandler(studentService)

	// Create the Gin router.
	router := gin.Default()

	// Register student routes.
	router.GET("/students", studentHandler.GetAll)
	router.POST("/students", studentHandler.Create)

	// Start the server.
	router.Run(":8080")
}
