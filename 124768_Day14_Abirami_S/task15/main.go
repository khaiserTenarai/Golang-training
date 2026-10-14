package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Employee struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Department string `json:"department"`
    Salary     int    `json:"salary"`
}

var employees = []Employee{
    {ID: 1, Name: "Abirami", Department: "IT", Salary: 40000},
    {ID: 2, Name: "Priya", Department: "HR", Salary: 35000},
}

func getEmployeeIndex(id int) int {
    for i, employee := range employees {
        if employee.ID == id {
            return i
        }
    }
    return -1
}

func getAllEmployees(c *gin.Context) {
    c.JSON(http.StatusOK, employees)
}

func getEmployee(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
        return
    }

    index := getEmployeeIndex(id)
    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    c.JSON(http.StatusOK, employees[index])
}

func createEmployee(c *gin.Context) {
    var employee Employee

    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    employee.ID = 1
    if len(employees) > 0 {
        employee.ID = employees[len(employees)-1].ID + 1
    }

    employees = append(employees, employee)

    c.JSON(http.StatusCreated, employee)
}

func updateEmployee(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
        return
    }

    index := getEmployeeIndex(id)
    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    var employee Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    employee.ID = id
    employees[index] = employee

    c.JSON(http.StatusOK, employee)
}

func patchEmployee(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
        return
    }

    index := getEmployeeIndex(id)
    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    var updates map[string]interface{}
    if err := c.ShouldBindJSON(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if name, ok := updates["name"].(string); ok {
        employees[index].Name = name
    }
    if department, ok := updates["department"].(string); ok {
        employees[index].Department = department
    }
    if salary, ok := updates["salary"].(float64); ok {
        employees[index].Salary = int(salary)
    }

    c.JSON(http.StatusOK, employees[index])
}

func deleteEmployee(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
        return
    }

    index := getEmployeeIndex(id)
    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    deleted := employees[index]
    employees = append(employees[:index], employees[index+1:]...)

    c.JSON(http.StatusOK, gin.H{
        "message":  "Employee deleted successfully",
        "employee": deleted,
    })
}

func main() {
    r := gin.Default()

    employeeGroup := r.Group("/employees")
    {
        employeeGroup.GET("", getAllEmployees)
        employeeGroup.GET("/:id", getEmployee)
        employeeGroup.POST("", createEmployee)
        employeeGroup.PUT("/:id", updateEmployee)
        employeeGroup.PATCH("/:id", patchEmployee)
        employeeGroup.DELETE("/:id", deleteEmployee)
    }

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })

    r.NoMethod(func(c *gin.Context) {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
    })

    r.Run(":8080")
}
