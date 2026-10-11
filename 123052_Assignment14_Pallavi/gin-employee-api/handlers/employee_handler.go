package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"gin-employee-api/models"

	"github.com/gin-gonic/gin"
)

var employees = map[string]models.Employee{
	"101": {ID: "101", Name: "Pallavi", Role: "Backend Engineer", Department: "IT", Salary: 75000},
	"102": {ID: "102", Name: "Khaiser", Role: "Tech Lead", Department: "Engineering", Salary: 95000},
}

// Topic 2 & 9: GET with Query Parameters
func GetEmployees(c *gin.Context) {
	dept := c.Query("department") // Query parameter
	if dept != "" {
		filtered := []models.Employee{}
		for _, emp := range employees {
			if strings.EqualFold(emp.Department, dept) {
				filtered = append(filtered, emp)
			}
		}
		c.JSON(http.StatusOK, gin.H{"data": filtered})
		return
	}

	all := make([]models.Employee, 0, len(employees))
	for _, emp := range employees {
		all = append(all, emp)
	}
	c.JSON(http.StatusOK, gin.H{"data": all})
}

// Topic 8 & 11: GET with Path Parameter & JSON Response
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id") // Path parameter
	emp, exists := employees[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": emp})
}

// Topic 3 & 10: POST Route with JSON Binding
func CreateEmployee(c *gin.Context) {
	var newEmp models.Employee
	if err := c.ShouldBindJSON(&newEmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := employees[newEmp.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Employee ID already exists"})
		return
	}

	employees[newEmp.ID] = newEmp
	c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully", "data": newEmp})
}

// Topic 4: PUT Route (Full Update)
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	if _, exists := employees[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var updatedEmp models.Employee
	if err := c.ShouldBindJSON(&updatedEmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEmp.ID = id
	employees[id] = updatedEmp
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated fully", "data": updatedEmp})
}

// Topic 5: PATCH Route (Partial Update)
func PartialUpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	emp, exists := employees[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var updates models.UpdateEmployee
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updates.Name != "" {
		emp.Name = updates.Name
	}
	if updates.Role != "" {
		emp.Role = updates.Role
	}
	if updates.Department != "" {
		emp.Department = updates.Department
	}
	if updates.Salary > 0 {
		emp.Salary = updates.Salary
	}

	employees[id] = emp
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated partially", "data": emp})
}

// Topic 6: DELETE Route
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if _, exists := employees[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	delete(employees, id)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// Topic 12: Handle Form Parameters
func ProcessForm(c *gin.Context) {
	name := c.PostForm("name")
	role := c.PostForm("role")
	salaryStr := c.PostForm("salary")

	salary, _ := strconv.ParseFloat(salaryStr, 64)

	c.JSON(http.StatusOK, gin.H{
		"message": "Form processed successfully",
		"received": gin.H{
			"name":   name,
			"role":   role,
			"salary": salary,
		},
	})
}