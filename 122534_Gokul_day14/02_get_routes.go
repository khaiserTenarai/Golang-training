/*
Question 2: Create GET routes.

Try: curl localhost:8080/about
     curl localhost:8080/employees
     curl localhost:8080/employees/count
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

var employees = []Employee{
	{1, "Asha", 30000},
	{2, "Ravi", 45000},
	{3, "Meena", 38000},
}

func readPort() string {
	fmt.Print("Enter the port (press Enter for 8080): ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return "8080"
	}
	return text
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Employee service is running")
	})

	router.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "A small service that lists employees")
	})

	router.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	router.GET("/employees/count", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"count": len(employees)})
	})

	router.Run(":" + readPort())
}
