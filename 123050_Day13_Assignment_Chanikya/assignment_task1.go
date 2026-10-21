package main

import (
	"fmt"
	"os"
)

// Task 1: Basic Layered Architecture Setup & Env Configuration Reading
// Demonstrates direct environment variable reading at the entry level.
func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Fallback fallback
	}
	fmt.Printf("[Task 1] Read configuration from environment variables. DB_HOST: %s\n", dbHost)
}
