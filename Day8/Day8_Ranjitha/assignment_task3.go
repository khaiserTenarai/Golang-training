package main

import (
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Info Level
	logger.Info("Application started")

	// Warn Level
	logger.Warn("Salary is zero, please check record")

	// Error Level
	logger.Error("Invalid employee ID provided")
}