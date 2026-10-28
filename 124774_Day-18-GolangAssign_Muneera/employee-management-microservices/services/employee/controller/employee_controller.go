package controller

import (
	"errors"
	"net/http"
	"strconv"

	"employee-management-microservices/services/employee/model"
	"employee-management-microservices/services/employee/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller handles employee HTTP requests.
type Controller struct {
	service *service.Service
}

func New(s *service.Service) *Controller {
	return &Controller{service: s}
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return 0, false
	}
	return uint(id), true
}

func (h *Controller) Create(c *gin.Context) {
	var e model.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.Create(c.Request.Context(), &e); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, e)
}

func (h *Controller) GetAll(c *gin.Context) {
	employees, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employees"})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *Controller) GetByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	e, err := h.service.GetByID(id)
	if err != nil {
		if service.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employee"})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (h *Controller) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var e model.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	e.ID = id

	if err := h.service.Update(c.Request.Context(), &e); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, e)
}

func (h *Controller) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	if err := h.service.Delete(id); err != nil {
		if service.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}

func (h *Controller) writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrDepartmentNotFound):
		c.JSON(http.StatusBadRequest, gin.H{"error": "Department does not exist"})
	case errors.Is(err, gorm.ErrDuplicatedKey):
		c.JSON(http.StatusConflict, gin.H{"error": "Employee email already exists"})
	default:
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "Could not complete employee operation",
		})
	}
}
