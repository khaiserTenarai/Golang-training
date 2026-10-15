package controller

import (
	"net/http"

	"employee-management/service"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct{ service service.DepartmentService }

func NewDepartmentController(s service.DepartmentService) *DepartmentController {
	return &DepartmentController{s}
}

func (c *DepartmentController) GetAll(ctx *gin.Context) {
	list, err := c.service.List()
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}
