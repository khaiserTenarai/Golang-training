package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Role     string  `json:"role" binding:"required"`
	Department string `json:"department" binding:"required"`
	Salary   float64 `json:"salary" binding:"required,gt=0"`
}

type UpdateEmployee struct {
	Name       *string  `json:"name,omitempty"`
	Role       *string  `json:"role,omitempty"`
	Department *string  `json:"department,omitempty"`
	Salary     *float64 `json:"salary,omitempty"`
}

type EmployeeStore struct {
	mu        sync.RWMutex
	employees map[int]Employee
	nextID    int
}

func NewEmployeeStore() *EmployeeStore {
	store := &EmployeeStore{
		employees: make(map[int]Employee),
		nextID:    1,
	}
	// Initial seed data
	store.employees[1] = Employee{ID: 1, Name: "Alice Smith", Role: "Backend Engineer", Department: "Engineering", Salary: 95000}
	store.nextID = 2
	return store
}

func main() {
	r := gin.Default()
	store := NewEmployeeStore()

	r.HandleMethodNotAllowed = true

	// Custom 404 / 405
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "HTTP method not allowed on this endpoint"})
	})

	// API Group
	api := r.Group("/api/v1")
	{
		api.GET("/employees", func(c *gin.Context) {
			dept := c.Query("department")
			store.mu.RLock()
			defer store.mu.RUnlock()

			list := []Employee{}
			for _, emp := range store.employees {
				if dept == "" || emp.Department == dept {
					list = append(list, emp)
				}
			}
			c.JSON(http.StatusOK, list)
		})

		api.GET("/employees/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
				return
			}

			store.mu.RLock()
			emp, exists := store.employees[id]
			store.mu.RUnlock()

			if !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				return
			}
			c.JSON(http.StatusOK, emp)
		})

		api.POST("/employees", func(c *gin.Context) {
			var newEmp Employee
			if err := c.ShouldBindJSON(&newEmp); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			store.mu.Lock()
			newEmp.ID = store.nextID
			store.nextID++
			store.employees[newEmp.ID] = newEmp
			store.mu.Unlock()

			c.JSON(http.StatusCreated, newEmp)
		})

		api.PUT("/employees/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
				return
			}

			var updatedEmp Employee
			if err := c.ShouldBindJSON(&updatedEmp); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			if _, exists := store.employees[id]; !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				return
			}

			updatedEmp.ID = id
			store.employees[id] = updatedEmp
			c.JSON(http.StatusOK, updatedEmp)
		})

		api.PATCH("/employees/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
				return
			}

			var req UpdateEmployee
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			emp, exists := store.employees[id]
			if !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				return
			}

			if req.Name != nil {
				emp.Name = *req.Name
			}
			if req.Role != nil {
				emp.Role = *req.Role
			}
			if req.Department != nil {
				emp.Department = *req.Department
			}
			if req.Salary != nil {
				emp.Salary = *req.Salary
			}

			store.employees[id] = emp
			c.JSON(http.StatusOK, emp)
		})

		api.DELETE("/employees/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
				return
			}

			store.mu.Lock()
			defer store.mu.Unlock()

			if _, exists := store.employees[id]; !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				return
			}

			delete(store.employees, id)
			c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully", "id": id})
		})
	}

	r.Run(":8080")
}
