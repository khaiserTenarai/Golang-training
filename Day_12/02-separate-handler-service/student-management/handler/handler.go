package handler

import (
	"net/http"
	"student-management/model"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

// StudentHandler handles HTTP requests.
type StudentHandler struct {
	service *service.StudentService
}

// NewStudentHandler creates the handler.
func NewStudentHandler(s *service.StudentService) *StudentHandler {
	return &StudentHandler{service: s}
}

// GetAll handles GET requests.
func (h *StudentHandler) GetAll(c *gin.Context) {
	students := h.service.GetAll()

	c.JSON(http.StatusOK, students)
}

// Create handles POST requests.
func (h *StudentHandler) Create(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student data",
		})
		return
	}

	createdStudent := h.service.Create(student)

	c.JSON(http.StatusCreated, createdStudent)
}
