package controller

import (
	"net/http"

	"employee-api/internal/model"
	"employee-api/internal/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	svc service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{svc: svc}
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	var emp model.Employee
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.svc.CreateEmployee(ctx.Request.Context(), &emp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, emp)
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	emp, err := c.svc.GetEmployee(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, emp)
}
