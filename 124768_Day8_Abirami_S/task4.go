package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger.Info("Employee created",
		"id", 101,
		"name", "Tom",
		"salary", 70000,
	)
	logger.Warn("Employee salary is low",
		"id", 102,
		"name", "John",
		"salary", 700,
	)
	logger.Error("Employee data not found", "id", 103)
}
