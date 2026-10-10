// Package controller handles employee HTTP requests.
package controller

import (
	"net/http"
	"strconv"
	"strings"

	"employee-gin-api/model"

	"github.com/gin-gonic/gin"
)

// Sample employee data used for this demonstration.
// Data resets when the application restarts.
var employees = []model.Employee{
	{ID: 1, Name: "Ravi", Age: 25, Email: "ravi@gmail.com"},
	{ID: 2, Name: "ganesh", Age: 21, Email: "ganesh@gmail.com"},
}

// GetAllEmployees returns all employees.
func GetAllEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

// GetEmployeeByID returns an employee by ID.
func GetEmployeeByID(c *gin.Context) {
	// Read the ID from the URL and convert it to an integer.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid employee ID",
		})
		return
	}

	// Search for the requested employee.
	for _, employee := range employees {
		if employee.ID == id {
			c.JSON(http.StatusOK, employee)
			return
		}
	}

	// Return 404 when the employee does not exist.
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Employee not found",
	})
}

// CreateEmployee creates a new employee.
func CreateEmployee(c *gin.Context) {
	var employee model.Employee

	// Convert the incoming JSON body into an Employee struct.
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON request",
		})
		return
	}

	// Validate the employee details.
	if strings.TrimSpace(employee.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	if employee.Age < 18 || employee.Age > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Age must be between 18 and 100",
		})
		return
	}

	if !strings.Contains(employee.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valid email is required",
		})
		return
	}

	// Generate the next ID using the existing employee records.
	nextID := 1
	for _, existing := range employees {
		if existing.ID >= nextID {
			nextID = existing.ID + 1
		}
	}

	// Assign the generated ID and save the employee.
	employee.ID = nextID
	employees = append(employees, employee)

	// Return HTTP 201 Created with the new employee.
	c.JSON(http.StatusCreated, employee)
}
