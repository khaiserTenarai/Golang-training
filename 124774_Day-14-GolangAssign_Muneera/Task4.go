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

	r.PUT("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")

		var employee Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid JSON",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"employee": employee,
			"message":  "Employee replaced",
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees/1
// Method: PUT

// Input (JSON Body):
// {
//   "name": "Raj",
//   "email": "raj@example.com"
// }

// Output:
// {
//   "id": "1",
//   "employee": {
//     "name": "Raj",
//     "email": "raj@example.com"
//   },
//   "message": "Employee replaced"
// }
// Status: 200 OK
