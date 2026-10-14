/*
Question 15: Day 13 Mini Project - Build a configurable Employee API with
environment configuration, secrets, logging, feature flags and graceful shutdown.

Layers: controller -> service -> dao (plus model, config and logger).

Run with: go run .
Every setting is asked at startup. Press Enter to accept the value in brackets,
which comes from an environment variable when one is set.
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"employeeapi/config"
	"employeeapi/controller"
	"employeeapi/dao"
	"employeeapi/logger"
	"employeeapi/service"
)

func main() {
	prompter := config.NewPrompter()

	fmt.Println("=== Employee API setup ===")
	cfg, err := config.Load(prompter)
	if err != nil {
		fmt.Println("configuration error:", err)
		os.Exit(1)
	}
	secrets := config.LoadSecrets(prompter, cfg)

	problems := append(cfg.Validate(), secrets.Validate(cfg)...)
	if len(problems) > 0 {
		fmt.Println("configuration is not valid:")
		for _, problem := range problems {
			fmt.Println(" -", problem)
		}
		os.Exit(1)
	}

	log := logger.New(cfg)
	log.Info("configuration loaded",
		"env", cfg.Env,
		"port", cfg.Port,
		"log_level", cfg.LogLevel,
		"shutdown_timeout_secs", cfg.ShutdownTimeout,
		"feature_salary_report", cfg.Features.SalaryReport,
		"feature_delete_employee", cfg.Features.DeleteEmployee,
		"secrets", secrets,
	)

	employeeDAO := dao.NewMemoryEmployeeDAO()
	employeeService := service.NewEmployeeService(employeeDAO, cfg.Features, log)
	employeeController := controller.NewEmployeeController(employeeService)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: controller.LogRequests(log, controller.RegisterRoutes(employeeController, secrets.APIKey)),
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Info("server starting", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server failed", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("shutdown did not finish in time", "error", err)
		return
	}
	log.Info("server stopped cleanly")
}
