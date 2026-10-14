/*
Question 6: Create DELETE routes.

Try: curl -X DELETE localhost:8080/employees/2
     curl localhost:8080/employees
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

func findIndex(id int) int {
	for i, e := range employees {
		if e.ID == id {
			return i
		}
	}
	return -1
}

func main() {
	router := gin.Default()

	router.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	router.DELETE("/employees/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}

		pos := findIndex(id)
		if pos == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}

		employees = append(employees[:pos], employees[pos+1:]...)
		c.Status(http.StatusNoContent)
	})

	router.Run(":" + readPort())
}
