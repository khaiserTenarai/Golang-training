package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    employeeGroup := r.Group("/employees")
    {
        employeeGroup.GET("/", func(c *gin.Context) {
            c.String(http.StatusOK, "All employees")
        })
        employeeGroup.GET("/:id", func(c *gin.Context) {
            c.String(http.StatusOK, "Employee "+c.Param("id"))
        })
    }

    r.Run(":8080")
}
