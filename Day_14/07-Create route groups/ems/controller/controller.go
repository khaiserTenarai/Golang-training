package controller

import (
	"net/http"
	"strconv"

	"employee-management/model"

	"github.com/gin-gonic/gin"
)

// Employees stores employee data in memory.
var Employees = []model.Employee{
	{ID: 1, Name: "Ganesh", Age: 23, Email: "ganesh@example.com"},
	{ID: 2, Name: "Ravi", Age: 25, Email: "ravi@example.com"},
}

// GetEmployees returns all employees.
func GetEmployees(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Employees)
}

// GetEmployee returns one employee by ID.
func GetEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	for _, employee := range Employees {
		if employee.ID == id {
			ctx.JSON(http.StatusOK, employee)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

// CreateEmployee adds a new employee.
func CreateEmployee(ctx *gin.Context) {
	var employee model.Employee

	// Convert the JSON request body into an Employee struct.
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validate the required employee details.
	if employee.Name == "" || employee.Age <= 0 || employee.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, valid age, and email are required",
		})
		return
	}

	// Generate the next ID automatically.
	employee.ID = 1
	for _, existing := range Employees {
		if existing.ID >= employee.ID {
			employee.ID = existing.ID + 1
		}
	}

	Employees = append(Employees, employee)

	ctx.JSON(http.StatusCreated, employee)
}

// UpdateEmployee updates an employee by ID.
func UpdateEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var updated model.Employee
	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if updated.Name == "" || updated.Age <= 0 || updated.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, valid age, and email are required",
		})
		return
	}

	for i, employee := range Employees {
		if employee.ID == id {
			// Keep the original ID from the URL.
			updated.ID = id
			Employees[i] = updated

			ctx.JSON(http.StatusOK, updated)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

// DeleteEmployee removes an employee by ID.
func DeleteEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	for i, employee := range Employees {
		if employee.ID == id {
			// Remove the employee from the slice.
			Employees = append(Employees[:i], Employees[i+1:]...)

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Employee deleted successfully",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}
