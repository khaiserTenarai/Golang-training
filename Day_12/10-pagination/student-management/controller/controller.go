package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"student-management/model"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

// StudentController handles student HTTP requests.
type StudentController struct {
	service *service.StudentService
}

// NewStudentController creates the controller.
func NewStudentController(
	s *service.StudentService,
) *StudentController {
	return &StudentController{service: s}
}

// GetStudents lists, filters, and paginates students.
func (c *StudentController) GetStudents(ctx *gin.Context) {
	page, err := queryInt(ctx, "page", 1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "page must be a number",
		})
		return
	}

	limit, err := queryInt(ctx, "limit", 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "limit must be a number",
		})
		return
	}

	// Stop the database operation if it takes too long.
	requestCtx, cancel := context.WithTimeout(
		ctx.Request.Context(),
		5*time.Second,
	)
	defer cancel()

	students, total, err := c.service.GetStudents(
		requestCtx,
		ctx.Query("name"),
		ctx.Query("grade"),
		page,
		limit,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page":          page,
		"limit":         limit,
		"total_records": total,
		"students":      students,
	})
}

// CreateStudent creates a student from a JSON request.
func (c *StudentController) CreateStudent(ctx *gin.Context) {
	var student model.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON body",
		})
		return
	}

	requestCtx, cancel := context.WithTimeout(
		ctx.Request.Context(),
		5*time.Second,
	)
	defer cancel()

	created, err := c.service.CreateStudent(requestCtx, student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// queryInt reads an integer query parameter.
func queryInt(
	ctx *gin.Context,
	key string,
	defaultValue int,
) (int, error) {
	value := ctx.Query(key)

	if value == "" {
		return defaultValue, nil
	}

	return strconv.Atoi(value)
}
