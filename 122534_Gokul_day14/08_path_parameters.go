/*
Question 8: Implement path parameters.

Try: curl localhost:8080/employees/2
     curl localhost:8080/departments/Tech/employees/2
     curl localhost:8080/files/reports/2024/jan.txt
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
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var employees = []Employee{
	{1, "Asha", "Sales", 30000},
	{2, "Ravi", "Tech", 45000},
	{3, "Meena", "HR", 38000},
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

	router.GET("/employees/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}

		for _, e := range employees {
			if e.ID == id {
				c.JSON(http.StatusOK, e)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
	})

	router.GET("/departments/:dept/employees/:id", func(c *gin.Context) {
		dept := c.Param("dept")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}

		for _, e := range employees {
			if e.ID == id && strings.EqualFold(e.Department, dept) {
				c.JSON(http.StatusOK, e)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "no such employee in " + dept})
	})

	router.GET("/files/*path", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"requested_path": c.Param("path")})
	})

	router.Run(":" + readPort())
}
