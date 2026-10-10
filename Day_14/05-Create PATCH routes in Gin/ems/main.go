package main

import (
	// "patch-demo/routes"

	"patch-demo/routes.go"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
