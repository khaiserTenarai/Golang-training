// Day 17 — Advanced GraphQL (All 15 Tasks)
// Task 14: REST + GraphQL on same server
// Task 15: Mini Project — REST + GraphQL Employee platform with security and testing
//
// Run:  go mod tidy && go run main.go
// Test: go test -v ./tests/...
//
// GraphQL: POST http://localhost:8080/graphql
// REST:    GET/POST/PUT/DELETE http://localhost:8080/api/employees

package main

import (
	"context"
	//"encoding/json"
	"fmt"
	//"net/http"
	"strconv"
	"strings"

	"graphql_employee/models"
	"graphql_employee/schema"
	"graphql_employee/service"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var gqlSchema graphql.Schema
var empSvc *service.EmployeeService
var deptSvc *service.DepartmentService
var authSvc *service.AuthService

func main() {
	empSvc = service.NewEmployeeService()
	deptSvc = service.NewDepartmentService()
	authSvc = service.NewAuthService()

	var err error
	gqlSchema, err = schema.BuildSchema(empSvc, deptSvc, authSvc)
	if err != nil {
		panic("Failed to build schema: " + err.Error())
	}

	r := gin.Default()

	// Public
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "healthy"}) })

	// Auth
	r.POST("/auth/login", loginREST)
	r.POST("/auth/register", registerREST)

	// GraphQL endpoint
	r.POST("/graphql", graphqlHandler)
	r.GET("/graphql", playgroundHandler)

	// Task 14: REST endpoints alongside GraphQL
	api := r.Group("/api")
	api.Use(jwtMiddleware())
	{
		api.GET("/employees", restGetAll)
		api.GET("/employees/:id", restGetByID)
		api.POST("/employees", requireRoleMiddleware("admin", "manager"), restCreate)
		api.PUT("/employees/:id", requireRoleMiddleware("admin", "manager"), restUpdate)
		api.DELETE("/employees/:id", requireRoleMiddleware("admin"), restDelete)
	}

	fmt.Println("*********************************************************")
	fmt.Println("   Day 17: Advanced GraphQL + REST Employee Platform  ")
	fmt.Println("**********************************************************")
	fmt.Println("->  GraphQL:  POST /graphql                             ")
	fmt.Println("->  Playground: GET /graphql (browser)                  ")
	fmt.Println("->  REST:     /api/employees (CRUD)                     ")
	fmt.Println("->  Auth:     POST /auth/login, /auth/register          ")
	fmt.Println("->  Users:    admin/admin123, piyush/piyush123          ")
	fmt.Println("->  Tests:    go test -v ./tests/...                    ")
	fmt.Println("**********************************************************")

	r.Run(":8080")
}

// ---- GraphQL Handler ----

func graphqlHandler(c *gin.Context) {
	var req struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Extract JWT from header and add user info to context
	ctx := context.Background()
	if token := extractToken(c); token != "" {
		if uid, uname, role, err := service.ParseJWT(token); err == nil {
			ctx = context.WithValue(ctx, "user_id", uid)
			ctx = context.WithValue(ctx, "username", uname)
			ctx = context.WithValue(ctx, "role", role)
		}
	}

	result := graphql.Do(graphql.Params{
		Schema:         gqlSchema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
		Context:        ctx,
	})
	c.JSON(200, result)
}

func playgroundHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(200, `<!DOCTYPE html><html><head><title>GraphQL Playground</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css"/>
<script src="https://cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
</head><body><div id="root"></div><script>
window.addEventListener('load',function(){GraphQLPlayground.init(document.getElementById('root'),{endpoint:'/graphql'})})
</script></body></html>`)
}

// ---- REST Handlers (Task 14) ----

func loginREST(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"success": false, "message": "username and password required"})
		return
	}
	token, user, err := authSvc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "token": token, "user": gin.H{"id": user.ID, "username": user.Username, "role": user.Role}})
}

func registerREST(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role"`
	}
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"success": false, "message": "invalid input"})
		return
	}
	token, user, err := authSvc.Register(req.Username, req.Password, req.Role)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"success": true, "token": token, "user": gin.H{"id": user.ID, "username": user.Username, "role": user.Role}})
}

func restGetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	filter := models.EmployeeFilter{
		Name:      c.Query("name"),
		SortBy:    c.DefaultQuery("sort_by", "id"),
		SortOrder: c.DefaultQuery("sort_order", "asc"),
	}
	if v, err := strconv.ParseFloat(c.Query("min_salary"), 64); err == nil {
		filter.MinSalary = v
	}
	list, total := empSvc.GetAll(filter, models.PaginationInput{Page: page, PageSize: pageSize})
	c.JSON(200, gin.H{"success": true, "data": list, "total": total, "page": page, "page_size": pageSize})
}

func restGetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	emp, err := empSvc.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": emp})
}

func restCreate(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	created, err := empSvc.Create(emp)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(201, gin.H{"success": true, "data": created})
}

func restUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var emp models.Employee
	if c.ShouldBindJSON(&emp) != nil {
		c.JSON(400, gin.H{"success": false, "message": "invalid JSON"})
		return
	}
	updated, err := empSvc.Update(id, emp)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": updated})
}

func restDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := empSvc.Delete(id); err != nil {
		c.JSON(404, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "message": "deleted"})
}

// ---- Middleware ----

func extractToken(c *gin.Context) string {
	return strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
}

func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" || token == c.GetHeader("Authorization") {
			c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Bearer token required"})
			return
		}
		uid, uname, role, err := service.ParseJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Invalid token"})
			return
		}
		c.Set("user_id", uid)
		c.Set("username", uname)
		c.Set("role", role)
		c.Next()
	}
}

func requireRoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		for _, r := range roles {
			if r == role {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(403, gin.H{"success": false, "message": "access denied"})
	}
}

/*
════════════════════════════════════════════
  CURL COMMANDS
════════════════════════════════════════════

# Login
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"admin123\"}"

# GraphQL — Login
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -d "{\"query\":\"mutation { login(username: \\\"admin\\\", password: \\\"admin123\\\") { token username role } }\"}"

# GraphQL — List employees (use token from login)
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -H "Authorization: Bearer TOKEN" -d "{\"query\":\"{ employees(page:1, page_size:3, sort_by:\\\"salary\\\", sort_order:\\\"desc\\\") { employees { id name salary department { name location } } total total_pages } }\"}"

# GraphQL — Get single employee
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -H "Authorization: Bearer TOKEN" -d "{\"query\":\"{ employee(id:1) { id name email salary department { name } } }\"}"

# GraphQL — Create employee (admin only)
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -H "Authorization: Bearer TOKEN" -d "{\"query\":\"mutation { createEmployee(input: {name: \\\"Frank\\\", email: \\\"frank@test.com\\\", age: 30, department_id: 1, salary: 70000}) { id name } }\"}"

# REST — List (with token)
curl -H "Authorization: Bearer TOKEN" http://localhost:8080/api/employees

# REST — Create (admin token)
curl -X POST http://localhost:8080/api/employees -H "Authorization: Bearer TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"Grace\",\"email\":\"grace@test.com\",\"age\":26,\"department_id\":2,\"salary\":55000}"

# Playground: open http://localhost:8080/graphql in browser
*/
