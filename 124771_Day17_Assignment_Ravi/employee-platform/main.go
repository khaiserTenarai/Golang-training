package main

import (
	"log"
	"os"

	"employee-platform/config"
	"employee-platform/database"
	graphqlserver "employee-platform/graphql"
	"employee-platform/handler"
	"employee-platform/middleware"
	"employee-platform/model"
	"employee-platform/repository"
	"employee-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Task 3: load environment-based configuration. .env is optional in deployed environments.
	_ = godotenv.Load()
	cfg := config.Load()

	db, err := database.Open(cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	// Database setup: automatically create/update the two project tables.
	if err := db.AutoMigrate(&model.User{}, &model.Department{}, &model.Employee{}); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	// Dependency injection: repositories -> services -> REST/GraphQL handlers.
	userRepo := repository.NewUserRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTHours)
	employeeService := service.NewEmployeeService(employeeRepo, departmentRepo)
	departmentService := service.NewDepartmentService(departmentRepo)

	// Create a predictable demo admin if the configured email does not exist.
	// Change DEMO_ADMIN_EMAIL / DEMO_ADMIN_PASSWORD before using outside local development.
	if os.Getenv("SEED_DEMO_ADMIN") == "true" {
		if err := authService.SeedAdmin(
			envOr("DEMO_ADMIN_EMAIL", "admin@example.com"),
			envOr("DEMO_ADMIN_PASSWORD", "Admin123!"),
			envOr("DEMO_ADMIN_NAME", "Demo Admin"),
		); err != nil {
			log.Fatalf("seed admin failed: %v", err)
		}
	}

	r := gin.Default()
	r.Use(middleware.RequestLogger())

	// Task 3: inject the same services into the GraphQL server used by REST.
	gqlServer, err := graphqlserver.NewServer(employeeService, departmentService, authService)
	if err != nil {
		log.Fatalf("GraphQL schema initialization failed: %v", err)
	}
	// Use the same graphql-go/handler pattern as the previous project to show classic GraphiQL.
	gqlHTTP := graphqlserver.WithAuthContext(cfg.JWTSecret, graphqlserver.NewHTTPHandler(gqlServer))
	handler.RegisterRoutes(r, employeeService, departmentService, authService, gqlHTTP, cfg.JWTSecret)

	port := cfg.ServerPort
	log.Printf("Employee platform listening on http://localhost:%s", port)
	log.Printf("GraphiQL IDE: http://localhost:%s/graphql", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
