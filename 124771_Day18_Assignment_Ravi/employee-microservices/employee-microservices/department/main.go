package main

import (
	"employee-microservices/department/handler"
	"employee-microservices/department/repository"
	"employee-microservices/department/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	repo := repository.NewDepartmentRepository()
	svc := service.NewDepartmentService(repo)
	h := handler.NewDepartmentHandler(svc)
	r := gin.Default()
	h.RegisterRoutes(r)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "department-service", "status": "ok"}) })
	port := os.Getenv("DEPARTMENT_PORT")
	if port == "" {
		port = "8082"
	}
	log.Printf("department-service starting on :%s", port)
	log.Fatal(r.Run(":" + port))
}
