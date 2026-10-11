// Package main starts the Gin Employee Management API.
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/config"
	"employee-management/controller"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"
)

func main() {
	cfg := config.Load()
	logger := logging.New()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		logger.Error(err.Error())
		return
	}

	// Auto-create table
	_, _ = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		department VARCHAR(50) NOT NULL,
		salary NUMERIC(10,2) NOT NULL
	);`)

	// Layer Dependency Injection
	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctl := controller.NewEmployeeController(svc)

	// Point 1: Create Gin Application
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging(logger))

	// Point 14: Implement custom 404 and 405 handlers
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Resource path not found",
		})
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    405,
			"message": "HTTP method not allowed on this route",
		})
	})

	// Register Routes
	ctl.RegisterRoutes(r)

	// Point 15: Run Gin Application
	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Gin Employee Management API started on http://localhost:" + cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}