package main

import (
	"fmt"

	"employee-platform/config"
	"employee-platform/controllers"
	"employee-platform/middleware"
	"employee-platform/repository"
	"employee-platform/router"
	"employee-platform/schema"
	"employee-platform/service"
)

func main() {
	cfg := config.LoadConfig()

	// 1. Repository
	repo := repository.NewMemoryRepository()

	// 2. Service
	svc := service.NewEmployeeService(repo)

	// 3. Controllers
	restController := controllers.NewEmployeeController(svc)

	// 4. GraphQL Schema & Resolver Injection
	gqlSchema, err := schema.BuildSchema(svc)
	if err != nil {
		panic(fmt.Sprintf("Failed to build GraphQL schema: %v", err))
	}

	// 5. Setup Router
	r := router.SetupRouter(restController, gqlSchema, cfg.JWTSecret)

	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Starting Day 17 REST + GraphQL Platform on %s (Env: %s)...
", serverAddr, cfg.Env)

	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
