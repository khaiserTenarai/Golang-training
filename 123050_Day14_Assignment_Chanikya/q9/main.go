package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/employees", func(c *gin.Context) {
        department := c.DefaultQuery("department", "all")
        limitText := c.DefaultQuery("limit", "10")
        limit, err := strconv.Atoi(limitText)
        if err != nil || limit < 1 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a positive integer"})
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "department": department,
            "limit": limit,
            "message": "Query parameters received",
        })
    })
    router.Run(":10014")
}
