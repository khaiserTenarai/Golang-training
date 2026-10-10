package controller

import (
	"net/http"
	"strconv"

	"put-demo/model"

	"github.com/gin-gonic/gin"
)

// Temporary employee data.
var employees = []model.Employee{
	{ID: 1, Name: "Ganesh", Age: 22, Salary: 30000},
	{ID: 2, Name: "Ravi", Age: 25, Salary: 40000},
}

// UpdateEmployee updates an existing employee.
func UpdateEmployee(c *gin.Context) {
	// Read employee ID from URL.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Read updated details from JSON.
	var updatedEmployee model.Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Find and update the employee.
	for i, employee := range employees {
		if employee.ID == id {
			updatedEmployee.ID = id
			employees[i] = updatedEmployee

			c.JSON(http.StatusOK, gin.H{
				"message":  "Employee updated successfully",
				"employee": employees[i],
			})
			return
		}
	}

	// Employee does not exist.
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}
