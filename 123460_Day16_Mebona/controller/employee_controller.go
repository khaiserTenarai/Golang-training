package controller

import (
	"net/http"

	"employee-management/model"
	"employee-management/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct{ service service.EmployeeService }

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}

func (c *EmployeeController) GetAll(ctx *gin.Context) {
	list, err := c.service.List(model.EmployeeFilter{Limit: 100})
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	e, err := c.service.Get(id)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, e)
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	var in model.CreateEmployeeInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}
	e, err := c.service.Create(in)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, e)
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	var in model.UpdateEmployeeInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}
	e, err := c.service.Update(id, in)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, e)
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	if err := c.service.Delete(id); err != nil {
		respondError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
