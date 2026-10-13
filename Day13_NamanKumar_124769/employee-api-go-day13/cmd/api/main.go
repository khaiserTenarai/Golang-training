package main

import (
	"log"
	"net/http"

	"employee-api-go/internal/app"
	"employee-api-go/internal/auth"
	"employee-api-go/internal/config"
	"employee-api-go/internal/handler"
	"employee-api-go/internal/middleware"
	"employee-api-go/internal/repository"
	"employee-api-go/internal/router"
	"employee-api-go/internal/service"
)

func main() {
	cfg, secrets, err := config.Load()
	if err != nil {
		log.Fatalf("configuration_error error=%q", err.Error())
	}

	repo := repository.NewMemoryEmployeeRepository()
	employeeService := service.NewEmployeeService(repo)
	tokenManager := auth.NewTokenManager([]byte(secrets.JWTSecret), cfg.TokenTTL)

	employeeHandler := handler.NewEmployeeHandler(employeeService)
	authHandler := handler.NewAuthHandler(tokenManager)
	systemHandler := handler.NewSystemHandler(repo)

	mux := router.New(employeeHandler, authHandler, systemHandler, tokenManager, cfg.EnableEmployeeDelete)
	httpServer := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           middleware.RequestID(middleware.Logging(mux)),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	log.Printf("startup_complete environment=%q port=%q log_level=%q employee_delete_enabled=%t",
		cfg.AppEnv, cfg.Port, cfg.LogLevel, cfg.EnableEmployeeDelete)

	server := app.Server{HTTP: httpServer, ShutdownTimeout: cfg.ShutdownTimeout}
	if err := server.Run(); err != nil {
		log.Fatalf("server_error error=%q", err.Error())
	}
}
