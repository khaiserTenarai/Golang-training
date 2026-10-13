package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"employee-api-go/internal/auth"
	"employee-api-go/internal/handler"
	"employee-api-go/internal/middleware"
	"employee-api-go/internal/repository"
	"employee-api-go/internal/router"
	"employee-api-go/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "development-secret-change-before-production"
	}

	repo := repository.NewMemoryEmployeeRepository()
	employeeService := service.NewEmployeeService(repo)
	tokenManager := auth.NewTokenManager([]byte(secret), 8*time.Hour)

	employeeHandler := handler.NewEmployeeHandler(employeeService)
	authHandler := handler.NewAuthHandler(tokenManager)
	systemHandler := handler.NewSystemHandler(repo)

	mux := router.New(employeeHandler, authHandler, systemHandler, tokenManager)
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           middleware.RequestID(middleware.Logging(mux)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("employee API listening on :%s", port)
	log.Fatal(server.ListenAndServe())
}
