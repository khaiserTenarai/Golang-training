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

	// Create the router.
	router := gin.Default()

	// Register routes.
	router.GET("/students", studentHandler.GetAll)
	router.POST("/students", studentHandler.Create)

	// Start the server.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
