package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

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

type EmployeeForm struct {
	Name       string  `form:"name" binding:"required"`
	Email      string  `form:"email" binding:"required,email"`
	Department string  `form:"department"`
	Salary     float64 `form:"salary" binding:"required"`
}

var (
	employees = []Employee{}
	nextID    = 1
	mu        sync.Mutex
)

func save(e Employee) Employee {
	mu.Lock()
	defer mu.Unlock()
	e.ID = nextID
	nextID++
	employees = append(employees, e)
	return e
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	port := askPort(reader)

	router := gin.Default()

	router.GET("/employee", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		c.JSON(http.StatusOK, employees)
	})

	router.POST("/employee/manual", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		department := c.DefaultPostForm("department", "General")
		salary, err := strconv.ParseFloat(c.PostForm("salary"), 64)

		if name == "" || email == "" || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name, email and a numeric salary are required"})
			return
		}
		c.JSON(http.StatusCreated, save(Employee{Name: name, Email: email, Department: department, Salary: salary}))
	})

	router.POST("/employee/bind", func(c *gin.Context) {
		var form EmployeeForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.Department == "" {
			form.Department = "General"
		}
		c.JSON(http.StatusCreated, save(Employee{Name: form.Name, Email: form.Email, Department: form.Department, Salary: form.Salary}))
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
