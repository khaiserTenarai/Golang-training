package controller

import (
	"employee-management/model"
	"employee-management/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeController struct{ service service.EmployeeService }

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}

func fail(ctx *gin.Context, err error) {
	switch {
	case service.IsNotFound(err):
		ctx.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
	case service.IsValidation(err):
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case service.IsDuplicate(err):
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	e, err := c.service.GetAll(model.EmployeeFilter{})
	if err != nil {
		fail(ctx, err)
		return
	}
	ctx.JSON(200, e)
}
func (c *EmployeeController) Create(ctx *gin.Context) {
	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Create(&e); err != nil {
		fail(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, e)
}
func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid employee id"})
		return
	}
	e, err := c.service.GetByID(uint(id))
	if err != nil {
		fail(ctx, err)
		return
	}
	ctx.JSON(200, e)
}
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid employee id"})
		return
	}
	var e model.Employee
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	x, err := c.service.Update(uint(id), &e)
	if err != nil {
		fail(ctx, err)
		return
	}
	ctx.JSON(200, x)
}
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid employee id"})
		return
	}
	if err := c.service.Delete(uint(id)); err != nil {
		fail(ctx, err)
		return
	}
	ctx.Status(204)
}
