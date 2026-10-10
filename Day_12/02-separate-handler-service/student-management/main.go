package main

import (
	"student-management/handler"
	"student-management/repository"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create and connect the layers.
	studentRepo := repository.NewStudentRepository()
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// Register API routes.
	router := gin.Default()

	router.GET("/students", studentHandler.GetAll)
	router.POST("/students", studentHandler.Create)

	// Start the server.
	router.Run(":8080")
}
