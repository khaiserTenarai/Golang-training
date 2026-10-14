package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/employees", func(c *gin.Context) {
        c.String(http.StatusCreated, "Employee created")
    })

    r.Run(":8080")
}
