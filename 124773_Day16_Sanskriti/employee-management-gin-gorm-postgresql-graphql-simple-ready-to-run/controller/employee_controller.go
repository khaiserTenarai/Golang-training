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
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	e, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
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
		ctx.JSON(400, gin.H{"error": err.Error()})
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
	if service.IsNotFound(err) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
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
	if service.IsNotFound(err) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
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
		if service.IsNotFound(err) {
			ctx.JSON(404, gin.H{"error": "employee not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(204)
}
