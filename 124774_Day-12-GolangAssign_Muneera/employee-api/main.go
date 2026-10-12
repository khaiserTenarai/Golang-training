package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"day12/config"
	"day12/handler"
	"day12/middleware"
	"day12/repository"
	"day12/service"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	ctx := context.Background()

	dbPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Unable to create database pool: ", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	log.Println("Database connected successfully")

	repo := repository.NewEmployeeRepository(dbPool)
	empService := service.NewEmployeeService(repo)
	empHandler := handler.NewEmployeeHandler(empService, dbPool)

	mux := http.NewServeMux()
	empHandler.RegisterRoutes(mux)

	// Employee routes require authentication.
	protected := http.NewServeMux()

	employeeHandler := middleware.RequireAuth(
		cfg.UserToken,
		cfg.AdminToken,
		middleware.AdminWrites(mux),
	)

	protected.Handle("/employees", employeeHandler)
	protected.Handle("/employees/", employeeHandler)

	// Health endpoints are public.
	protected.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})

	protected.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})

	// Request ID runs before logging.
	server := middleware.RequestID(
		middleware.Logging(protected),
	)

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, server))
}
