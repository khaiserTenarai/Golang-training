package handler

import (
	"employee-microservices/department/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DepartmentHandler is the HTTP layer. Task 5.
type DepartmentHandler struct{ service *service.DepartmentService }

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: s}
}
func (h *DepartmentHandler) RegisterRoutes(r *gin.Engine) { r.GET("/departments", h.List) }
func (h *DepartmentHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": h.service.List()})
}
