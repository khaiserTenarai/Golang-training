package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		       Level: slog.LevelInfo,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("User logged in",
		"user_id", 42,
		"ip_address", "192.168.1.100",
		"role", "admin",
	)

	slog.Error("Payment processing failed",
		"transaction_id", "txn_89321",
		"error_code", 502,
		"retries", 3,
	)
}