package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/employees", func(c *gin.Context) {
        name := c.Query("name")
        department := c.DefaultQuery("department", "IT")

        c.JSON(http.StatusOK, gin.H{
            "name":       name,
            "department": department,
        })
    })

    r.Run(":8080")
}
