package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Employee struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Salary int    `json:"salary"`
}

func main() {
    r := gin.Default()

    r.POST("/employees", func(c *gin.Context) {
        var employee Employee

        if err := c.ShouldBindJSON(&employee); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, employee)
    })

    r.Run(":8080")
}
