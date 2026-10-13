package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	cfg, secrets, problems := LoadConfig(reader)
	if len(problems) > 0 {
		fmt.Println("\nThe configuration isn't valid, not starting:")
		for _, p := range problems {
			fmt.Println(" -", p)
		}
		os.Exit(1)
	}

	logger := NewLogger(cfg.LogLevel)
	logger.Infof("starting in %s mode", cfg.Env)
	logger.Debugf("config: %+v", cfg)
	logger.Debugf("secrets: %+v", secrets) // prints ******** - never the real key
	logger.Infof("feature flags: bonus=%v search=%v", cfg.FeatureBonus, cfg.FeatureSearch)

	api := NewAPI(cfg, secrets, logger)
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           api.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Infof("listening on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("server error: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	logger.Infof("shutdown signal received, waiting up to %v for active requests", cfg.ShutdownTimeout)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Errorf("forced to stop early: %v", err)
		return
	}
	logger.Infof("server stopped cleanly")
}
