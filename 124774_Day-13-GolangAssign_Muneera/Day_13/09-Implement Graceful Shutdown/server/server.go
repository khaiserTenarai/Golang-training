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

// New creates the HTTP server and configures its routes.
func New(cfg *config.Config, db *gorm.DB) *http.Server {
	// Use release mode in production.
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a Gin router without default middleware.
	router := gin.New()

	// Log HTTP requests and recover from unexpected panics.
	router.Use(gin.Logger(), gin.Recovery())

	// Return the application status.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"environment": cfg.AppEnv,
		})
	})

	// Check whether PostgreSQL is reachable.
	router.GET("/health/db", func(c *gin.Context) {
		// Get the underlying SQL connection pool.
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "database unavailable",
			})
			return
		}

		// Stop waiting if the database check takes over two seconds.
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			2*time.Second,
		)
		defer cancel()

		// Ping the database.
		if err := sqlDB.PingContext(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "database unavailable",
			})
			return
		}

		// Return success when the database is reachable.
		c.JSON(http.StatusOK, gin.H{
			"status": "database connected",
		})
	})

	// Create the HTTP server and set basic timeout protection.
	return &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

// Start runs the server and waits for a shutdown signal.
func Start(ctx context.Context, httpServer *http.Server) error {
	// Use a buffered channel to receive the server's result.
	serverErrors := make(chan error, 1)

	// Run the HTTP server in the background.
	go func() {
		err := httpServer.ListenAndServe()

		// A normal shutdown returns ErrServerClosed.
		if errors.Is(err, http.ErrServerClosed) {
			serverErrors <- nil
			return
		}

		// Send any unexpected error to the main goroutine.
		serverErrors <- err
	}()

	// Wait for the server to fail or for shutdown to be requested.
	select {
	case err := <-serverErrors:
		// Report an unexpected server error.
		if err != nil {
			return fmt.Errorf("HTTP server failed: %w", err)
		}
		return nil

	case <-ctx.Done():
		// A shutdown signal was received.
		// Give active requests up to ten seconds to finish.
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		// Stop accepting new requests and wait for active requests.
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			// Force-close connections if graceful shutdown times out.
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
