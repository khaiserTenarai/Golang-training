package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"employee-management/config"
	"employee-management/controller"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()
	logger := logging.New()

	if cfg.DBPassword == "" {
		logger.Error("DB_PASSWORD is missing from .env")
		return
	}

	if cfg.JWTSecret == "" {
		logger.Error("JWT_SECRET is missing from .env")
		return
	}

	expiryMinutes, err := strconv.Atoi(cfg.JWTExpiryMinutes)
	if err != nil || expiryMinutes <= 0 {
		logger.Error("JWT_EXPIRY_MINUTES must be a positive number")
		return
	}

	// Connect to PostgreSQL.
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		logger.Error("PostgreSQL connection failed: " + err.Error())
		return
	}

	logger.Info("Connected to PostgreSQL")

	// Create the application layers using your reference structure.
	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)

	ctl := controller.NewEmployeeController(
		svc,
		cfg.JWTSecret,
		expiryMinutes,
	)

	// Set Gin to release mode when configured for production.
	if cfg.LogLevel == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Global middleware order.
	r.Use(middleware.RequestID())
	r.Use(middleware.CORS())
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.CustomMiddleware())
	r.Use(middleware.SecondMiddleware())
	r.Use(middleware.Logging(logger))

	// Register routes.
	ctl.RegisterRoutes(r)

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info(
		"Employee Management API started on http://localhost:" +
			cfg.ServerPort,
	)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}
