/*
Question 5: Create PATCH routes.

PATCH changes only the fields that are sent.
Try: curl -X PATCH localhost:8080/employees/1 -d '{"salary":40000}'
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

type EmployeePatch struct {
	Name   *string  `json:"name"`
	Salary *float64 `json:"salary"`
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

	router.PATCH("/employees/:id", func(c *gin.Context) {
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

		var patch EmployeePatch
		if err := c.ShouldBindJSON(&patch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if patch.Name != nil {
			employees[pos].Name = *patch.Name
		}
		if patch.Salary != nil {
			employees[pos].Salary = *patch.Salary
		}
		c.JSON(http.StatusOK, employees[pos])
	})

	router.Run(":" + readPort())
}
