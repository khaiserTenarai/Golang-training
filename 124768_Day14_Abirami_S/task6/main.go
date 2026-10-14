package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.DELETE("/employees/:id", func(c *gin.Context) {
        c.String(http.StatusOK, "Employee "+c.Param("id")+" deleted")
    })

    r.Run(":8080")
}
