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

// GetAll returns all students as JSON.
func (h *StudentHandler) GetAll(c *gin.Context) {
	students := h.service.GetAll()
	c.JSON(http.StatusOK, students)
}

// Create adds a new student.
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
