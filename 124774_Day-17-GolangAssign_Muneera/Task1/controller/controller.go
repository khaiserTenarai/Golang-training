package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-management/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: svc}
}

func (c *EmployeeController) GetAll(ctx *gin.Context) {
	employees, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve employees",
		})
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee ID",
		})
		return
	}

	employee, err := c.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve employee",
		})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}
