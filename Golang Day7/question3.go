package main

import (
	"log/slog"
	"os"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Debug("Loading configuration files...")
	slog.Info("Application started successfully.", "port", 8080)
	slog.Warn("Memory usage is getting high.", "usage", "85%")
	slog.Error("Failed to connect to the database.", "db", "users")
}