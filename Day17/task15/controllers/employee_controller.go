package controllers

import (
	"net/http"
	"strconv"

	"employee-platform/models"
	"employee-platform/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	svc service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{svc: svc}
}

func (ctrl *EmployeeController) GetEmployees(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	sortBy := c.DefaultQuery("sortBy", "name")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	var filter models.EmployeeFilter
	if dept := c.Query("departmentId"); dept != "" {
		filter.DepartmentID = &dept
	}

	emps, total, err := ctrl.svc.ListEmployees(filter, models.Pagination{Limit: limit, Offset: offset}, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   emps,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.svc.CreateEmployee(&emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, emp)
}
