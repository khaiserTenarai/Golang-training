package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Employee struct {
    ID     int     `json:"id"`
    Name   string  `json:"name"`
    Email  string  `json:"email"`
    Age    int     `json:"age"`
    Salary float64 `json:"salary"`
}

func main() {
    router := gin.Default()

    router.POST("/employees", func(c *gin.Context) {
        var employee Employee
        if err := c.ShouldBindJSON(&employee); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{
            "message": "Employee accepted (in-memory example; not persisted)",
            "employee": employee,
        })
    })

    router.Run(":10014")
}
