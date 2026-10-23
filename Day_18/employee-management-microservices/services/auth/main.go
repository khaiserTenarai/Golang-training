package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"employee-management-microservices/database"
	"employee-management-microservices/services/auth/config"
	"employee-management-microservices/services/auth/controller"
	"employee-management-microservices/services/auth/model"
	"employee-management-microservices/services/auth/repository"
	"employee-management-microservices/services/auth/service"
	"employee-management-microservices/shared/correlation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	if len(cfg.JWTSecret) < 32 {
		log.Fatal("JWT_SECRET must contain at least 32 characters")
	}

	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatal("Auth database connection failed: ", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Auth migration failed: ", err)
	}

	// Create the initial admin only when a password is configured.
	adminPassword := os.Getenv("INITIAL_ADMIN_PASSWORD")
	if adminPassword != "" {
		var count int64
		db.Model(&model.User{}).Where("username = ?", "admin").Count(&count)

		if count == 0 {
			hash, err := bcrypt.GenerateFromPassword(
				[]byte(adminPassword), bcrypt.DefaultCost,
			)
			if err != nil {
				log.Fatal(err)
			}

			admin := model.User{
				Username:     "admin",
				PasswordHash: string(hash),
				Role:         "ADMIN",
				Active:       true,
			}
			if err := db.Create(&admin).Error; err != nil {
				log.Fatal("Failed to create initial admin: ", err)
			}
		}
	}

	repo := repository.New(db)
	svc := service.New(repo, cfg.JWTSecret)
	handler := controller.New(svc)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(correlation.Middleware())

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

	router.POST("/auth/login", handler.Login)

	// Keep registration private during initial setup.
	// Protect this endpoint before enabling it in a deployed system.
	router.POST("/auth/register", handler.Register)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("Authentication Service running on port %s", cfg.Port)
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
