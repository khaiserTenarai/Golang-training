package handler

import (
	"day19/department-service/dto"
	"day19/department-service/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// DepartmentHandler translates HTTP requests into service calls (handler layer).
type DepartmentHandler struct{ service *service.DepartmentService }

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: s}
}
func (h *DepartmentHandler) Register(r *gin.Engine) {
	g := r.Group("/departments")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.Get)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
func (h *DepartmentHandler) Create(c *gin.Context) {
	var in dto.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}
	out, err := h.service.Create(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not create department", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, out)
}
func (h *DepartmentHandler) List(c *gin.Context) {
	out, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list departments"})
		return
	}
	c.JSON(http.StatusOK, out)
}
func parseID(c *gin.Context) (uint, error) {
	n, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || n == 0 {
		return 0, errors.New("id must be a positive integer")
	}
	return uint(n), nil
}
func (h *DepartmentHandler) Get(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	out, err := h.service.Get(c.Request.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "department not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "could not get department"})
		return
	}
	c.JSON(200, out)
}
func (h *DepartmentHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var in dto.UpdateDepartmentRequest
	if err = c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}
	out, err := h.service.Update(c.Request.Context(), id, in)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "department not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "could not update department"})
		return
	}
	c.JSON(200, out)
}
func (h *DepartmentHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.service.Delete(c.Request.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "department not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "could not delete department"})
		return
	}
	c.Status(http.StatusNoContent)
}
