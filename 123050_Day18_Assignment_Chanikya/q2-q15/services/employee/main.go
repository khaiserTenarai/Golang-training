package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"employee-management-microservices/database"
	"employee-management-microservices/events/outbox"
	"employee-management-microservices/events/worker"
	"employee-management-microservices/services/employee/client"
	"employee-management-microservices/services/employee/config"
	"employee-management-microservices/services/employee/controller"
	"employee-management-microservices/services/employee/model"
	"employee-management-microservices/services/employee/repository"
	"employee-management-microservices/services/employee/service"
	"employee-management-microservices/shared/correlation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatal("Employee database connection failed: ", err)
	}

	// Create both the employee and outbox tables.
	if err := db.AutoMigrate(&model.Employee{}, &outbox.Event{}); err != nil {
		log.Fatal("Employee migration failed: ", err)
	}

	repo := repository.New(db)
	departmentClient := client.NewDepartmentClient(cfg.DepartmentServiceURL)
	svc := service.New(repo, departmentClient)
	handler := controller.New(svc)

	// Start the asynchronous event worker.
	workerCtx, cancelWorker := context.WithCancel(context.Background())
	defer cancelWorker()

	eventURL := os.Getenv("EVENT_RECEIVER_URL")
	if eventURL != "" {
		eventWorker := worker.New(db, eventURL)
		go eventWorker.Run(workerCtx)
		log.Println("Outbox worker started")
	} else {
		log.Println("EVENT_RECEIVER_URL is empty; outbox events will remain pending")
	}

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

	employees := router.Group("/employees")
	{
		employees.POST("", handler.Create)
		employees.GET("", handler.GetAll)
		employees.GET("/:id", handler.GetByID)
		employees.PUT("/:id", handler.Update)
		employees.DELETE("/:id", handler.Delete)
	}

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("Employee Service running on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("Employee server error: %v", err)
			cancelWorker()
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	cancelWorker()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second,
	)
	defer cancel()

	_ = server.Shutdown(shutdownCtx)

	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
}
