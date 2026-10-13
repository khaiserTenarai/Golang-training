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

	router.GET("/map", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "this reply is a gin.H map", "employees": len(employees)})
	})

	router.GET("/struct", func(c *gin.Context) {
		if len(employees) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "there are no employees yet"})
			return
		}
		c.JSON(http.StatusOK, employees[0])
	})

	router.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	router.GET("/pretty", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, employees)
	})

	router.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this is what an error reply looks like"})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
