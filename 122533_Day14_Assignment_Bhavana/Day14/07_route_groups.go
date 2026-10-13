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

func adminOnly(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Admin-Token") != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "admin token missing or wrong"})
			return
		}
		c.Next()
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	token := ""
	for token == "" {
		token = ask(reader, "Choose an admin token for the /admin routes: ")
	}
	employees := askEmployees(reader)
	port := askPort(reader)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/employee", func(c *gin.Context) {
			c.JSON(http.StatusOK, employees)
		})
		v1.GET("/employee/:id", func(c *gin.Context) {
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
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/employee", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"count": len(employees), "employees": employees})
		})
	}

	admin := router.Group("/admin", adminOnly(token))
	{
		admin.GET("/stats", func(c *gin.Context) {
			var total float64
			for _, e := range employees {
				total += e.Salary
			}
			c.JSON(http.StatusOK, gin.H{"employees": len(employees), "total_salary": total})
		})
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
