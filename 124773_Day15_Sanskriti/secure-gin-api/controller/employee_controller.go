package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"secure-gin-api/model"
	"secure-gin-api/repository"
	"secure-gin-api/service"
	"secure-gin-api/utility"
	"strconv"
	"strings"
)

type EmployeeController struct {
	Service          *service.EmployeeService
	JWTSecret        string
	JWTExpiryMinutes int
}

func NewEmployeeController(s *service.EmployeeService, secret string, expiry int) *EmployeeController {
	return &EmployeeController{Service: s, JWTSecret: secret, JWTExpiryMinutes: expiry}
}
func (c *EmployeeController) Health(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "UP"}) }
func (c *EmployeeController) Login(ctx *gin.Context) {
	var in model.LoginInput
	if err := ctx.ShouldBindJSON(&in); err != nil || strings.TrimSpace(in.Username) == "" || in.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}
	// Demo credentials are hardcoded as requested. Replace with a database-backed
	// user lookup and password hashes before deploying this API.
	var userID int64
	var role string
	switch {
	case in.Username == "admin" && in.Password == "admin123":
		userID = 1
		role = "admin"
	case in.Username == "user" && in.Password == "user123":
		userID = 2
		role = "user"
	default:
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := utility.GenerateToken(c.JWTSecret, c.JWTExpiryMinutes, userID, in.Username, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "token_type": "Bearer", "expires_in_minutes": c.JWTExpiryMinutes, "username": in.Username, "role": role})
}
func (c *EmployeeController) Create(ctx *gin.Context) {
	var in model.EmployeeInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON request"})
		return
	}
	e, err := c.Service.Create(ctx.Request.Context(), in)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, e)
}
func (c *EmployeeController) Get(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	e, err := c.Service.Get(ctx.Request.Context(), id)
	if errors.Is(err, repository.ErrNotFound) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "database error"})
		return
	}
	ctx.JSON(200, e)
}
func (c *EmployeeController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if page < 1 || limit < 1 || limit > 100 {
		ctx.JSON(400, gin.H{"error": "page must be >= 1 and limit must be between 1 and 100"})
		return
	}
	sort := ctx.DefaultQuery("sort", "id")
	order := ctx.DefaultQuery("order", "asc")
	if order != "asc" && order != "desc" {
		ctx.JSON(400, gin.H{"error": "order must be asc or desc"})
		return
	}
	items, total, err := c.Service.List(ctx.Request.Context(), page, limit, sort, order)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "database error"})
		return
	}
	pages := (total + limit - 1) / limit
	ctx.JSON(200, gin.H{"data": items, "pagination": gin.H{"page": page, "limit": limit, "total": total, "total_pages": pages}})
}
func (c *EmployeeController) Update(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	var in model.EmployeeInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON request"})
		return
	}
	e, err := c.Service.Update(ctx.Request.Context(), id, in)
	if errors.Is(err, repository.ErrNotFound) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, e)
}
func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	err := c.Service.Delete(ctx.Request.Context(), id)
	if errors.Is(err, repository.ErrNotFound) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "database error"})
		return
	}
	ctx.Status(204)
}
func parseID(ctx *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id < 1 {
		ctx.JSON(400, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}
