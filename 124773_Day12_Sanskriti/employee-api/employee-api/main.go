package main

import (
	"context"
	"employee-api/config"
	"employee-api/handler"
	"employee-api/middleware"
	"employee-api/repository"
	"employee-api/service"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Configuration error: ", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)
	db, err := pgxpool.New(ctx, cfg.DatabaseURL())
	cancel()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	ctx, cancel = context.WithTimeout(
		context.Background(), 5*time.Second,
	)
	err = db.Ping(ctx)
	cancel()
	if err != nil {
		log.Fatal("Database ping failed: ", err)
	}
	log.Println("Connected to PostgreSQL")

	// Dependency injection
	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	h := handler.NewEmployeeHandler(svc)
	auth := middleware.NewAuth(cfg)

	mux := http.NewServeMux()
	// Public login
	mux.HandleFunc("POST /login", auth.Login)

	// Authentication wrapper
	protect := func(next http.Handler) http.Handler {
		return middleware.RequireAuth(auth.Secret)(next)
	}

	// Both admin and normal user can read employees
	mux.Handle("GET /employees",
		protect(http.HandlerFunc(h.Collection)))

	mux.Handle("GET /employees/{id}",
		protect(http.HandlerFunc(h.Detail)))

	// Only admin can modify employees
	adminOnly := func(next http.Handler) http.Handler {
		return protect(middleware.RequireRole("admin")(next))
	}

	mux.Handle("POST /employees",
		adminOnly(http.HandlerFunc(h.Collection)))

	mux.Handle("PUT /employees/{id}",
		adminOnly(http.HandlerFunc(h.Detail)))

	mux.Handle("DELETE /employees/{id}",
		adminOnly(http.HandlerFunc(h.Detail)))

	// Middleware chain
	var app http.Handler = mux

	app = middleware.Logging(app)
	app = middleware.RequestID(app)

	server := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		Handler:           app,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Log level: %s", cfg.LogLevel)
	log.Printf("Server running on port %s", cfg.ServerPort)

	log.Fatal(server.ListenAndServe())
}
