package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/status", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"status": "ok"})
        })
        employees := api.Group("/employees")
        {
            employees.GET("", func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{"message": "Employee list"})
            })
            employees.POST("", func(c *gin.Context) {
                c.JSON(http.StatusCreated, gin.H{"message": "Employee created"})
            })
        }
    }

    router.Run(":10014")
}
