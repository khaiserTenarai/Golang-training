/*
Question 3: Create POST routes.

Try: curl -X POST localhost:8080/employees -d '{"name":"Karthik","salary":52000}'
     curl localhost:8080/employees
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

	router.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	router.POST("/employees", func(c *gin.Context) {
		var e Employee
		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		e.ID = len(employees) + 1
		employees = append(employees, e)
		c.JSON(http.StatusCreated, e)
	})

	router.Run(":" + readPort())
}
