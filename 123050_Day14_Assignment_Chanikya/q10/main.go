package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Employee struct {
    Name   string  `json:"name" binding:"required"`
    Email  string  `json:"email" binding:"required,email"`
    Age    int     `json:"age" binding:"required,min=18"`
    Salary float64 `json:"salary" binding:"required,gt=0"`
}

func main() {
    router := gin.Default()
    router.POST("/employees", func(c *gin.Context) {
        var employee Employee
        if err := c.ShouldBindJSON(&employee); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"employee": employee})
    })
    router.Run(":10014")
}
