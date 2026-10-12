package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin"})
    })
    router.GET("/employees", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "List of employees"})
    })
    router.GET("/employees/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"employee_id": c.Param("id")})
    })

    router.Run(":10014")
}
