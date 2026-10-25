package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"employee-management-microservices/database"
	"employee-management-microservices/services/department/config"
	"employee-management-microservices/services/department/controller"
	"employee-management-microservices/services/department/model"
	"employee-management-microservices/services/department/repository"
	"employee-management-microservices/services/department/service"
	"employee-management-microservices/shared/correlation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load local environment variables.
	_ = godotenv.Load()

	cfg := config.Load()

	// Connect to this service's database.
	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatal("Department database connection failed: ", err)
	}

	// Create the departments table if it does not exist.
	if err := db.AutoMigrate(&model.Department{}); err != nil {
		log.Fatal("Department migration failed: ", err)
	}

	repo := repository.New(db)
	svc := service.New(repo)
	handler := controller.New(svc)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(correlation.Middleware())

	// Basic health endpoints.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	router.GET("/ready", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	// Department API routes.
	departments := router.Group("/departments")
	{
		departments.POST("", handler.Create)
		departments.GET("", handler.GetAll)
		departments.GET("/:id", handler.GetByID)
		departments.PUT("/:id", handler.Update)
		departments.DELETE("/:id", handler.Delete)
	}

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Stop the server gracefully when interrupted.
	go func() {
		log.Printf("Department Service running on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	_ = server.Close()

	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
}
