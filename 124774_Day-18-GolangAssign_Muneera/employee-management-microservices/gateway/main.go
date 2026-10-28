package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"employee-management-microservices/discovery"
	"employee-management-microservices/gateway/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file.
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	// Validate JWT secret.
	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) < 32 {
		log.Fatalf(
			"JWT_SECRET must contain at least 32 characters; got %d",
			len(jwtSecret),
		)
	}

	// Configure Consul address.
	consulAddress := strings.TrimSpace(os.Getenv("CONSUL_HTTP_ADDR"))
	if consulAddress == "" {
		consulAddress = "http://127.0.0.1:8500"
		_ = os.Setenv("CONSUL_HTTP_ADDR", consulAddress)
	}

	// Create the Gateway router.
	router := gin.Default()

	// Configure authentication middleware.
	authMiddleware := middleware.Auth(jwtSecret)
	adminMiddleware := middleware.RequireRole("ADMIN")

	// Gateway health endpoint.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Authentication Service.
	// Login is public.
	authProxy := discovery.Proxy(
		"auth-service",
		"/api",
	)
	router.POST("/api/auth/login", authProxy)

	// Employee Service.
	employeeProxy := discovery.Proxy(
		"employee-service",
		"/api",
	)
	registerServiceRoutes(
		router,
		"/api/employees",
		authMiddleware,
		adminMiddleware,
		employeeProxy,
	)

	// Department Service.
	departmentProxy := discovery.Proxy(
		"department-service",
		"/api",
	)
	registerServiceRoutes(
		router,
		"/api/departments",
		authMiddleware,
		adminMiddleware,
		departmentProxy,
	)

	// Start API Gateway.
	port := strings.TrimSpace(os.Getenv("GATEWAY_PORT"))
	if port == "" {
		port = "8080"
	}

	log.Printf("API Gateway listening on port %s", port)
	log.Printf("Consul address: %s", consulAddress)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

// registerServiceRoutes configures service routes.
// GET and HEAD require authentication.
// POST, PUT, PATCH and DELETE require ADMIN authorization.
func registerServiceRoutes(
	router *gin.Engine,
	basePath string,
	authMiddleware gin.HandlerFunc,
	adminMiddleware gin.HandlerFunc,
	proxy gin.HandlerFunc,
) {
	// Read operations.
	router.GET(basePath, authMiddleware, proxy)
	router.GET(basePath+"/*path", authMiddleware, proxy)

	router.HEAD(basePath, authMiddleware, proxy)
	router.HEAD(basePath+"/*path", authMiddleware, proxy)

	// Create operations.
	router.POST(basePath, authMiddleware, adminMiddleware, proxy)
	router.POST(basePath+"/*path", authMiddleware, adminMiddleware, proxy)

	// Update operations.
	router.PUT(basePath, authMiddleware, adminMiddleware, proxy)
	router.PUT(basePath+"/*path", authMiddleware, adminMiddleware, proxy)

	router.PATCH(basePath, authMiddleware, adminMiddleware, proxy)
	router.PATCH(basePath+"/*path", authMiddleware, adminMiddleware, proxy)

	// Delete operations.
	router.DELETE(basePath, authMiddleware, adminMiddleware, proxy)
	router.DELETE(basePath+"/*path", authMiddleware, adminMiddleware, proxy)
}
