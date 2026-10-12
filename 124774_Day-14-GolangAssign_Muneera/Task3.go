package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/employees", func(c *gin.Context) {
		var employee Employee

		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid JSON",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Employee created",
			"employee": employee,
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees
// Method: POST
// Input (JSON Body):
// {
//   "name": "Raj",
//   "email": "raj@example.com"
// }

// Output:
// {
//   "message": "Employee created",
//   "employee": {
//     "name": "Raj",
//     "email": "raj@example.com"
//   }
// }
// Status: 201 Created
