package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"employee-api/model"
	"employee-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EmployeeController defines the employee HTTP handlers.
type EmployeeController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
}

// employeeController implements EmployeeController.
type employeeController struct {
	service service.EmployeeService
}

// NewEmployeeController creates a controller with its service dependency.
func NewEmployeeController(svc service.EmployeeService) EmployeeController {
	return &employeeController{service: svc}
}

// Create handles POST /employees.
func (h *employeeController) Create(c *gin.Context) {
	var employee model.Employee

	// Decode the incoming JSON request.
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee JSON",
		})
		return
	}

	// Pass the request context to the service.
	err := h.service.Create(c.Request.Context(), &employee)
	if err != nil {
		switch {
		case service.IsValidationError(err):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		case errors.Is(err, context.Canceled),
			errors.Is(err, context.DeadlineExceeded):
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error": "request cancelled or timed out",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create employee",
			})
		}
		return
	}

	// Return the saved employee.
	c.JSON(http.StatusCreated, employee)
}

// GetByID handles GET /employees/:id.
func (h *employeeController) GetByID(c *gin.Context) {
	// Convert the URL parameter to a positive integer.
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "employee ID must be a positive integer",
		})
		return
	}

	// Retrieve the employee using the request context.
	employee, err := h.service.GetByID(
		c.Request.Context(),
		uint(id64),
	)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": "employee not found",
			})

		case errors.Is(err, context.Canceled),
			errors.Is(err, context.DeadlineExceeded):
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error": "request cancelled or timed out",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to retrieve employee",
			})
		}
		return
	}

	c.JSON(http.StatusOK, employee)
}