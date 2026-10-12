package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type Employee struct {
    ID     int     `json:"id"`
    Name   string  `json:"name"`
    Email  string  `json:"email"`
    Age    int     `json:"age"`
    Salary float64 `json:"salary"`
}

var employees = map[int]Employee{
    1: {ID: 1, Name: "Asha", Email: "asha@example.com", Age: 25, Salary: 45000},
}

func main() {
    router := gin.Default()
    router.PUT("/employees/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
            return
        }
        var employee Employee
        if err := c.ShouldBindJSON(&employee); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if _, ok := employees[id]; !ok {
            c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
            return
        }
        employee.ID = id
        employees[id] = employee
        c.JSON(http.StatusOK, gin.H{"message": "Employee replaced", "employee": employee})
    })
    router.Run(":10014")
}
