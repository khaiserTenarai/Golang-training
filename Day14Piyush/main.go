// Day 14 — Gin Framework (All 15 Tasks)
// Task 1  — Create a Gin application
// Task 7  — Create route groups
// Task 13 — Serve static files
// Task 14 — Custom 404/405 handling
// Task 15 — Mini Project: Gin Employee CRUD REST API
//
// Run: go mod tidy && go run main.go
// Open: http://localhost:8080 (static page)
// API:  http://localhost:8080/api/employees

package main

import (
	"fmt"
	"net/http"

	"gin_employee_api/handler"
	"gin_employee_api/middleware"
	"gin_employee_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Task 1: Create a Gin application
	r := gin.Default()

	// Add custom middleware
	r.Use(middleware.RequestLogger())

	// Task 13: Serve static files
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")

	// Task 14: Custom 404/405 handling
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Message: "Route not found: " + c.Request.Method + " " + c.Request.URL.Path,
		})
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Message: "Method " + c.Request.Method + " not allowed for " + c.Request.URL.Path,
		})
	})

	// Health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "framework": "gin"})
	})

	// Initialize handler
	empHandler := handler.NewEmployeeHandler()

	// Task 7: Route groups
	api := r.Group("/api")
	{
		employees := api.Group("/employees")
		{
			employees.GET("", empHandler.GetAll)             // Task 2: GET
			employees.GET("/:id", empHandler.GetByID)        // Task 2, 8: GET + path param
			employees.POST("", empHandler.Create)            // Task 3, 10: POST + JSON bind
			employees.POST("/form", empHandler.CreateFromForm) // Task 12: Form params
			employees.PUT("/:id", empHandler.Update)         // Task 4: PUT
			employees.PATCH("/:id", empHandler.Patch)        // Task 5: PATCH
			employees.DELETE("/:id", empHandler.Delete)      // Task 6: DELETE
		}
	}

	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║   Day 14: Gin Employee CRUD REST API        ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║  Static:  http://localhost:8080              ║")
	fmt.Println("║  Health:  GET  /health                      ║")
	fmt.Println("║  List:    GET  /api/employees               ║")
	fmt.Println("║  Detail:  GET  /api/employees/:id           ║")
	fmt.Println("║  Create:  POST /api/employees (JSON)        ║")
	fmt.Println("║  Form:    POST /api/employees/form          ║")
	fmt.Println("║  Update:  PUT  /api/employees/:id           ║")
	fmt.Println("║  Patch:   PATCH /api/employees/:id          ║")
	fmt.Println("║  Delete:  DELETE /api/employees/:id         ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	r.Run(":8080")
}

/*
════════════════════════════════════════
  CURL COMMANDS
════════════════════════════════════════

# Task 13: Static page
curl http://localhost:8080

# Task 2: GET all (Task 9: query params, Task 11: JSON response)
curl http://localhost:8080/api/employees
curl "http://localhost:8080/api/employees?name=piyush"
curl "http://localhost:8080/api/employees?department=Engineering&sort_by=salary&sort_order=desc"
curl "http://localhost:8080/api/employees?page=1&page_size=2"

# Task 2, 8: GET by ID (path parameter)
curl http://localhost:8080/api/employees/1

# Task 3, 10: POST with JSON binding
curl -X POST http://localhost:8080/api/employees -H "Content-Type: application/json" -d "{\"name\":\"Eve\",\"email\":\"eve@test.com\",\"age\":26,\"department\":\"HR\",\"salary\":52000}"

# Task 12: POST with form data
curl -X POST http://localhost:8080/api/employees/form -d "name=Frank&email=frank@test.com&age=29&department=Finance&salary=58000"

# Task 4: PUT (full update)
curl -X PUT http://localhost:8080/api/employees/1 -H "Content-Type: application/json" -d "{\"name\":\"Piyush Updated\",\"email\":\"p@new.com\",\"age\":26,\"department\":\"Backend\",\"salary\":75000}"

# Task 5: PATCH (partial update — only salary)
curl -X PATCH http://localhost:8080/api/employees/1 -H "Content-Type: application/json" -d "{\"salary\":80000}"

# Task 6: DELETE
curl -X DELETE http://localhost:8080/api/employees/3

# Task 14: Custom 404
curl http://localhost:8080/nonexistent

# Validation error (binding fails)
curl -X POST http://localhost:8080/api/employees -H "Content-Type: application/json" -d "{\"name\":\"A\"}"
*/
