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

func main() {
	reader := bufio.NewReader(os.Stdin)
	employees := askEmployees(reader)
	port := askPort(reader)

	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		department := c.Query("department")
		minSalary, _ := strconv.ParseFloat(c.Query("min_salary"), 64)
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}
		_, filtered := c.GetQuery("department")

		matches := []Employee{}
		for _, e := range employees {
			if department != "" && !strings.EqualFold(e.Department, department) {
				continue
			}
			if minSalary > 0 && e.Salary < minSalary {
				continue
			}
			matches = append(matches, e)
		}

		start := (page - 1) * limit
		if start > len(matches) {
			start = len(matches)
		}
		end := start + limit
		if end > len(matches) {
			end = len(matches)
		}

		c.JSON(http.StatusOK, gin.H{
			"filtered_by_department": filtered,
			"page":                   page,
			"limit":                  limit,
			"total_matches":          len(matches),
			"employees":              matches[start:end],
		})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
