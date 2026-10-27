package main

import (
	gatewayhandler "employee-microservices/gateway/handler"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	gatewayhandler.NewGatewayHandler().RegisterRoutes(r)
	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("api-gateway starting on :%s", port)
	log.Fatal(r.Run(":" + port))
}
