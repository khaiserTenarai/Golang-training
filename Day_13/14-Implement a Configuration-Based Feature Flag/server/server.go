package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"employee-api/config"
	"employee-api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// New creates the HTTP server and registers routes.
func New(
	cfg *config.Config,
	db *gorm.DB,
	employeeController controller.EmployeeController,
) *http.Server {
	router := gin.Default()

	// Basic health check.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Check whether the database is reachable.
	router.GET("/health/db", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "database unavailable",
			})
			return
		}

		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			3*time.Second,
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

	// Register POST only when employee creation is enabled.
	if cfg.EnableEmployeeCreate {
		router.POST("/employees", employeeController.Create)
	} else {
		// This message helps during development.
		fmt.Println("Feature flag OFF: employee creation is disabled")
	}

	// Employee lookup is available regardless of the feature flag.
	router.GET("/employees/:id", employeeController.GetByID)

	return &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

// Start runs the HTTP server and shuts it down gracefully.
func Start(ctx context.Context, httpServer *http.Server) error {
	serverErrors := make(chan error, 1)

	// Run the server in the background.
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
			return
		}

		serverErrors <- nil
	}()

	select {
	case err := <-serverErrors:
		if err != nil {
			return fmt.Errorf("HTTP server failed: %w", err)
		}
		return nil

	case <-ctx.Done():
		fmt.Println("Shutdown signal received")

		// Give active requests time to finish.
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Println("Graceful shutdown failed; forcing close")
			_ = httpServer.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		fmt.Println("HTTP server stopped gracefully")
		return nil
	}
}