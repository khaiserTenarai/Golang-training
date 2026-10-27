package logger

import (
	"log/slog"
	"os"
	"strings"

	"order-management/config"
)

var appLogger *slog.Logger

// Init initializes the application logger from configuration.
func Init() {
	level := slog.LevelInfo

	switch strings.ToLower(config.Get("LOG_LEVEL", "info")) {
	case "debug":
		level = slog.LevelDebug
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	appLogger = slog.New(handler)
	slog.SetDefault(appLogger)
}

// Get returns the initialized logger, or a default logger.
func Get() *slog.Logger {
	if appLogger == nil {
		return slog.Default()
	}
	return appLogger
}
