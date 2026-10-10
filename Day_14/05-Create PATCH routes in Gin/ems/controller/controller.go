package controller

import (
	"net/http"
	"strconv"

	"patch-demo/model"

	"github.com/gin-gonic/gin"
)

var employees = []model.Employee{
	{ID: 1, Name: "Ganesh", Age: 22, Salary: 30000},
	{ID: 2, Name: "Ravi", Age: 25, Salary: 40000},
}

// PatchEmployee updates only the supplied fields.
func PatchEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Pointers distinguish omitted fields from supplied values.
	var updates map[string]*string
	if err := c.ShouldBindJSON(&updates); err != nil || updates == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i := range employees {
		if employees[i].ID == id {
			for key, value := range updates {
				if value == nil {
					continue
				}

				switch key {
				case "name":
					employees[i].Name = *value
				case "age":
					age, err := strconv.Atoi(*value)
					if err != nil || age < 0 {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age"})
						return
					}
					employees[i].Age = age
				case "salary":
					salary, err := strconv.ParseFloat(*value, 64)
					if err != nil || salary < 0 {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid salary"})
						return
					}
					employees[i].Salary = salary
				default:
					c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown field: " + key})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{
				"message":  "Employee updated successfully",
				"employee": employees[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}
