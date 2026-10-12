package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"message": "Employee found",
			"id":      id,
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees/101
// Method: GET

// Output:
// {
//   "message": "Employee found",
//   "id": "101"
// }
