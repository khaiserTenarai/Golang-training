package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"day12/model"
	"day12/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeUseCase interface {
	Create(context.Context, model.EmployeeInput) (model.Employee, error)
	GetByID(context.Context, int64) (model.Employee, error)
	GetAll(context.Context, int, int, string, string, string) ([]model.Employee, int, error)
	Update(context.Context, int64, model.EmployeeInput) (model.Employee, error)
	Delete(context.Context, int64) error
}

type EmployeeController struct {
	service EmployeeUseCase
	db      *pgxpool.Pool
}

func NewEmployeeController(s EmployeeUseCase, db *pgxpool.Pool) *EmployeeController {
	return &EmployeeController{service: s, db: db}
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	var in model.EmployeeInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON body"})
		return
	}
	e, err := c.service.Create(ctx.Request.Context(), in)
	if errors.Is(err, service.ErrInvalidInput) {
		ctx.JSON(400, gin.H{"error": "name, valid email, department and non-negative salary are required"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not create employee (email may already exist)"})
		return
	}
	ctx.JSON(201, e)
}

func (c *EmployeeController) GetAll(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(400, gin.H{"error": "page must be a positive integer"})
		return
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		ctx.JSON(400, gin.H{"error": "limit must be a positive integer"})
		return
	}
	if limit > 100 {
		limit = 100
	}
	department := ctx.Query("department")
	if len(department) > 100 {
		ctx.JSON(400, gin.H{"error": "department filter is too long"})
		return
	}
	sortBy := ctx.DefaultQuery("sort", "id")
	order := strings.ToLower(ctx.DefaultQuery("order", "asc"))
	if order != "asc" && order != "desc" {
		ctx.JSON(400, gin.H{"error": "order must be asc or desc"})
		return
	}
	if sortBy != "id" && sortBy != "name" && sortBy != "salary" && sortBy != "department" && sortBy != "created_at" {
		ctx.JSON(400, gin.H{"error": "sort must be id, name, salary, department, or created_at"})
		return
	}
	items, total, err := c.service.GetAll(ctx.Request.Context(), page, limit, department, sortBy, order)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not fetch employees"})
		return
	}
	ctx.JSON(200, gin.H{"data": items, "page": page, "limit": limit, "total": total})
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	e, err := c.service.GetByID(ctx.Request.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not fetch employee"})
		return
	}
	ctx.JSON(200, e)
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	var in model.EmployeeInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON body"})
		return
	}
	e, err := c.service.Update(ctx.Request.Context(), id, in)
	if errors.Is(err, service.ErrInvalidInput) {
		ctx.JSON(400, gin.H{"error": "name, valid email, department and non-negative salary are required"})
		return
	}
	if errors.Is(err, pgx.ErrNoRows) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not update employee (email may already exist)"})
		return
	}
	ctx.JSON(200, e)
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	id, ok := parseID(ctx)
	if !ok {
		return
	}
	err := c.service.Delete(ctx.Request.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		ctx.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not delete employee"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func parseID(ctx *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id < 1 {
		ctx.JSON(400, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func (c *EmployeeController) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "UP", "time": time.Now().UTC()})
}

func (c *EmployeeController) Ready(ctx *gin.Context) {
	checkCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()
	if err := c.db.Ping(checkCtx); err != nil {
		ctx.JSON(503, gin.H{"status": "NOT_READY", "database": "DOWN"})
		return
	}
	ctx.JSON(200, gin.H{"status": "READY", "database": "UP"})
}
