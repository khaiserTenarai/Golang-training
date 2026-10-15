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

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	total, err := strconv.Atoi(ask("How many sample employees should be created? "))
	if err != nil || total < 1 {
		fmt.Println("Please enter a whole number of 1 or more.")
		return
	}

	var employees []Employee
	for i := 1; i <= total; i++ {
		employees = append(employees, Employee{ID: i, Name: fmt.Sprintf("Employee %d", i)})
	}

	router := gin.New()

	router.GET("/employees", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}
		if limit > 100 {
			limit = 100
		}

		start := (page - 1) * limit
		if start > len(employees) {
			start = len(employees)
		}
		end := start + limit
		if end > len(employees) {
			end = len(employees)
		}

		totalPages := (len(employees) + limit - 1) / limit

		c.JSON(http.StatusOK, gin.H{
			"page":        page,
			"limit":       limit,
			"total":       len(employees),
			"total_pages": totalPages,
			"data":        employees[start:end],
		})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
