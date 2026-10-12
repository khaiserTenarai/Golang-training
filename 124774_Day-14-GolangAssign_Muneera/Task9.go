package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/employees", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		size := c.DefaultQuery("size", "10")

		fmt.Fprintf(c.Writer, "page=%s size=%s", page, size)
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/employees?page=2&size=5
// Method: GET
// Output:
// page=2 size=5

// Request URL: http://localhost:8080/employees
// Method: GET
// Output (default values):
// page=1 size=10
