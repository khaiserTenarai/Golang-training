package controller

import (
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: svc}
}

func (ctrl *EmployeeController) RegisterRoutes(r *gin.Engine, jwtSecret string) {
	r.POST("/login", ctrl.Login)

	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware(jwtSecret))
	{
		api.GET("/employees", ctrl.GetEmployees)

		admin := api.Group("")
		admin.Use(middleware.RBACMiddleware("ADMIN"))
		{
			admin.POST("/employees", ctrl.CreateEmployee)
			admin.POST("/employees/:id/avatar", ctrl.UploadAvatar)
			admin.DELETE("/employees/:id", ctrl.DeleteEmployee)
		}
	}
}

func (ctrl *EmployeeController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := &middleware.Claims{
		Username: req.Username,
		Role:     req.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var emp model.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateEmployee(c.Request.Context(), &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, emp)
}

func (ctrl *EmployeeController) GetEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	employees, total, err := ctrl.service.GetEmployees(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employees,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}



func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := ctrl.service.DeleteEmployee(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}


















/*func (ctrl *EmployeeController) UploadAvatar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	uploadDir := "./uploads"
	_ = os.MkdirAll(uploadDir, os.ModePerm)
	dst := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if err := ctrl.service.UpdateAvatar(c.Request.Context(), id, dst); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully", "path": dst})
}