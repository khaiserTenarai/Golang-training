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

	// Register global middleware.
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Logging())

	// All student routes require authentication.
	students := router.Group("/students")
	students.Use(middleware.Authentication())

	// Both students and admins can view students.
	students.GET("", studentHandler.GetAll)

	// Only admins can create students.
	students.POST(
		"",
		middleware.Authorization("admin"),
		studentHandler.Create,
	)

	// Start the server.
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
