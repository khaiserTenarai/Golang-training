package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-management/model"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

// GET /employees
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	employees, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve employees"})
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

// GET /employees/:id
func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}

	employee, err := c.service.GetByID(id)
	if err != nil {
		if service.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve employee"})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// POST /employees
func (c *EmployeeController) Create(ctx *gin.Context) {
	var employee model.Employee

	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON request"})
		return
	}

	// The client cannot choose the database-generated ID.
	employee.ID = 0

	if err := c.service.Create(&employee); err != nil {
		if isDuplicateKey(err) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}

		if isValidationError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create employee"})
		return
	}

	ctx.JSON(http.StatusCreated, employee)
}

// PUT /employees/:id
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}

	var employee model.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON request"})
		return
	}

	employee.ID = id

	updated, err := c.service.Update(id, employee)
	if err != nil {
		if service.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}

		if isDuplicateKey(err) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}

		if isValidationError(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update employee"})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

// DELETE /employees/:id
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}

	err := c.service.Delete(id)
	if err != nil {
		if service.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete employee"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func parseID(ctx *gin.Context) (uint, bool) {
	value, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil || value == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a positive integer"})
		return 0, false
	}

	return uint(value), true
}

func isDuplicateKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}

	// PostgreSQL unique-constraint violations commonly use SQLSTATE 23505.
	type sqlStater interface {
		SQLState() string
	}

	var stateErr sqlStater
	return errors.As(err, &stateErr) && stateErr.SQLState() == "23505"
}

func isValidationError(err error) bool {
	// Service validation errors are returned as ordinary errors.
	// Database errors should not be treated as client validation errors.
	switch err.Error() {
	case "employee cannot be nil",
		"name is required",
		"a valid email is required",
		"department is required",
		"salary cannot be negative",
		"employee ID must be positive":
		return true
	default:
		return false
	}
}
