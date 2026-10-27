package handler

import (
	"net/http"
	"strconv"

	"employee-platform/middleware"
	"employee-platform/service"
	"github.com/gin-gonic/gin"
)

type API struct {
	employees   *service.EmployeeService
	departments *service.DepartmentService
	auth        *service.AuthService
}

func RegisterRoutes(r *gin.Engine, employees *service.EmployeeService, departments *service.DepartmentService, auth *service.AuthService, gqlHTTP http.Handler, jwtSecret string) {
	api := &API{employees: employees, departments: departments, auth: auth}
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.POST("/api/auth/register", api.register)
	r.POST("/api/auth/login", api.login)
	// Public department list makes it possible to discover valid department IDs for initial data setup.
	r.GET("/api/departments", api.listDepartments)
	r.POST("/api/departments", middleware.Auth(jwtSecret), middleware.RequireRole("admin"), api.createDepartment)
	secured := r.Group("/api")
	secured.Use(middleware.Auth(jwtSecret))
	{
		secured.GET("/employees", api.listEmployees)
		secured.GET("/employees/:id", api.getEmployee)
		secured.POST("/employees", middleware.RequireRole("admin"), api.createEmployee)
		secured.PUT("/employees/:id", middleware.RequireRole("admin"), api.updateEmployee)
		secured.DELETE("/employees/:id", middleware.RequireRole("admin"), api.deleteEmployee)
	}
	r.Any("/graphql", gin.WrapH(gqlHTTP))
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *API) register(c *gin.Context) {
	var req registerRequest
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	user, err := a.auth.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"id": user.ID, "name": user.Name, "email": user.Email, "role": user.Role})
}
func (a *API) login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	user, token, err := a.auth.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}
	c.JSON(200, gin.H{"token": token, "user": gin.H{"id": user.ID, "name": user.Name, "email": user.Email, "role": user.Role}})
}
func (a *API) listDepartments(c *gin.Context) {
	items, err := a.departments.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, items)
}
func (a *API) createDepartment(c *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	d, err := a.departments.Create(req.Name, req.Description)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, d)
}
func (a *API) listEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	dept, _ := strconv.ParseUint(c.Query("department_id"), 10, 64)
	result, err := a.employees.List(page, limit, uint(dept), c.Query("search"), c.DefaultQuery("sort_by", "id"), c.DefaultQuery("order", "asc"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}
func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(id), err == nil && id > 0
}
func (a *API) getEmployee(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	e, err := a.employees.ByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(200, e)
}
func (a *API) createEmployee(c *gin.Context) {
	var req struct {
		Name         string  `json:"name"`
		Email        string  `json:"email"`
		Salary       float64 `json:"salary"`
		DepartmentID uint    `json:"department_id"`
	}
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	e, err := a.employees.Create(req.Name, req.Email, req.Salary, req.DepartmentID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, e)
}
func (a *API) updateEmployee(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var req struct {
		Name         string  `json:"name"`
		Email        string  `json:"email"`
		Salary       float64 `json:"salary"`
		DepartmentID uint    `json:"department_id"`
	}
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	e, err := a.employees.Update(id, req.Name, req.Email, req.Salary, req.DepartmentID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, e)
}
func (a *API) deleteEmployee(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := a.employees.Delete(id); err != nil {
		c.JSON(404, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(200, gin.H{"message": "employee deleted"})
}
