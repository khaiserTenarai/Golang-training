package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"employee-api/config"
	"employee-api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// New registers the routes and creates the HTTP server.
func New(
	cfg *config.Config,
	db *gorm.DB,
	employeeController controller.EmployeeController,
) *http.Server {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Application health check.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Database health check.
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

	// Feature flag controls employee creation.
	if cfg.EnableEmployeeCreate {
		router.POST("/employees", employeeController.Create)
		log.Println("Feature flag enabled: POST /employees")
	} else {
		log.Println("Feature flag disabled: POST /employees is unavailable")
	}

	// Employee lookup is always available.
	router.GET("/employees/:id", employeeController.GetByID)

	return &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

// Start runs the server and shuts it down gracefully.
func Start(
	ctx context.Context,
	httpServer *http.Server,
) error {
	serverErrors := make(chan error, 1)

	go func() {
		err := httpServer.ListenAndServe()

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- fmt.Errorf("HTTP server failed: %w", err)
			return
		}

		serverErrors <- nil
	}()

	select {
	case err := <-serverErrors:
		return err

	case <-ctx.Done():
		log.Println("Shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
			_ = httpServer.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		// Wait for ListenAndServe to return.
		if err := <-serverErrors; err != nil {
			return err
		}

		log.Println("HTTP server stopped gracefully")
		return nil
	}
}