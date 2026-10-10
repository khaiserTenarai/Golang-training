package controller

import (
	"net/http"
	"strconv"
	"strings"

	"student-management/model"

	"github.com/gin-gonic/gin"
)

// Students stores student records in memory.
var Students = []model.Student{
	{ID: 1, Name: "Ganesh", Age: 21, Email: "ganesh@example.com"},
	{ID: 2, Name: "Ravi", Age: 22, Email: "ravi@example.com"},
}

// GetStudents returns students and supports query parameters.
func GetStudents(ctx *gin.Context) {
	name := ctx.Query("name")
	ageText := ctx.Query("age")

	var result []model.Student

	for _, student := range Students {
		if name != "" && !strings.Contains(
			strings.ToLower(student.Name),
			strings.ToLower(name),
		) {
			continue
		}

		if ageText != "" {
			age, err := strconv.Atoi(ageText)
			if err != nil || age <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid age query parameter",
				})
				return
			}

			if student.Age != age {
				continue
			}
		}

		result = append(result, student)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Students retrieved successfully",
		"data":    result,
	})
}

// GetStudent returns a student by ID.
func GetStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	for _, student := range Students {
		if student.ID == id {
			ctx.JSON(http.StatusOK, gin.H{"data": student})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// CreateStudent creates a student using JSON.
func CreateStudent(ctx *gin.Context) {
	var student model.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if student.Name == "" || student.Age <= 0 || student.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, valid age, and email are required",
		})
		return
	}

	student.ID = nextStudentID()
	Students = append(Students, student)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"data":    student,
	})
}

// CreateStudentForm creates a student using form parameters.
func CreateStudentForm(ctx *gin.Context) {
	name := strings.TrimSpace(ctx.PostForm("name"))
	ageText := ctx.PostForm("age")
	email := strings.TrimSpace(ctx.PostForm("email"))

	age, err := strconv.Atoi(ageText)
	if err != nil || age <= 0 || name == "" || email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Valid name, age, and email are required",
		})
		return
	}

	student := model.Student{
		ID:    nextStudentID(),
		Name:  name,
		Age:   age,
		Email: email,
	}

	Students = append(Students, student)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully using form data",
		"data":    student,
	})
}

// UpdateStudent updates a student by ID.
func UpdateStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var updated model.Student
	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
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
			updated.ID = id
			Students[i] = updated

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Student updated successfully",
				"data":    updated,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// DeleteStudent deletes a student by ID.
func DeleteStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	for i, student := range Students {
		if student.ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Student deleted successfully",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// nextStudentID returns the next available student ID.
func nextStudentID() int {
	id := 1
	for _, student := range Students {
		if student.ID >= id {
			id = student.ID + 1
		}
	}
	return id
}
