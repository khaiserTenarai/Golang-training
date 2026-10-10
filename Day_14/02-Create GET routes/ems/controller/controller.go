// Package controller handles HTTP requests.
package controller

import (
	"net/http"
	"strconv"

	"employee-gin-api/model"

	"github.com/gin-gonic/gin"
)

// Sample employee data for demonstrating GET routes.
var employees = []model.Employee{
	{ID: 1, Name: "Ravi", Age: 25, Email: "ravi@gmail.com"},
	{ID: 2, Name: "Priya", Age: 27, Email: "priya@gmail.com"},
}

// GetAllEmployees returns all employees.
func GetAllEmployees(c *gin.Context) {
	// Send the employee slice as a JSON response.
	c.JSON(http.StatusOK, employees)
}

// GetEmployeeByID returns one employee using its ID.
func GetEmployeeByID(c *gin.Context) {
	// Read the ID from the URL.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid employee ID",
		})
		return
	}

	// Search for the employee with the requested ID.
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
