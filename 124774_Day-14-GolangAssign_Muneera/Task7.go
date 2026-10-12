package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "employee list",
		})
	})

	v1.GET("/departments", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "department list",
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/api/v1/employees
// Method: GET
// Output:
// {
//   "message": "employee list"
// }

// Request URL: http://localhost:8080/api/v1/departments
// Method: GET
// Output:
// {
//   "message": "department list"
// }
