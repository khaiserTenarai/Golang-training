package api

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gin-secure-api-day15/internal/config"
	"gin-secure-api-day15/internal/middleware"
	"gin-secure-api-day15/internal/model"
	"gin-secure-api-day15/internal/store"
)

type Router struct {
	cfg  config.Config
	repo store.EmployeeStore
}

func NewRouter(cfg config.Config, repo store.EmployeeStore) *gin.Engine {
	r := &Router{cfg: cfg, repo: repo}
	engine := gin.New()
	engine.Use(middleware.SafeRecovery(), middleware.RequestID(), middleware.TraceOrder(), middleware.Timing())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"X-Request-ID", "X-Response-Time"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	engine.POST("/api/auth/token", r.issueToken)

	protected := engine.Group("/api")
	protected.Use(middleware.Authenticate(cfg.JWTSecret))
	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get(middleware.UserIDKey)
		role, _ := c.Get(middleware.RoleKey)
		c.JSON(http.StatusOK, gin.H{"user_id": userID, "role": role})
	})
	protected.GET("/employees", r.listEmployees)
	protected.GET("/employees/:id", r.getEmployee)
	protected.POST("/employees", middleware.RequireRoles("admin", "manager"), r.createEmployee)
	protected.PUT("/employees/:id", middleware.RequireRoles("admin", "manager"), r.updateEmployee)
	protected.PATCH("/employees/:id", middleware.RequireRoles("admin", "manager"), r.updateEmployee)
	protected.DELETE("/employees/:id", middleware.RequireRoles("admin"), r.deleteEmployee)
	protected.GET("/admin/overview", middleware.RequireRoles("admin"), func(c *gin.Context) {
		_, total := r.repo.List("", "", 0, 1)
		c.JSON(http.StatusOK, gin.H{"employee_count": total, "message": "admin access granted"})
	})

	protected.POST("/files", middleware.RequireRoles("admin", "manager"), func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file field is required"})
			return
		}
		if file.Size > 5<<20 {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file must be 5 MB or smaller"})
			return
		}
		if file.Size == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty files are not accepted"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"filename": filepath.Base(file.Filename), "size": file.Size})
	})

	engine.GET("/demo/next-abort", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "request reached endpoint"})
	})
	engine.GET("/demo/panic", func(c *gin.Context) {
		panic("demo panic")
	})
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
	})
	engine.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
	})
	return engine
}

func (r *Router) issueToken(c *gin.Context) {
	var request struct {
		UserID int    `json:"user_id" binding:"required,gt=0"`
		Role   string `json:"role" binding:"required,oneof=admin manager employee"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	token, err := middleware.CreateToken(r.cfg.JWTSecret, request.UserID, request.Role, time.Duration(r.cfg.TokenTTLHours)*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": token, "token_type": "Bearer", "expires_in": r.cfg.TokenTTLHours * 3600})
}

func (r *Router) listEmployees(c *gin.Context) {
	page, err := positiveQueryInt(c, "page", 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page must be a positive integer"})
		return
	}
	limit, err := positiveQueryInt(c, "limit", 10)
	if err != nil || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be between 1 and 100"})
		return
	}
	items, total := r.repo.List(c.Query("search"), c.Query("department"), (page-1)*limit, limit)
	c.JSON(http.StatusOK, gin.H{
		"data": items,
		"pagination": gin.H{
			"page": page, "limit": limit, "total": total,
			"total_pages": (total + limit - 1) / limit,
		},
	})
}

func (r *Router) getEmployee(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	employee, exists := r.repo.Get(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func (r *Router) createEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}
	created, err := r.repo.Create(employee)
	if err == store.ErrConflict {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": created})
}

func (r *Router) updateEmployee(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}
	updated, err := r.repo.Update(id, employee)
	if err == store.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	if err == store.ErrConflict {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updated})
}

func (r *Router) deleteEmployee(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if !r.repo.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func positiveQueryInt(c *gin.Context, key string, fallback int) (int, error) {
	value := c.Query(key)
	if value == "" {
		return fallback, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return 0, strconv.ErrSyntax
	}
	return parsed, nil
}
