package main

import (
	"context"
	"net/http"
	"time"

	"employee-management/config"
	"employee-management/controller"
	gqlserver "employee-management/graphql"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()
	logger := logging.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, cfg.DSN())
	if err != nil {
		logger.Error("DB pool error: " + err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		logger.Error("DB ping failed: " + err.Error())
		return
	}
	logger.Info("connected to PostgreSQL")

	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)

	gql, err := gqlserver.NewServer(svc)
	if err != nil {
		logger.Error("GraphQL init error: " + err.Error())
		return
	}

	mux := http.NewServeMux()
	ctrl.RegisterRoutes(mux)

	gqlHandler := gqlserver.NewHTTPHandler(gql)
	mux.Handle("/graphql", gqlHandler)

	handler := middleware.Logging(mux, logger)

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("REST API running on http://localhost:" + cfg.ServerPort + "/employees")
	logger.Info("GraphQL endpoint running on http://localhost:" + cfg.ServerPort + "/graphql")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}