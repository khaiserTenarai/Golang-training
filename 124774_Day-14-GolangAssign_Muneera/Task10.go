package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.POST("/employees", func(c *gin.Context) {
		var e Employee

		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid JSON",
			})
			return
		}

		c.JSON(http.StatusCreated, e)
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees
// Method: POST

// Input (JSON Body):
// {
//   "name": "Raj",
//   "email": "raj@example.com",
//   "age": 25
// }

// Output:
// {
//   "name": "Raj",
//   "email": "raj@example.com",
//   "age": 25
// }
// Status: 201 Created
