package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"employee-api/config"
	"employee-api/controller"
	"employee-api/middleware"
	"employee-api/model"
	"employee-api/repository"
	"employee-api/service"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "fatal:", err)
		os.Exit(1)
	}
}

// run contains the whole application lifecycle:
//
//	STARTUP   1 configuration + logging -> 2 repository -> 3 services -> 4 HTTP handlers -> 5 listen
//	RUNNING   wait for SIGINT (Ctrl+C) or SIGTERM
//	SHUTDOWN  1 stop taking traffic -> 2 drain in-flight requests -> 3 close repository -> 4 done
//
// Resources are shut down in the REVERSE order of how they were started.
func run() error {
	// ---------------- STARTUP 1/5: configuration (fail fast) and logging ----------------
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("invalid configuration:\n%w", err)
	}
	logger := newLogger(cfg)
	slog.SetDefault(logger)
	logger.Info("startup 1/5: configuration loaded and validated", "config", cfg.Summary())

	// Cancelled by Ctrl+C (SIGINT) or SIGTERM (what Docker/Kubernetes send on stop).
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Deadline for everything done during startup.
	initCtx, cancelInit := context.WithTimeout(ctx, cfg.Server.StartupTimeout)
	defer cancelInit()

	// ---------------- STARTUP 2/5: repository ----------------
	repo := repository.NewMemoryEmployeeRepository()
	if cfg.Features.SeedData {
		if err := seed(initCtx, repo); err != nil {
			return fmt.Errorf("seeding sample data: %w", err)
		}
	}
	logger.Info("startup 2/5: repository ready", "seedData", cfg.Features.SeedData)

	// ---------------- STARTUP 3/5: services ----------------
	// Configuration is copied into the service Options; the service never reads the environment.
	employeeService := service.NewEmployeeService(repo, service.Options{
		RequestTimeout:  cfg.Server.RequestTimeout,
		DefaultPageSize: cfg.Pagination.DefaultPageSize,
		MaxPageSize:     cfg.Pagination.MaxPageSize,
		SearchEnabled:   cfg.Features.SearchEnabled, // feature flag
	}, logger)
	logger.Info("startup 3/5: services ready")

	// ---------------- STARTUP 4/5: controllers, routes, middleware ----------------
	employeeController := controller.NewEmployeeController(employeeService, logger, cfg.Features.ExposeSalary)
	health := controller.NewHealthController(employeeService, cfg.Summary)

	mux := http.NewServeMux()
	employeeController.Register(mux, middleware.APIKey(cfg.Secrets.APIKey))
	mux.HandleFunc("GET /health", health.Live)
	mux.HandleFunc("GET /ready", health.Ready)
	if !cfg.IsProduction() { // debugging helpers never exist in production
		mux.HandleFunc("GET /config", health.Config)
		mux.HandleFunc("GET /debug/slow", health.Slow)
	}
	handler := middleware.Chain(mux,
		middleware.RequestID,         // outermost: id exists for everything below
		middleware.Logging(logger),   // logs the final status code, also after a recovered panic
		middleware.Recover(logger),   // innermost: panic -> 500
	)
	logger.Info("startup 4/5: routes ready")

	// ---------------- STARTUP 5/5: HTTP server ----------------
	// Requests get their context from baseCtx. Normally it is never cancelled; it is cancelled only
	// when graceful shutdown runs out of time, to abort requests that are still running.
	baseCtx, cancelBase := context.WithCancel(context.Background())
	defer cancelBase()

	srv := &http.Server{
		Addr:              cfg.Server.Addr(),
		Handler:           handler,
		ReadTimeout:       cfg.Server.ReadTimeout,
		ReadHeaderTimeout: cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
		BaseContext:       func(net.Listener) context.Context { return baseCtx },
	}

	ln, err := net.Listen("tcp", srv.Addr) // fails immediately if the port is already in use
	if err != nil {
		return fmt.Errorf("cannot listen on %s: %w", srv.Addr, err)
	}
	serverErr := make(chan error, 1)
	go func() {
		if err := srv.Serve(ln); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()
	health.SetReady(true)
	logger.Info("startup 5/5: server listening - press Ctrl+C to stop", "addr", ln.Addr().String())

	// ---------------- RUNNING ----------------
	select {
	case err := <-serverErr:
		return fmt.Errorf("server failed: %w", err)
	case <-ctx.Done():
		logger.Info("shutdown signal received")
	}

	// ---------------- SHUTDOWN ----------------
	stop() // from now on a second Ctrl+C kills the process immediately

	logger.Info("shutdown 1/4: not ready any more (no new traffic)")
	health.SetReady(false)

	logger.Info("shutdown 2/4: finishing in-flight requests", "timeout", cfg.Server.ShutdownTimeout.String())
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown timed out, closing connections", "error", err)
		cancelBase() // cancel the context of requests that are still running
		_ = srv.Close()
	}

	logger.Info("shutdown 3/4: closing repository")
	if err := repo.Close(); err != nil {
		logger.Error("closing repository", "error", err)
	}

	logger.Info("shutdown 4/4: bye")
	return nil
}

func newLogger(cfg *config.Config) *slog.Logger {
	var level slog.Level
	switch cfg.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: level}

	var h slog.Handler
	if cfg.Log.Format == "json" {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}
	return slog.New(h).With("app", cfg.App.Name, "env", cfg.App.Env)
}

func seed(ctx context.Context, repo repository.EmployeeRepository) error {
	samples := []model.Employee{
		{Name: "Rajesh", Email: "rajesh@example.com", Department: "IT", Salary: 60000},
		{Name: "Amit", Email: "amit@example.com", Department: "HR", Salary: 50000},
		{Name: "Priya", Email: "priya@example.com", Department: "IT", Salary: 75000},
	}
	for _, e := range samples {
		if _, err := repo.Create(ctx, e); err != nil {
			return err
		}
	}
	return nil
}
