package main

import (
	"student-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.StudentRoutes(router)

	router.Run(":8081")
}

// http://localhost:8080/api/students?name=Ganesh&age=21
