package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{AllowOrigins: []string{"http://localhost:3000"}, AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}, AllowCredentials: true}))
	r.GET("/hello", func(c *gin.Context) { c.JSON(200, gin.H{"message": "CORS enabled"}) })
	r.Run(":8080")
}
