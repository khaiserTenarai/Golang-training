package handler

import (
	"day19/employee-service/dto"
	"day19/employee-service/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type EmployeeHandler struct{ service *service.EmployeeService }

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}
func (h *EmployeeHandler) Register(r *gin.Engine) {
	g := r.Group("/employees")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:id", h.Get)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}
func cid(c *gin.Context) string { v, _ := c.Get("correlation_id"); s, _ := v.(string); return s }
func (h *EmployeeHandler) Create(c *gin.Context) {
	var in dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}
	out, err := h.service.Create(c.Request.Context(), in, cid(c))
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, out)
}
func (h *EmployeeHandler) List(c *gin.Context) {
	out, err := h.service.List(c.Request.Context(), cid(c))
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(200, out)
}
func parseID(c *gin.Context) (uint, error) {
	n, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || n == 0 {
		return 0, errors.New("id must be a positive integer")
	}
	return uint(n), nil
}
func (h *EmployeeHandler) Get(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	out, err := h.service.Get(c.Request.Context(), id, cid(c))
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(200, out)
}
func (h *EmployeeHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var in dto.UpdateEmployeeRequest
	if err = c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}
	out, err := h.service.Update(c.Request.Context(), id, in, cid(c))
	if err != nil {
		writeServiceError(c, err)
		return
	}
	c.JSON(200, out)
}
func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err = h.service.Delete(c.Request.Context(), id); err != nil {
		writeServiceError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
func writeServiceError(c *gin.Context, err error) {
	// Task 9: convert internal/downstream errors into a predictable HTTP response.
	if errors.Is(err, gorm.ErrRecordNotFound) || strings.Contains(err.Error(), "not found") {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if strings.Contains(err.Error(), "circuit breaker is open") {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "department service temporarily unavailable", "details": err.Error()})
		return
	}
	if strings.Contains(err.Error(), "department service") || strings.Contains(err.Error(), "Department Service") || strings.Contains(err.Error(), "department details") || strings.Contains(err.Error(), "validate department") || strings.Contains(err.Error(), "load department") {
		c.JSON(http.StatusBadGateway, gin.H{"error": "department service request failed", "details": err.Error()})
		return
	}
	if strings.Contains(strings.ToLower(err.Error()), "duplicate key") || strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
		c.JSON(http.StatusConflict, gin.H{"error": "employee email already exists"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
}
