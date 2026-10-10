package controller

import (
	"net/http"
	"strconv"

	"employee-management/model"

	"github.com/gin-gonic/gin"
)

// Employees stores employee data in memory.
var Employees = []model.Employee{
	{ID: 1, Name: "Ganesh"},
	{ID: 2, Name: "Ravi"},
	{ID: 3, Name: "Priya"},
}

// DeleteEmployee deletes an employee by ID.
func DeleteEmployee(ctx *gin.Context) {

	// Read employee ID from the URL.
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid employee ID",
		})
		return
	}

	// Search for the employee.
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

	// Employee was not found.
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Employee not found",
	})
}
