package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.DELETE("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": "Employee deleted",
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees/1
// Method: DELETE

// Output:
// {
//   "id": "1",
//   "message": "Employee deleted"
// }
// Status: 200 OK
