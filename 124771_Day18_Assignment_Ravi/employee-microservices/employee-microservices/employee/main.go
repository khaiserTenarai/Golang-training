package main

import (
	"log"
	"os"

	"employee-microservices/employee/handler"
	"employee-microservices/employee/repository"
	"employee-microservices/employee/service"
	"employee-microservices/shared/observability"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := observability.NewLogger("employee-service") // Task 14: basic logging.
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo) // Dependency injection: repository -> service.
	h := handler.NewEmployeeHandler(svc)    // Dependency injection: service -> handler.

	r := gin.Default()
	h.RegisterRoutes(r)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "employee-service", "status": "ok"}) })
	port := os.Getenv("EMPLOYEE_PORT")
	if port == "" {
		port = "8081"
	}
	logger.Println("starting on :" + port)
	log.Fatal(r.Run(":" + port))
}
