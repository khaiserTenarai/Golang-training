package handler

import (
	"net/http"
	"student-management/model"
	"student-management/service"
	"student-management/utility"

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

// GetAll returns all students.
func (h *StudentHandler) GetAll(c *gin.Context) {
	students := h.service.GetAll()
	c.JSON(http.StatusOK, students)
}

// Create validates and adds a student.
func (h *StudentHandler) Create(c *gin.Context) {
	var student model.Student

	// Read the JSON request.
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
		return
	}

	// Validate student details.
	if err := utility.ValidateStudent(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save the valid student.
	createdStudent := h.service.Create(student)

	c.JSON(http.StatusCreated, createdStudent)
}
