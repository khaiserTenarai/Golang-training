// Day 15 — Advanced Gin (All 15 Tasks)
// Task 15: Mini Project — Secure Gin API with JWT, RBAC, Middleware,
//          Validation, Pagination, File Upload, Testing
//
// Run:   go mod tidy && go run main.go
// Test:  go test -v ./handler/...

package main

import (
	"fmt"
	"os"

	"advanced_gin/handler"
	"advanced_gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create uploads directory for Task 13
	os.MkdirAll("./uploads", 0755)

	r := gin.New()

	// Middleware chain (Tasks 1-7)
	r.Use(middleware.Recovery())          // Task 6: Recovery
	r.Use(middleware.RequestID())         // Task 5: Request ID
	r.Use(middleware.CORS())              // Task 7: CORS
	r.Use(middleware.TimingMiddleware())  // Task 1: Custom middleware
	r.Use(gin.Logger())

	// Initialize handlers
	empHandler := handler.NewEmployeeHandler()
	authHandler := handler.NewAuthHandler()

	// --- Public routes ---
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Task 2: Middleware execution order demo
	r.GET("/order", middleware.FirstMiddleware(), middleware.SecondMiddleware(),
		func(c *gin.Context) {
			fmt.Println("  [Order] Handler executed")
			c.JSON(200, gin.H{"message": "Check server console for middleware order"})
		})

	// Task 4: Abort demo — requires X-API-Key header
	r.GET("/api-key-test", middleware.ApiKeyMiddleware("my-secret-key"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "API key valid!", "validated": c.GetBool("api_key_validated")})
		})

	// Auth routes (Task 8: JWT)
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// --- Protected routes (Task 9: JWT required) ---
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/profile", authHandler.Profile)

		// Employee CRUD (Task 10: role-based)
		employees := api.Group("/employees")
		{
			// Any authenticated user can read
			employees.GET("", empHandler.GetAll)          // Task 12: pagination
			employees.GET("/:id", empHandler.GetByID)

			// Only admin and manager can write (Task 10: RBAC)
			write := employees.Group("")
			write.Use(middleware.RequireRole("admin", "manager"))
			{
				write.POST("", empHandler.Create)          // Task 11: validation
				write.PUT("/:id", empHandler.Update)
				write.DELETE("/:id", empHandler.Delete)
			}
		}

		// Task 13: File upload (admin only)
		upload := api.Group("/upload")
		upload.Use(middleware.RequireRole("admin"))
		{
			upload.POST("", empHandler.UploadFile)
		}
	}

	fmt.Println("╔═══════════════════════════════════════════════════╗")
	fmt.Println("║   Day 15: Advanced Gin — Secure Employee API     ║")
	fmt.Println("╠═══════════════════════════════════════════════════╣")
	fmt.Println("║  Public:                                         ║")
	fmt.Println("║    GET  /health                                  ║")
	fmt.Println("║    GET  /order          (middleware order demo)   ║")
	fmt.Println("║    GET  /api-key-test   (X-API-Key: my-secret-key)║")
	fmt.Println("║    POST /auth/register                           ║")
	fmt.Println("║    POST /auth/login                              ║")
	fmt.Println("║  Protected (Bearer JWT):                         ║")
	fmt.Println("║    GET  /api/profile                             ║")
	fmt.Println("║    GET  /api/employees      (any role)           ║")
	fmt.Println("║    GET  /api/employees/:id  (any role)           ║")
	fmt.Println("║    POST /api/employees      (admin/manager)      ║")
	fmt.Println("║    PUT  /api/employees/:id  (admin/manager)      ║")
	fmt.Println("║    DELETE /api/employees/:id (admin/manager)     ║")
	fmt.Println("║    POST /api/upload          (admin only)        ║")
	fmt.Println("╠═══════════════════════════════════════════════════╣")
	fmt.Println("║  Seeded users: admin/admin123, manager/manager123║")
	fmt.Println("║                piyush/piyush123                  ║")
	fmt.Println("║  Tests: go test -v ./handler/...                 ║")
	fmt.Println("╚═══════════════════════════════════════════════════╝")

	r.Run(":8080")
}

/*
════════════════════════════════════
  CURL COMMANDS
════════════════════════════════════

# Health
curl http://localhost:8080/health

# Task 2: Middleware order (check server console)
curl http://localhost:8080/order

# Task 4: Abort demo
curl http://localhost:8080/api-key-test
curl -H "X-API-Key: my-secret-key" http://localhost:8080/api-key-test

# Task 8: Login (get JWT)
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"admin123\"}"

# Copy the token from login response and use below:
# TOKEN=<paste_token_here>

# Task 9: Protected route
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/profile

# Task 12: Pagination + filtering
curl -H "Authorization: Bearer $TOKEN" "http://localhost:8080/api/employees?page=1&page_size=2&sort_by=salary&sort_order=desc"

# Task 10: RBAC — admin can create
curl -X POST http://localhost:8080/api/employees -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"Eve\",\"email\":\"eve@test.com\",\"age\":26,\"department\":\"HR\",\"salary\":52000}"

# Task 10: RBAC — user cannot create (login as piyush first)
# USER_TOKEN=$(curl -s -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d "{\"username\":\"piyush\",\"password\":\"piyush123\"}" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
# curl -X POST http://localhost:8080/api/employees -H "Authorization: Bearer $USER_TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"Test\"}"

# Task 13: File upload (admin only)
# curl -X POST http://localhost:8080/api/upload -H "Authorization: Bearer $TOKEN" -F "file=@./README.md"

# Task 14: Run tests
# go test -v ./handler/...
*/
