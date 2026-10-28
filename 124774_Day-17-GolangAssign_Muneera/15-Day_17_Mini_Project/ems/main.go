package main

import (
	"context"
	"employee-management/auth"
	"employee-management/config"
	"employee-management/controller"
	"employee-management/database"
	"errors"
	"log"
	"net/http"
	"time"

	graph "employee-management/graphql"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("configuration error: ", err)
	}

	// Initialize the shared PostgreSQL connection through GORM.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("database initialization failed: ", err)
	}
	defer sqlDB.Close()

	// Create or update the employee table and indexes.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	// Construct one repository and one shared service.
	repo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(repo)

	// Initialize JWT authentication for the two configured users.
	authService, err := auth.New(
		cfg.JWTSecret,
		cfg.JWTExpiryMinutes,
		cfg.AdminUsername, cfg.AdminPassword,
		cfg.UserUsername, cfg.UserPassword,
	)
	if err != nil {
		log.Fatal("authentication initialization failed: ", err)
	}

	authController := controller.NewAuthController(authService)
	employeeController := controller.NewEmployeeController(employeeService)

	// Construct GraphQL with the exact same employee service.
	graphServer, err := graph.NewServer(employeeService)
	if err != nil {
		log.Fatal("GraphQL initialization failed: ", err)
	}
	graphHandler := graph.NewHTTPHandler(graphServer, authService)

	mux := http.NewServeMux()

	// Public routes: login, liveness, and readiness.
	mux.HandleFunc("/auth/login", authController.Login)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"UP"}`))
	})

	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := timeLimitContext(r)
		defer cancel()

		if err := sqlDB.PingContext(ctx); err != nil {
			http.Error(w, `{"status":"NOT_READY"}`, http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"READY"}`))
	})

	// REST employee endpoints require a valid JWT.
	employeeMux := http.NewServeMux()
	employeeController.RegisterRoutes(employeeMux)

	protectedREST := middleware.RequireAuth(authService, employeeMux)
	mux.Handle("/employees", protectedREST)
	mux.Handle("/employees/", protectedREST)

	// GraphQL queries and mutations require a valid JWT on POST requests.
	mux.Handle("/graphql", graphHandler)

	// Logging wraps all routes, including health and login.
	handler := middleware.Logging(mux)

	server := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("REST API: http://localhost:%s/employees", cfg.ServerPort)
	log.Printf("Login:     http://localhost:%s/auth/login", cfg.ServerPort)
	log.Printf("GraphQL:   http://localhost:%s/graphql", cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) &&
		!errors.Is(err, gorm.ErrInvalidDB) {
		log.Fatal("HTTP server failed: ", err)
	}
}

// timeLimitContext gives readiness checks a short timeout.
func timeLimitContext(r *http.Request) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), 3*time.Second)
}
