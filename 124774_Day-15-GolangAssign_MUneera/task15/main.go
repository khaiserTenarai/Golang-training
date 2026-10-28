package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"employee-management/auth"
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

	if cfg.JWTSecret == "" {
		logger.Error("JWT_SECRET is missing from .env")
		return
	}

	if cfg.AdminPasswordHash == "" || cfg.UserPasswordHash == "" {
		logger.Error("ADMIN_PASSWORD_HASH and USER_PASSWORD_HASH must be set in .env")
		return
	}

	expiryMinutes, err := strconv.Atoi(cfg.JWTExpiryMinutes)
	if err != nil || expiryMinutes <= 0 {
		expiryMinutes = 60
	}

	// Build the PostgreSQL connection string.
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

	// Create the layers.
	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)

	authService := auth.New(
		cfg.JWTSecret,
		expiryMinutes,
		cfg.AdminUsername,
		cfg.AdminPasswordHash,
		cfg.UserUsername,
		cfg.UserPasswordHash,
	)

	ctl := controller.NewEmployeeController(svc, authService)

	// Create Gin and register global middleware.
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(
		middleware.RequestID(),
		middleware.Logging(logger),
		middleware.Recovery(logger),
		middleware.CORS(),
	)

	// Register application endpoints.
	ctl.RegisterRoutes(r)

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Employee Management API started at http://localhost:" + cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}
