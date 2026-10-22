package main

import (
	"fmt"

	"employee-api/config"
	"employee-api/controllers"
	"employee-api/repository"
	"employee-api/router"
	"employee-api/service"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Initialize Repository Layer
	repo := repository.NewMemoryRepository()

	// 3. Initialize Service Layer
	svc := service.NewEmployeeService(repo)

	// 4. Initialize Controller Layer
	controller := controllers.NewEmployeeController(svc)

	// 5. Setup Router
	r := router.SetupRouter(controller)

	// 6. Start Server
	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Starting Employee API on %s (Env: %s)...", serverAddr, cfg.Env)

	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
