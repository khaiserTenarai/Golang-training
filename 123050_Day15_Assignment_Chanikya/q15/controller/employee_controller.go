package controller

import (
	"errors"
	"net/http"
	"strconv"

	"ems/model"
	"ems/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

// GET /employees?page=1&pageSize=10
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid page"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid pageSize"})
		return
	}

	employees, total, err := c.service.GetAll(page, pageSize)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	ctx.JSON(200, gin.H{
		"data": employees,
		"pagination": gin.H{
			"page":       page,
			"pageSize":   pageSize,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}

	e, err := c.service.GetByID(id)
	if err != nil {
		writeError(ctx, err)
		return
	}

	ctx.JSON(200, e)
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	var e model.Employee

	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid employee JSON"})
		return
	}

	if err := c.service.Create(&e); err != nil {
		writeError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, e)
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}

	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid employee JSON"})
		return
	}

	updated, err := c.service.Update(id, &e)
	if err != nil {
		writeError(ctx, err)
		return
	}

	ctx.JSON(200, updated)
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}

	if err := c.service.Delete(id); err != nil {
		writeError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func parseID(ctx *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(400, gin.H{"error": "ID must be a positive integer"})
		return 0, false
	}
	return uint(id), true
}

func writeError(ctx *gin.Context, err error) {
	switch {
	case service.IsNotFound(err):
		ctx.JSON(404, gin.H{"error": "employee not found"})
	case errors.Is(err, gorm.ErrDuplicatedKey):
		ctx.JSON(409, gin.H{"error": "email already exists"})
	case err.Error() == "employee not found":
		ctx.JSON(404, gin.H{"error": "employee not found"})
	case err.Error() == "invalid employee ID":
		ctx.JSON(400, gin.H{"error": err.Error()})
	case err.Error() == "name is required and must be at most 100 characters",
		err.Error() == "a valid email is required",
		err.Error() == "department is required and must be at most 100 characters",
		err.Error() == "salary cannot be negative",
		err.Error() == "employee is required":
		ctx.JSON(400, gin.H{"error": err.Error()})
	default:
		ctx.JSON(500, gin.H{"error": "internal server error"})
	}
}
