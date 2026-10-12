package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.PATCH("/employees/:id", func(c *gin.Context) {
		var updates map[string]interface{}

		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid JSON",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      c.Param("id"),
			"updates": updates,
			"message": "Employee partially updated",
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees/1
// Method: PATCH

// Input (JSON Body):
// {
//   "name": "Raj Kumar"
// }

// Output:
// {
//   "id": "1",
//   "updates": {
//     "name": "Raj Kumar"
//   },
//   "message": "Employee partially updated"
// }
// Status: 200 OK
