package main

import (
	"io"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		os.Exit(1)
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	handler := slog.NewJSONHandler(mw, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Sets the minimum log level
	})
	logger := slog.New(handler)

	slog.SetDefault(logger)

	// Example logs with structured key-value pairs
	slog.Debug("Initializing database connection", "host", "localhost", "port", 5432)
	slog.Info("Application started", "version", "1.0.0", "env", "production")
	slog.Info("User logged in", "user", "admin", "ip", "192.168.1.50")
	slog.Warn("High memory usage", "usage_percent", 88.5, "free_mb", 256)
	slog.Error("Database query failed", "query_id", 402, "error", "connection timeout")

	// Grouping related structured data
	userLogger := logger.WithGroup("user_context")
	userLogger.Info("Profile updated", "id", 42, "role", "editor")
}
