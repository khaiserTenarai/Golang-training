package controller

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service       service.EmployeeService
	jwtSecret     string
	expiryMinutes int
}

func NewEmployeeController(
	s service.EmployeeService,
	secret string,
	expiryMinutes int,
) *EmployeeController {
	return &EmployeeController{
		service:       s,
		jwtSecret:     secret,
		expiryMinutes: expiryMinutes,
	}
}

// RegisterRoutes registers all application endpoints.
func (c *EmployeeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/healthz", c.Health)
	r.POST("/login", c.Login)

	// Demonstration routes for middleware questions.
	r.GET("/demo/middleware", c.DemoMiddleware)
	r.GET("/demo/abort", middleware.AbortDemo(), c.AbortTarget)

	// Authenticated routes.
	r.GET(
		"/employees",
		middleware.Auth(c.jwtSecret),
		c.GetAll,
	)

	r.GET(
		"/employees/:id",
		middleware.Auth(c.jwtSecret),
		c.GetByID,
	)

	// Only ADMIN can create, update, delete, or upload.
	r.POST(
		"/employees",
		middleware.Auth(c.jwtSecret),
		middleware.RequireRole("ADMIN"),
		c.Create,
	)

	r.PUT(
		"/employees/:id",
		middleware.Auth(c.jwtSecret),
		middleware.RequireRole("ADMIN"),
		c.Update,
	)

	r.DELETE(
		"/employees/:id",
		middleware.Auth(c.jwtSecret),
		middleware.RequireRole("ADMIN"),
		c.Delete,
	)

	r.POST(
		"/upload",
		middleware.Auth(c.jwtSecret),
		middleware.RequireRole("ADMIN"),
		c.Upload,
	)
}

func (c *EmployeeController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (c *EmployeeController) DemoMiddleware(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "middleware executed successfully",
	})
}

func (c *EmployeeController) AbortTarget(ctx *gin.Context) {
	// This should not execute because AbortDemo stops the chain.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "this handler should not execute",
	})
}

// Login is a demonstration login.
// Replace these demo credentials with real user storage in production.
func (c *EmployeeController) Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required",
		})
		return
	}

	role := ""

	switch {
	case input.Username == "admin" && input.Password == "admin123":
		role = "ADMIN"

	case input.Username == "user" && input.Password == "user123":
		role = "USER"

	default:
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	token, err := middleware.GenerateToken(
		input.Username,
		role,
		c.jwtSecret,
		c.expiryMinutes,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  role,
	})
}

// Question 11: Request validation.
func (c *EmployeeController) Create(ctx *gin.Context) {
	var e model.Employee

	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee data: name, valid email, department and non-negative salary are required",
		})
		return
	}

	if err := c.service.Create(ctx.Request.Context(), &e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, e)
}

// Question 12: Pagination and filtering.
func (c *EmployeeController) GetAll(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "page must be a positive number",
		})
		return
	}

	size, err := strconv.Atoi(ctx.DefaultQuery("size", "20"))
	if err != nil || size < 1 || size > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "size must be between 1 and 100",
		})
		return
	}

	department := ctx.Query("department")
	sortBy := ctx.DefaultQuery("sort", "id")

	employees, total, err := c.service.GetAll(
		ctx.Request.Context(),
		page,
		size,
		department,
		sortBy,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not retrieve employees",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page":       page,
		"size":       size,
		"total":      total,
		"department": department,
		"data":       employees,
	})
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "employee id must be a positive number",
		})
		return
	}

	employee, err := c.service.GetByID(ctx.Request.Context(), id)

	if repository.IsNotFound(err) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "employee not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not retrieve employee",
		})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "employee id must be a positive number",
		})
		return
	}

	var e model.Employee

	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee data",
		})
		return
	}

	updated, err := c.service.Update(
		ctx.Request.Context(),
		id,
		&e,
	)

	if repository.IsNotFound(err) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "employee not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "employee id must be a positive number",
		})
		return
	}

	err = c.service.Delete(ctx.Request.Context(), id)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) ||
			strings.Contains(err.Error(), "employee not found") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "employee not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not delete employee",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// Question 13: File upload.
func (c *EmployeeController) Upload(ctx *gin.Context) {
	const maxSize = 5 << 20 // 5 MiB

	ctx.Request.Body = http.MaxBytesReader(
		ctx.Writer,
		ctx.Request.Body,
		maxSize,
	)

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "please upload a file using the file field",
		})
		return
	}
	defer file.Close()

	if header.Size > maxSize {
		ctx.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"error": "file size must not exceed 5 MB",
		})
		return
	}

	if err := os.MkdirAll("uploads", 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create upload directory",
		})
		return
	}

	// Use only the extension, not the supplied path or filename.
	extension := filepath.Ext(filepath.Base(header.Filename))
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + extension
	destination := filepath.Join("uploads", filename)

	if err := ctx.SaveUploadedFile(header, destination); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not save uploaded file",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "file uploaded successfully",
		"filename": filename,
	})
}
