package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-management-microservices/services/employee/dto"
	"employee-management-microservices/services/employee/entity"
	"employee-management-microservices/services/employee/service"
	"employee-management-microservices/shared/httpclient"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	Service *service.EmployeeService
}

func NewEmployeeController(s *service.EmployeeService) *EmployeeController {
	return &EmployeeController{Service: s}
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	employee, err := c.Service.Create(
		ctx.Request.Context(),
		ctx.GetHeader("X-Correlation-ID"),
		req.Name,
		req.Email,
		req.Salary,
		req.DepartmentID,
	)
	if err != nil {
		if errors.Is(err, httpclient.ErrDepartmentNotFound) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "department not found"})
			return
		}

		if errors.Is(err, httpclient.ErrDepartmentUnavailable) {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "department service unavailable"})
			return
		}

		ctx.JSON(http.StatusConflict, gin.H{"error": "could not create employee"})
		return
	}

	ctx.JSON(http.StatusCreated, toEmployeeResponse(employee))
}

func (c *EmployeeController) GetAll(ctx *gin.Context) {
	employees, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch employees"})
		return
	}

	response := make([]dto.EmployeeResponse, 0, len(employees))
	for i := range employees {
		response = append(response, toEmployeeResponse(&employees[i]))
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee ID"})
		return
	}

	employee, err := c.Service.GetByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch employee"})
		return
	}

	ctx.JSON(http.StatusOK, toEmployeeResponse(employee))
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee ID"})
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	employee, err := c.Service.Update(
		ctx.Request.Context(),
		ctx.GetHeader("X-Correlation-ID"),
		uint(id),
		req.Name,
		req.Email,
		req.Salary,
		req.DepartmentID,
	)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if errors.Is(err, httpclient.ErrDepartmentNotFound) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "department not found"})
		return
	}
	if errors.Is(err, httpclient.ErrDepartmentUnavailable) {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "department service unavailable"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update employee"})
		return
	}

	ctx.JSON(http.StatusOK, toEmployeeResponse(employee))
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee ID"})
		return
	}

	err = c.Service.Delete(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete employee"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func toEmployeeResponse(employee *entity.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		ID:           employee.ID,
		Name:         employee.Name,
		Email:        employee.Email,
		Salary:       employee.Salary,
		DepartmentID: employee.DepartmentID,
	}
}
