package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"employee-rest-api/config"
	"employee-rest-api/controller"
	"employee-rest-api/database"
	"employee-rest-api/repository"
	"employee-rest-api/routes"
	"employee-rest-api/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	if err := database.CreateSchema(db); err != nil {
		log.Fatalf("create database schema: %v", err)
	}

	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctl := controller.NewEmployeeController(svc)

	mux := http.NewServeMux()
	routes.Register(mux, ctl)

	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           loggingMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Employee API running at http://localhost:%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("server error: %v", err)
		os.Exit(1)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("completed %s %s in %s", r.Method, r.URL.Path, time.Since(start))
	})
}
