package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.New()
    router.Use(gin.Logger(), gin.Recovery())

    router.GET("/employees", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Employee list"})
    })

    // Custom 404: route not found.
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })

    // Custom 405: path exists, but HTTP method is not allowed.
    router.HandleMethodNotAllowed = true
    router.NoMethod(func(c *gin.Context) {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "HTTP method not allowed"})
    })

    router.Run(":10014")
}
