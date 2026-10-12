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
    router.PATCH("/employees/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
            return
        }
        employee, ok := employees[id]
        if !ok {
            c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
            return
        }
        var changes map[string]interface{}
        if err := c.ShouldBindJSON(&changes); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if name, ok := changes["name"].(string); ok { employee.Name = name }
        if email, ok := changes["email"].(string); ok { employee.Email = email }
        if age, ok := changes["age"].(float64); ok { employee.Age = int(age) }
        if salary, ok := changes["salary"].(float64); ok { employee.Salary = salary }
        employees[id] = employee
        c.JSON(http.StatusOK, gin.H{"message": "Employee partially updated", "employee": employee})
    })
    router.Run(":10014")
}
