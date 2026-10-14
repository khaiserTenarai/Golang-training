package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.PUT("/employees/:id", func(c *gin.Context) {
        c.String(http.StatusOK, "Employee "+c.Param("id")+" replaced")
    })

    r.Run(":8080")
}
