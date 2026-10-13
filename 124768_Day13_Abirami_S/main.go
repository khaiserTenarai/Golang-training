package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	logger := newLogger(cfg.LogLevel)
	logger.Info("starting Employee API", "environment", cfg.AppEnv, "port", cfg.Port)

	secretValues, err := loadSecrets()
	if err != nil {
		logger.Error("secret configuration error", "error", err)
		os.Exit(1)
	}

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	startupCtx, startupCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer startupCancel()
	db, err := pgxpool.New(startupCtx, cfg.DatabaseURL)
	if err != nil {
		logger.Error("database initialization failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	if err := db.Ping(startupCtx); err != nil {
		logger.Error("database is unavailable", "error", err)
		os.Exit(1)
	}

	svc := newEmployeeService(db)
	api := newAPI(svc, db, secretValues, logger, cfg.EmployeeExportEnabled)
	server := &http.Server{Addr: ":" + cfg.Port, Handler: api.routes(), ReadHeaderTimeout: 5 * time.Second}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("server listening", "address", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	select {
	case <-signalCtx.Done():
		logger.Info("shutdown signal received")
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server stopped unexpectedly", "error", err)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeoutSeconds)*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown timed out; forcing close", "error", err)
		_ = server.Close()
	}
	logger.Info("Employee API stopped")
}
