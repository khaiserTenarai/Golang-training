package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-api/model"
	"employee-api/service"
	"employee-api/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EmployeeController defines employee HTTP handlers.
type EmployeeController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
}

// employeeController implements EmployeeController.
type employeeController struct {
	service service.EmployeeService
}

// NewEmployeeController creates an employee controller.
func NewEmployeeController(
	svc service.EmployeeService,
) EmployeeController {
	return &employeeController{service: svc}
}

// Create handles POST /employees.
func (ctrl *employeeController) Create(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON request body",
		})
		return
	}

	err := ctrl.service.Create(c.Request.Context(), &employee)
	if err != nil {
		switch {
		case errors.Is(err, validation.ErrNameRequired),
			errors.Is(err, validation.ErrEmailRequired),
			errors.Is(err, validation.ErrEmailInvalid):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not create employee",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// GetByID handles GET /employees/:id.
func (ctrl *employeeController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "employee ID must be a positive number",
		})
		return
	}

	employee, err := ctrl.service.GetByID(
		c.Request.Context(),
		uint(id),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "employee not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not retrieve employee",
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}
