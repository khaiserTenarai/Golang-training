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

// GetStudents returns all students or filters by query parameters.
func GetStudents(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")

	var result []model.Student

	for _, student := range Students {
		if name != "" && !strings.Contains(
			strings.ToLower(student.Name),
			strings.ToLower(name),
		) {
			continue
		}

		if age != "" {
			studentAge, err := strconv.Atoi(age)
			if err != nil || studentAge <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid age query parameter",
				})
				return
			}

			if student.Age != studentAge {
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}

	for _, student := range Students {
		if student.ID == id {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Student retrieved successfully",
				"data":    student,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Student not found",
	})
}

// CreateStudent binds JSON and creates a student.
func CreateStudent(ctx *gin.Context) {
	var student model.Student

	// Convert the JSON request body into a Student struct.
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
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

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"data":    student,
	})
}

// UpdateStudent updates a student by ID.
func UpdateStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}

	var updated model.Student
	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
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

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Student not found",
	})
}

// DeleteStudent deletes a student by ID.
func DeleteStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
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

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Student not found",
	})
}
