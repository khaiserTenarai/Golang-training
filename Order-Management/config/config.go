package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Load reads environment variables from the .env file.
// Existing environment variables take precedence over values in .env.
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using environment variables")
	}
}

// Get returns an environment variable or a fallback value.
func Get(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// GetDuration returns a duration from an environment variable.
func GetDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Invalid duration for %s; using default", key)
		return fallback
	}

	return duration
}
