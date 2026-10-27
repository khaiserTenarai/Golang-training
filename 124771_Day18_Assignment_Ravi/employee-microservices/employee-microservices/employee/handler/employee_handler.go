package handler

import (
	"net/http"

	"employee-microservices/employee/model"
	"employee-microservices/employee/service"
	"github.com/gin-gonic/gin"
)

// EmployeeHandler is the HTTP/controller layer. Task 4.
type EmployeeHandler struct{ service *service.EmployeeService }

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/employees", h.List)
	r.POST("/employees", h.Create)
}

func (h *EmployeeHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": h.service.List()})
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var e model.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON or missing name"})
		return
	}
	created, err := h.service.Create(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": created})
}
