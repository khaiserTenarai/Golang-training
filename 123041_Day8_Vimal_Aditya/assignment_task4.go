package main

import (
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// log structured data using key value parameters
	logger.Info("application started", "environment", "production", "version", "1.0.0")

	logger.Info("processing employee", "emp_id", 101, "name", "Vimal", "salary", 500000.0)

	logger.Warn("high salary bonus calculated", "emp_id", 101, "bonus", 50000.0)

	logger.Error("failed to update record", "emp_id", -1, "reason", "invalid ID")
}