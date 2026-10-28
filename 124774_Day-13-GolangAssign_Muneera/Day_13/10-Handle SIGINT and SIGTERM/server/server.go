package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"employee-api/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// New creates the HTTP server and registers its routes.
func New(cfg *config.Config, db *gorm.DB) *http.Server {
	// Use release mode in production.
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create the Gin router and middleware.
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Check whether the application is running.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"environment": cfg.AppEnv,
		})
	})

	// Check whether PostgreSQL is reachable.
	router.GET("/health/db", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "database unavailable",
			})
			return
		}

		// Stop waiting if the database check takes too long.
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			2*time.Second,
		)
		defer cancel()

		if err := sqlDB.PingContext(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "database unavailable",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "database connected",
		})
	})

	// Create the HTTP server with a header timeout.
	return &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

// Start runs the server until it fails or shutdown is requested.
func Start(ctx context.Context, httpServer *http.Server) error {
	// Receive the HTTP server's result.
	serverErrors := make(chan error, 1)

	// Run the HTTP server in a goroutine.
	go func() {
		err := httpServer.ListenAndServe()

		// ErrServerClosed is expected during normal shutdown.
		if errors.Is(err, http.ErrServerClosed) {
			serverErrors <- nil
			return
		}

		serverErrors <- err
	}()

	// Wait for a server error or a shutdown request.
	select {
	case err := <-serverErrors:
		if err != nil {
			return fmt.Errorf("HTTP server failed: %w", err)
		}
		return nil

	case <-ctx.Done():
		// Allow active requests up to ten seconds to finish.
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		// Stop accepting new requests and wait for active requests.
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			// Force-close connections if graceful shutdown fails.
			_ = httpServer.Close()
			return fmt.Errorf("graceful shutdown failed: %w", err)
		}

		// Wait for the server goroutine to finish.
		if err := <-serverErrors; err != nil {
			return fmt.Errorf("HTTP server failed during shutdown: %w", err)
		}

		return nil
	}
}