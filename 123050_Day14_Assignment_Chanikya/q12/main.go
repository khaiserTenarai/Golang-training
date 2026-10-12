package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.POST("/contact", func(c *gin.Context) {
        name := c.PostForm("name")
        email := c.PostForm("email")
        if name == "" || email == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "name and email are required"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Form received", "name": name, "email": email})
    })
    router.Run(":10014")
}
