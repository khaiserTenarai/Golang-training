package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"day18/config"
	"day18/consumer"
	gql "day18/graphql"
	"day18/logging"
	"day18/middleware"
	"day18/repository"
	"day18/service"
	"day18/worker"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type OnboardRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int64  `json:"department_id" binding:"required"`
}

func main() {
	cfg := config.LoadConfig()
	logger := logging.NewLogger()

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	poolConfig, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		logger.Error("Unable to parse DB connection string: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.Error("Unable to create connection pool: %v", err)
		return
	}
	defer pool.Close()

	empRepo := repository.NewEmployeeRepository(pool)
	empService := service.NewEmployeeService(empRepo, logger)

	// Start Background Kafka Outbox Relay Worker
	kafkaBroker := "localhost:9092"
	outboxRelay := worker.NewOutboxRelay(pool, kafkaBroker, logger)
	go outboxRelay.Start(ctx)

	// Start Background Saga Compensating Consumer
	sagaConsumer := consumer.NewSagaConsumer(kafkaBroker, empRepo, logger)
	sagaConsumer.Start(ctx)

	gqlServer, err := gql.NewServer(empService)
	if err != nil {
		logger.Error("Failed to build GraphQL schema: %v", err)
		return
	}
	gqlHandler := gql.NewHTTPHandler(gqlServer)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.GinLoggingMiddleware(logger))

	// Prometheus Endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/employees", func(c *gin.Context) {
			var req OnboardRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			emp, err := empService.OnboardEmployee(c.Request.Context(), req.FirstName, req.LastName, req.Email, req.DepartmentID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to onboard employee"})
				return
			}

			c.JSON(http.StatusCreated, emp)
		})

		api.GET("/employees/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format, must be an integer"})
				return
			}

			emp, err := empService.GetEmployee(c.Request.Context(), id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				return
			}

			c.JSON(http.StatusOK, emp)
		})
	}

	r.Any("/graphql", gin.WrapH(gqlHandler))

	serverAddr := fmt.Sprintf(":%s", cfg.AppPort)
	logger.Info("Employee Microservice running on port %s...", cfg.AppPort)
	if err := r.Run(serverAddr); err != nil {
		logger.Error("Server crashed: %v", err)
	}
}