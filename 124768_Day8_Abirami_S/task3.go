package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger.Debug("Debugged")
	logger.Info("Application Started")
	logger.Warn("Employee salary is low")
	logger.Error("Employee data not found")
}
