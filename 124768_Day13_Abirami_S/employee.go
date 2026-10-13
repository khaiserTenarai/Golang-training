package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Employee struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EmployeeInput struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// EmployeeService keeps request context flowing into database operations.
type EmployeeService struct {
	db *pgxpool.Pool
}

func newEmployeeService(db *pgxpool.Pool) *EmployeeService {
	return &EmployeeService{db: db}
}

func (s *EmployeeService) create(ctx context.Context, in EmployeeInput) (Employee, error) {
	in.Name, in.Email, in.Department = strings.TrimSpace(in.Name), strings.TrimSpace(in.Email), strings.TrimSpace(in.Department)
	if in.Name == "" || len(in.Name) > 100 || !strings.Contains(in.Email, "@") || len(in.Email) > 150 ||
		in.Department == "" || len(in.Department) > 100 || in.Salary < 0 {
		return Employee{}, errors.New("invalid employee fields")
	}
	var e Employee
	err := s.db.QueryRow(ctx, `INSERT INTO employees(name,email,department,salary)
        VALUES($1,$2,$3,$4) RETURNING id,name,email,department,salary,created_at,updated_at`,
		in.Name, in.Email, in.Department, in.Salary).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (s *EmployeeService) getByID(ctx context.Context, id int64) (Employee, error) {
	var e Employee
	err := s.db.QueryRow(ctx, `SELECT id,name,email,department,salary,created_at,updated_at FROM employees WHERE id=$1`, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (s *EmployeeService) getAll(ctx context.Context, page, limit int, department, sortBy, order string) ([]Employee, int, error) {
	allowed := map[string]string{"id": "id", "name": "name", "salary": "salary", "department": "department", "created_at": "created_at"}
	column := allowed[sortBy]
	if column == "" {
		column = "id"
	}
	direction := "ASC"
	if strings.EqualFold(order, "desc") {
		direction = "DESC"
	}

	var total int
	if err := s.db.QueryRow(ctx, `SELECT COUNT(*) FROM employees WHERE ($1='' OR department ILIKE $2)`,
		department, "%"+department+"%").Scan(&total); err != nil {
		return nil, 0, err
	}
	query := `SELECT id,name,email,department,salary,created_at,updated_at FROM employees
        WHERE ($1='' OR department ILIKE $2) ORDER BY ` + column + ` ` + direction + ` LIMIT $3 OFFSET $4`
	rows, err := s.db.Query(ctx, query, department, "%"+department+"%", limit, (page-1)*limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	result := make([]Employee, 0)
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, 0, err
		}
		result = append(result, e)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (s *EmployeeService) update(ctx context.Context, id int64, in EmployeeInput) (Employee, error) {
	in.Name, in.Email, in.Department = strings.TrimSpace(in.Name), strings.TrimSpace(in.Email), strings.TrimSpace(in.Department)
	if in.Name == "" || len(in.Name) > 100 || !strings.Contains(in.Email, "@") || len(in.Email) > 150 ||
		in.Department == "" || len(in.Department) > 100 || in.Salary < 0 {
		return Employee{}, errors.New("invalid employee fields")
	}
	var e Employee
	err := s.db.QueryRow(ctx, `UPDATE employees SET name=$1,email=$2,department=$3,salary=$4,updated_at=NOW()
        WHERE id=$5 RETURNING id,name,email,department,salary,created_at,updated_at`,
		in.Name, in.Email, in.Department, in.Salary, id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (s *EmployeeService) delete(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

type API struct {
	employees *EmployeeService
	db        *pgxpool.Pool
	secrets   Secrets
	logger    interface {
		Info(string, ...any)
	}
	exportEnabled bool
}

func newAPI(svc *EmployeeService, db *pgxpool.Pool, secrets Secrets, logger interface{ Info(string, ...any) }, exportEnabled bool) *API {
	return &API{employees: svc, db: db, secrets: secrets, logger: logger, exportEnabled: exportEnabled}
}

func (a *API) routes() *gin.Engine {
	r := gin.New()
	r.Use(recoveryMiddleware())
	r.Use(func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = strconv.FormatInt(time.Now().UnixNano(), 10)
		}
		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)
		c.Next()
	})
	r.Use(loggingMiddleware())

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "UP"}) })
	r.GET("/ready", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := a.db.Ping(ctx); err != nil {
			c.JSON(503, gin.H{"status": "NOT_READY"})
			return
		}
		c.JSON(200, gin.H{"status": "READY", "database": "UP"})
	})
	r.GET("/features", func(c *gin.Context) { c.JSON(200, gin.H{"employee_export_enabled": a.exportEnabled}) })

	group := r.Group("/employees")
	group.Use(a.authenticate())
	group.GET("", a.getAll)
	group.GET("/:id", a.getByID)
	group.POST("", a.requireAdmin(), a.create)
	group.PUT("/:id", a.requireAdmin(), a.update)
	group.DELETE("/:id", a.requireAdmin(), a.delete)
	return r
}

func (a *API) authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		parts := strings.SplitN(strings.TrimSpace(c.GetHeader("Authorization")), " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			c.Abort()
			return
		}
		switch parts[1] {
		case a.secrets.AdminToken:
			c.Set("role", "admin")
		case a.secrets.UserToken:
			c.Set("role", "user")
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
func (a *API) requireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(403, gin.H{"error": "admin role required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
func (a *API) create(c *gin.Context) {
	var in EmployeeInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	e, err := a.employees.create(c.Request.Context(), in)
	if err != nil {
		if err.Error() == "invalid employee fields" {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": "could not create employee; email may already exist"})
		return
	}
	c.JSON(201, e)
}
func (a *API) getAll(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(400, gin.H{"error": "page must be positive"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(400, gin.H{"error": "limit must be positive"})
		return
	}
	if limit > 100 {
		limit = 100
	}
	department := c.Query("department")
	sortBy, order := c.DefaultQuery("sort", "id"), strings.ToLower(c.DefaultQuery("order", "asc"))
	if order != "asc" && order != "desc" {
		c.JSON(400, gin.H{"error": "order must be asc or desc"})
		return
	}
	if len(department) > 100 {
		c.JSON(400, gin.H{"error": "department filter too long"})
		return
	}
	if sortBy != "id" && sortBy != "name" && sortBy != "salary" && sortBy != "department" && sortBy != "created_at" {
		c.JSON(400, gin.H{"error": "invalid sort field"})
		return
	}
	items, total, err := a.employees.getAll(c.Request.Context(), page, limit, department, sortBy, order)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not fetch employees"})
		return
	}
	c.JSON(200, gin.H{"data": items, "page": page, "limit": limit, "total": total})
}
func (a *API) getByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	e, err := a.employees.getByID(c.Request.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "could not fetch employee"})
		return
	}
	c.JSON(200, e)
}
func (a *API) update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var in EmployeeInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	e, err := a.employees.update(c.Request.Context(), id, in)
	if err != nil {
		if err.Error() == "invalid employee fields" {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(404, gin.H{"error": "employee not found"})
			return
		}
		c.JSON(500, gin.H{"error": "could not update employee; email may already exist"})
		return
	}
	c.JSON(200, e)
}
func (a *API) delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	err := a.employees.delete(c.Request.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "could not delete employee"})
		return
	}
	c.Status(http.StatusNoContent)
}
func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id < 1 {
		c.JSON(400, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}
