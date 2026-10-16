package controller

import (
	"employee-management/model"
	"employee-management/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DepartmentController struct{ service service.DepartmentService }

func NewDepartmentController(s service.DepartmentService) *DepartmentController {
	return &DepartmentController{s}
}

type departmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *DepartmentController) List(ctx *gin.Context) {
	list, err := c.service.GetAll()
	if err != nil {
		fail(ctx, err, "department")
		return
	}
	ctx.JSON(http.StatusOK, list)
}
func (c *DepartmentController) Create(ctx *gin.Context) {
	var req departmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d := model.Department{Name: req.Name, Description: req.Description}
	if err := c.service.Create(&d); err != nil {
		fail(ctx, err, "department")
		return
	}
	ctx.JSON(http.StatusCreated, d)
}
