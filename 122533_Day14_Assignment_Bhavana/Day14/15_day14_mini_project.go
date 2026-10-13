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

// what a caller has to send to create or fully replace an employee
type EmployeeInput struct {
	Name       string  `json:"name" binding:"required,min=2"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"required,gt=0"`
}

// for PATCH: pointers, so "not sent" can be told apart from "sent as empty"
type EmployeePatch struct {
	Name       *string  `json:"name" binding:"omitempty,min=2"`
	Email      *string  `json:"email" binding:"omitempty,email"`
	Department *string  `json:"department"`
	Salary     *float64 `json:"salary" binding:"omitempty,gt=0"`
}

type Store struct {
	mu        sync.Mutex
	employees []Employee
	nextID    int
}

func (s *Store) emailTaken(email string, exceptID int) bool {
	for _, e := range s.employees {
		if e.ID != exceptID && strings.EqualFold(e.Email, email) {
			return true
		}
	}
	return false
}

func (s *Store) indexOf(id int) int {
	for i, e := range s.employees {
		if e.ID == id {
			return i
		}
	}
	return -1
}

func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return 0, false
	}
	return id, true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	initial := askEmployees(reader)
	port := askPort(reader)

	store := &Store{employees: initial, nextID: len(initial) + 1}

	router := gin.Default()
	router.HandleMethodNotAllowed = true

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found", "path": c.Request.URL.Path})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed", "method": c.Request.Method})
	})

	api := router.Group("/api/v1")
	{
		// ---------- list ----------
		api.GET("/employee", func(c *gin.Context) {
			department := c.Query("department")
			search := strings.ToLower(c.Query("search"))
			page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
			if err != nil || page < 1 {
				page = 1
			}
			limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
			if err != nil || limit < 1 {
				limit = 10
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			matches := []Employee{}
			for _, e := range store.employees {
				if department != "" && !strings.EqualFold(e.Department, department) {
					continue
				}
				if search != "" && !strings.Contains(strings.ToLower(e.Name), search) {
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
				"data":  matches[start:end],
				"page":  page,
				"limit": limit,
				"total": len(matches),
			})
		})

		// ---------- read one ----------
		api.GET("/employee/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			store.mu.Lock()
			defer store.mu.Unlock()

			idx := store.indexOf(id)
			if idx == -1 {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.JSON(http.StatusOK, store.employees[idx])
		})

		// ---------- create ----------
		api.POST("/employee", func(c *gin.Context) {
			var input EmployeeInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			if store.emailTaken(input.Email, 0) {
				c.JSON(http.StatusConflict, gin.H{"error": "that email is already in use"})
				return
			}
			e := Employee{
				ID:         store.nextID,
				Name:       input.Name,
				Email:      input.Email,
				Department: input.Department,
				Salary:     input.Salary,
			}
			store.nextID++
			store.employees = append(store.employees, e)
			c.JSON(http.StatusCreated, e)
		})

		// ---------- replace ----------
		api.PUT("/employee/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			var input EmployeeInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			idx := store.indexOf(id)
			if idx == -1 {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			if store.emailTaken(input.Email, id) {
				c.JSON(http.StatusConflict, gin.H{"error": "that email is already in use"})
				return
			}
			store.employees[idx] = Employee{
				ID:         id,
				Name:       input.Name,
				Email:      input.Email,
				Department: input.Department,
				Salary:     input.Salary,
			}
			c.JSON(http.StatusOK, store.employees[idx])
		})

		// ---------- partial update ----------
		api.PATCH("/employee/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			var patch EmployeePatch
			if err := c.ShouldBindJSON(&patch); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			idx := store.indexOf(id)
			if idx == -1 {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			if patch.Email != nil && store.emailTaken(*patch.Email, id) {
				c.JSON(http.StatusConflict, gin.H{"error": "that email is already in use"})
				return
			}

			e := &store.employees[idx]
			if patch.Name != nil {
				e.Name = *patch.Name
			}
			if patch.Email != nil {
				e.Email = *patch.Email
			}
			if patch.Department != nil {
				e.Department = *patch.Department
			}
			if patch.Salary != nil {
				e.Salary = *patch.Salary
			}
			c.JSON(http.StatusOK, *e)
		})

		// ---------- delete ----------
		api.DELETE("/employee/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			store.mu.Lock()
			defer store.mu.Unlock()

			idx := store.indexOf(id)
			if idx == -1 {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			store.employees = append(store.employees[:idx], store.employees[idx+1:]...)
			c.Status(http.StatusNoContent)
		})
	}

	fmt.Println("Employee API running on port", port)
	if err := router.Run(":" + port); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
