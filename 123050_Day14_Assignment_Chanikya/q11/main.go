package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/success", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Request succeeded"})
    })
    router.GET("/created", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{"message": "Resource created"})
    })
    router.GET("/error", func(c *gin.Context) {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
    })
    router.Run(":10014")
}
