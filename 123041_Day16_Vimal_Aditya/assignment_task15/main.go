package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"employee-management/config"
	gqlserver "employee-management/graphql"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()
	logger := logging.New()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		logger.Error(err.Error())
		return
	}

	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)

	gs, err := gqlserver.NewServer(svc)
	if err != nil {
		logger.Error("GraphQL initialization failed: " + err.Error())
		return
	}

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging(logger))

	// Direct GraphQL Endpoint binding via Gin
	h := gqlserver.NewHTTPHandler(gs)
	r.Any("/graphql", gin.WrapH(h))

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Gin GraphQL Platform running on http://localhost:" + cfg.ServerPort + "/graphql")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}