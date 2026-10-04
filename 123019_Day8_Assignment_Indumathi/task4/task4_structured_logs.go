package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("user login", "username", "john_doe", "status", "success")
	logger.Error("database connection failed", "attempt", 3)
}