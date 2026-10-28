package controller

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"employee-management/auth"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service     service.EmployeeService
	authService *auth.AuthService
}

func NewEmployeeController(
	s service.EmployeeService,
	a *auth.AuthService,
) *EmployeeController {
	return &EmployeeController{
		service:     s,
		authService: a,
	}
}

func (ctl *EmployeeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/healthz", ctl.Health)
	r.POST("/login", ctl.Login)

	// Demonstration route for middleware exercises.
	r.GET("/demo/middleware", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "middleware executed",
			"requestID": c.GetString("requestID"),
		})
	})

	// Demonstration route for Abort.
	r.GET("/demo/abort", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "request stopped by Abort middleware",
		})
	})

	// All employee endpoints require authentication.
	employees := r.Group("/employees")
	employees.Use(middleware.RequireAuth(ctl.authService))

	employees.GET("", ctl.GetAll)
	employees.GET("/:id", ctl.GetByID)

	// Only admins can create, update, or delete employees.
	admin := employees.Group("")
	admin.Use(middleware.RequireRole("admin"))

	admin.POST("", ctl.Create)
	admin.PUT("/:id", ctl.Update)
	admin.DELETE("/:id", ctl.Delete)
}

// Health checks the database connection.
func (ctl *EmployeeController) Health(c *gin.Context) {
	if err := ctl.service.Health(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unavailable",
			"error":  "database unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Login authenticates a user and returns a JWT.
func (ctl *EmployeeController) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required",
		})
		return
	}

	token, err := ctl.authService.Login(
		request.Username,
		request.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

// GetAll returns a paginated list of employees.
func (ctl *EmployeeController) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page must be a positive integer",
		})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page_size must be between 1 and 100",
		})
		return
	}

	sort := strings.ToLower(c.DefaultQuery("sort", "id"))
	sortOrder := strings.ToLower(c.DefaultQuery("sort_order", "asc"))

	allowedSortFields := map[string]bool{
		"id":         true,
		"name":       true,
		"department": true,
		"salary":     true,
	}

	if !allowedSortFields[sort] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid sort field",
		})
		return
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "sort_order must be asc or desc",
		})
		return
	}

	result, err := ctl.service.GetAll(
		c.Request.Context(),
		c.Query("department"),
		c.Query("name"),
		sort,
		sortOrder,
		page,
		pageSize,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve employees",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetByID returns one employee by ID.
func (ctl *EmployeeController) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee ID",
		})
		return
	}

	employee, err := ctl.service.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "employee not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve employee",
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// Create adds a new employee.
func (ctl *EmployeeController) Create(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee data",
		})
		return
	}

	if err := ctl.service.Create(
		c.Request.Context(),
		&employee,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// Update modifies an existing employee.
func (ctl *EmployeeController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee ID",
		})
		return
	}

	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee data",
		})
		return
	}

	updatedEmployee, err := ctl.service.Update(
		c.Request.Context(),
		id,
		&employee,
	)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "employee not found",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

// Delete removes an employee by ID.
func (ctl *EmployeeController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee ID",
		})
		return
	}

	if err := ctl.service.Delete(
		c.Request.Context(),
		id,
	); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "employee not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete employee",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
