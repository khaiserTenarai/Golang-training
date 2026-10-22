package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-management-microservices/services/department/dto"
	"employee-management-microservices/services/department/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DepartmentController struct {
	Service *service.DepartmentService
}

func NewDepartmentController(s *service.DepartmentService) *DepartmentController {
	return &DepartmentController{Service: s}
}

func (c *DepartmentController) Create(ctx *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	department, err := c.Service.Create(req.Name)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "could not create department"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.DepartmentResponse{
		ID: department.ID, Name: department.Name,
	})
}

func (c *DepartmentController) GetAll(ctx *gin.Context) {
	departments, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch departments"})
		return
	}

	response := make([]dto.DepartmentResponse, 0, len(departments))
	for _, d := range departments {
		response = append(response, dto.DepartmentResponse{ID: d.ID, Name: d.Name})
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *DepartmentController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid department ID"})
		return
	}

	department, err := c.Service.GetByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch department"})
		return
	}

	ctx.JSON(http.StatusOK, dto.DepartmentResponse{ID: department.ID, Name: department.Name})
}

func (c *DepartmentController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid department ID"})
		return
	}

	var req dto.UpdateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	department, err := c.Service.Update(uint(id), req.Name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update department"})
		return
	}

	ctx.JSON(http.StatusOK, dto.DepartmentResponse{ID: department.ID, Name: department.Name})
}

func (c *DepartmentController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid department ID"})
		return
	}

	err = c.Service.Delete(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete department"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
