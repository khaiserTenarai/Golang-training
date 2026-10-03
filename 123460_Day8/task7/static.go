package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		slog.Error("Failed to open file", "err", err)
		return
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := slog.New(slog.NewJSONHandler(mw, nil))

	// Fixed: Use %s for string userName, and include userID so it's not unused
	userID := 101
	userName := "admin" 
	fmt.Printf("User ID %d (%s) logged in\n", userID, userName) 

	status := "pending"
	status = "completed"

	// Removed premature 'return' so this line actually executes
	logger.Info("App finished", "status", status)
}