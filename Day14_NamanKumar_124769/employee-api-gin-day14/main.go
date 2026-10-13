package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name" binding:"required,min=2,max=100"`
	Email      string    `json:"email" binding:"required,email,max=254"`
	Department string    `json:"department" binding:"required,min=2,max=80"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateEmployeeInput struct {
	Name       string `json:"name" form:"name" binding:"required,min=2,max=100"`
	Email      string `json:"email" form:"email" binding:"required,email,max=254"`
	Department string `json:"department" form:"department" binding:"required,min=2,max=80"`
}

type PatchEmployeeInput struct {
	Name       *string `json:"name" binding:"omitempty,min=2,max=100"`
	Email      *string `json:"email" binding:"omitempty,email,max=254"`
	Department *string `json:"department" binding:"omitempty,min=2,max=80"`
}

type EmployeeStore struct {
	mu        sync.RWMutex
	employees map[int64]Employee
	nextID    int64
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{employees: make(map[int64]Employee), nextID: 1}
}

func (s *EmployeeStore) Create(input CreateEmployeeInput) Employee {
	s.mu.Lock()
	defer s.mu.Unlock()

	employee := Employee{
		ID:         s.nextID,
		Name:       strings.TrimSpace(input.Name),
		Email:      strings.ToLower(strings.TrimSpace(input.Email)),
		Department: strings.TrimSpace(input.Department),
		CreatedAt:  time.Now().UTC(),
	}
	s.employees[employee.ID] = employee
	s.nextID++
	return employee
}

func (s *EmployeeStore) List(department, search string) []Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Employee, 0, len(s.employees))
	for _, employee := range s.employees {
		if department != "" && !strings.EqualFold(employee.Department, department) {
			continue
		}
		if search != "" &&
			!strings.Contains(strings.ToLower(employee.Name), strings.ToLower(search)) &&
			!strings.Contains(strings.ToLower(employee.Email), strings.ToLower(search)) {
			continue
		}
		result = append(result, employee)
	}
	return result
}

func (s *EmployeeStore) Get(id int64) (Employee, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	employee, ok := s.employees[id]
	return employee, ok
}

func (s *EmployeeStore) Update(id int64, input CreateEmployeeInput) (Employee, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	employee, ok := s.employees[id]
	if !ok {
		return Employee{}, false
	}
	employee.Name = strings.TrimSpace(input.Name)
	employee.Email = strings.ToLower(strings.TrimSpace(input.Email))
	employee.Department = strings.TrimSpace(input.Department)
	s.employees[id] = employee
	return employee, true
}

func (s *EmployeeStore) Patch(id int64, input PatchEmployeeInput) (Employee, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	employee, ok := s.employees[id]
	if !ok {
		return Employee{}, false
	}
	if input.Name != nil {
		employee.Name = strings.TrimSpace(*input.Name)
	}
	if input.Email != nil {
		employee.Email = strings.ToLower(strings.TrimSpace(*input.Email))
	}
	if input.Department != nil {
		employee.Department = strings.TrimSpace(*input.Department)
	}
	s.employees[id] = employee
	return employee, true
}

func (s *EmployeeStore) Delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.employees[id]; !ok {
		return false
	}
	delete(s.employees, id)
	return true
}

func main() {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	if parsed, err := strconv.Atoi(port); err != nil || parsed < 1 || parsed > 65535 {
		panic("PORT must be an integer between 1 and 65535")
	}

	store := NewEmployeeStore()
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.HandleMethodNotAllowed = true

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	{
		api.GET("/employees", func(c *gin.Context) {
			department := strings.TrimSpace(c.Query("department"))
			search := strings.TrimSpace(c.Query("search"))
			c.JSON(http.StatusOK, gin.H{
				"data":  store.List(department, search),
				"count": len(store.List(department, search)),
			})
		})

		api.GET("/employees/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			employee, exists := store.Get(id)
			if !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.JSON(http.StatusOK, employee)
		})

		api.POST("/employees", func(c *gin.Context) {
			var input CreateEmployeeInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee JSON", "details": err.Error()})
				return
			}
			employee := store.Create(input)
			c.JSON(http.StatusCreated, employee)
		})

		api.POST("/employees/form", func(c *gin.Context) {
			var input CreateEmployeeInput
			if err := c.ShouldBind(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee form", "details": err.Error()})
				return
			}
			employee := store.Create(input)
			c.JSON(http.StatusCreated, employee)
		})

		api.PUT("/employees/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			var input CreateEmployeeInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee JSON", "details": err.Error()})
				return
			}
			employee, exists := store.Update(id, input)
			if !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.JSON(http.StatusOK, employee)
		})

		api.PATCH("/employees/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			var input PatchEmployeeInput
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patch JSON", "details": err.Error()})
				return
			}
			if input.Name == nil && input.Email == nil && input.Department == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "at least one field is required"})
				return
			}
			employee, exists := store.Patch(id, input)
			if !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.JSON(http.StatusOK, employee)
		})

		api.DELETE("/employees/:id", func(c *gin.Context) {
			id, ok := parseID(c)
			if !ok {
				return
			}
			if !store.Delete(id) {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.Status(http.StatusNoContent)
		})
	}

	router.Static("/assets", "./static")

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
			"path":  c.Request.URL.Path,
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error":  "method not allowed",
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})
	})

	if err := router.Run(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}
