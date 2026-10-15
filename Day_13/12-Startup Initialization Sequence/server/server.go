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

// New creates the HTTP server and registers the routes.
func New(
	cfg *config.Config,
	db *gorm.DB,
	employeeController controller.EmployeeController,
) *http.Server {
	// Use release mode in production.
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a router with request logging and panic recovery.
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Application health check.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"environment": cfg.AppEnv,
		})
	})

	// Database health check.
	router.GET("/health/db", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "database unavailable",
			})
			return
		}

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

	// Employee routes.
	router.POST("/employees", employeeController.Create)
	router.GET("/employees/:id", employeeController.GetByID)

	// Configure the HTTP server.
	return &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

// Start runs the HTTP server until an error or shutdown request occurs.
func Start(ctx context.Context, httpServer *http.Server) error {
	serverErrors := make(chan error, 1)

	// Listen for incoming HTTP requests in the background.
	go func() {
		err := httpServer.ListenAndServe()

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
		// Give active requests ten seconds to finish.
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		// Stop accepting new requests and finish active requests.
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			_ = httpServer.Close()
			return fmt.Errorf("graceful shutdown failed: %w", err)
		}

		// Wait for the HTTP server goroutine to finish.
		if err := <-serverErrors; err != nil {
			return fmt.Errorf("HTTP server shutdown: %w", err)
		}

		return nil
	}
}