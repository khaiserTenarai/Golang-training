package controller

import (
	"net/http"
	"strconv"

	"student-management/model"

	"github.com/gin-gonic/gin"
)

// Students stores student records in memory.
var Students = []model.Student{
	{ID: 1, Name: "Ganesh", Age: 21, Email: "ganesh@example.com"},
	{ID: 2, Name: "Ravi", Age: 22, Email: "ravi@example.com"},
}

// GetStudents returns all students.
func GetStudents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Students)
}

// GetStudent returns one student using a path parameter.
func GetStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	for _, student := range Students {
		if student.ID == id {
			ctx.JSON(http.StatusOK, student)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// CreateStudent adds a new student.
func CreateStudent(ctx *gin.Context) {
	var student model.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if student.Name == "" || student.Age <= 0 || student.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, valid age, and email are required",
		})
		return
	}

	// Generate the next student ID.
	student.ID = 1
	for _, existing := range Students {
		if existing.ID >= student.ID {
			student.ID = existing.ID + 1
		}
	}

	Students = append(Students, student)
	ctx.JSON(http.StatusCreated, student)
}

// UpdateStudent updates a student using a path parameter.
func UpdateStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var updated model.Student
	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if updated.Name == "" || updated.Age <= 0 || updated.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, valid age, and email are required",
		})
		return
	}

	for i, student := range Students {
		if student.ID == id {
			// Keep the ID supplied in the URL.
			updated.ID = id
			Students[i] = updated

			ctx.JSON(http.StatusOK, updated)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// DeleteStudent deletes a student using a path parameter.
func DeleteStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	for i, student := range Students {
		if student.ID == id {
			// Remove the student from the slice.
			Students = append(Students[:i], Students[i+1:]...)

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Student deleted successfully",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
