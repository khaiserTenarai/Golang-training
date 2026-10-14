package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/employee", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "id":     1,
            "name":   "Abirami",
            "salary": 40000,
        })
    })

    r.Run(":8080")
}
