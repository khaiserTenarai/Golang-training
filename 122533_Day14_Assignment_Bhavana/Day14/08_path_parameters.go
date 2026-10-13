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
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func askPort(reader *bufio.Reader) string {
	port := ask(reader, "Port to run the server on (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}
	return port
}

func askEmployees(reader *bufio.Reader) []Employee {
	count, err := strconv.Atoi(ask(reader, "How many employees do you want to start with? "))
	if err != nil || count < 0 {
		fmt.Println("That's not a valid number, starting with no employees.")
		count = 0
	}

	list := []Employee{}
	for i := 1; i <= count; i++ {
		fmt.Printf("\n--- Employee %d ---\n", i)
		name := ask(reader, "Name: ")
		email := ask(reader, "Email: ")
		department := ask(reader, "Department: ")
		salary, _ := strconv.ParseFloat(ask(reader, "Salary: "), 64)
		list = append(list, Employee{ID: i, Name: name, Email: email, Department: department, Salary: salary})
	}
	return list
}

type idParam struct {
	ID int `uri:"id" binding:"required"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	employees := askEmployees(reader)
	port := askPort(reader)

	router := gin.Default()

	router.GET("/employee/:id", func(c *gin.Context) {
		var p idParam
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}
		for _, e := range employees {
			if e.ID == p.ID {
				c.JSON(http.StatusOK, e)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
	})

	router.GET("/department/:name", func(c *gin.Context) {
		name := c.Param("name")
		matches := []Employee{}
		for _, e := range employees {
			if strings.EqualFold(e.Department, name) {
				matches = append(matches, e)
			}
		}
		c.JSON(http.StatusOK, matches)
	})

	router.GET("/department/:name/employee/:id", func(c *gin.Context) {
		name := c.Param("name")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}
		for _, e := range employees {
			if e.ID == id && strings.EqualFold(e.Department, name) {
				c.JSON(http.StatusOK, e)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "no such employee in that department"})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
