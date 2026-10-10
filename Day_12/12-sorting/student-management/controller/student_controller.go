package controller

import (
	"net/http"
	"strconv"
	"student-management/model"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

type StudentController struct{ service service.StudentService }

func NewStudentController(s service.StudentService) *StudentController {
	return &StudentController{service: s}
}

func (c *StudentController) CreateStudent(ctx *gin.Context) {
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "send valid JSON with name and grade"})
		return
	}
	if err := c.service.CreateStudent(ctx.Request.Context(), student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "student created"})
}

func (c *StudentController) GetStudents(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page must be a positive number"})
		return
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "limit must be between 1 and 100"})
		return
	}
	sortBy := ctx.DefaultQuery("sortBy", "id")
	order := ctx.DefaultQuery("order", "asc")
	students, total, err := c.service.GetStudents(ctx.Request.Context(), ctx.Query("name"), ctx.Query("grade"), sortBy, order, page, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": students, "total": total, "page": page, "limit": limit,
		"sort": gin.H{"sortBy": sortBy, "order": order},
	})
}
