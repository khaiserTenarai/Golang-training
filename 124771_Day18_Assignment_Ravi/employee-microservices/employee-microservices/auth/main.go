package main

import (
	"employee-microservices/auth/handler"
	"employee-microservices/auth/repository"
	"employee-microservices/auth/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	repo := repository.NewUserRepository()
	svc := service.NewAuthService(repo)
	h := handler.NewAuthHandler(svc)
	r := gin.Default()
	h.RegisterRoutes(r)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "auth-service", "status": "ok"}) })
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8083"
	}
	log.Printf("auth-service starting on :%s", port)
	log.Fatal(r.Run(":" + port))
}
