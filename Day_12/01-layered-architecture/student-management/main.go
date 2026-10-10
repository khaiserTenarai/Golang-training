package main

import (
	"student-management/handler"
	"student-management/repository"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect the layers.
	studentRepo := repository.NewStudentRepository()
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// Register routes.
	router := gin.Default()

	router.GET("/students", studentHandler.GetAll)
	router.POST("/students", studentHandler.Create)

	// Start the server.
	router.Run(":8080")
}
