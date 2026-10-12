package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.GET("/employee", func(c *gin.Context) {
		employee := Employee{
			ID:    1,
			Name:  "Raj",
			Email: "raj@example.com",
		}

		c.JSON(http.StatusOK, employee)
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employee
// Method: GET

// Output:
// {
//   "id": 1,
//   "name": "Raj",
//   "email": "raj@example.com"
// }
// Status: 200 OK
