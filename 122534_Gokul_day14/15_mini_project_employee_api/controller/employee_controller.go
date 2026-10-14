package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"employeegin/dao"
	"employeegin/model"
	"employeegin/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func respondError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, dao.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrDuplicateEmail):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}
}

func readID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive number"})
		return 0, false
	}
	return id, true
}

func (ctl *EmployeeController) Create(c *gin.Context) {
	var req model.EmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := ctl.service.AddEmployee(req)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, employee)
}

func (ctl *EmployeeController) GetAll(c *gin.Context) {
	employees, err := ctl.service.GetAllEmployees(c.Query("department"))
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (ctl *EmployeeController) GetOne(c *gin.Context) {
	id, ok := readID(c)
	if !ok {
		return
	}

	employee, err := ctl.service.GetEmployee(id)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (ctl *EmployeeController) Replace(c *gin.Context) {
	id, ok := readID(c)
	if !ok {
		return
	}

	var req model.EmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := ctl.service.ReplaceEmployee(id, req)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (ctl *EmployeeController) Patch(c *gin.Context) {
	id, ok := readID(c)
	if !ok {
		return
	}

	var patch model.EmployeePatch
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := ctl.service.PatchEmployee(id, patch)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (ctl *EmployeeController) Delete(c *gin.Context) {
	id, ok := readID(c)
	if !ok {
		return
	}

	if err := ctl.service.DeleteEmployee(id); err != nil {
		respondError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
