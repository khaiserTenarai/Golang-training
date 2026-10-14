package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/form", func(c *gin.Context) {
        name := c.PostForm("name")
        email := c.PostForm("email")

        c.JSON(http.StatusOK, gin.H{
            "name":  name,
            "email": email,
        })
    })

    r.Run(":8080")
}
