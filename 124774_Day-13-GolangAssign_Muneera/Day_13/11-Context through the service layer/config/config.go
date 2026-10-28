package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores application and database settings.
type Config struct {
	AppEnv  string
	AppPort string

	DBHost    string
	DBPort    string
	DBUser    string
	DBName    string
	DBSSLMode string
}

// Load reads configuration from .env and process environment variables.
func Load() (*Config, error) {
	// Load .env without overwriting existing environment variables.
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("load .env file: %w", err)
	}

	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBName:    getEnv("DB_NAME", "employee_db"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// getEnv returns an environment variable or a fallback value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

// Validate checks application and database configuration.
func (c *Config) Validate() error {
	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	if c.DBHost == "" || c.DBUser == "" || c.DBName == "" {
		return fmt.Errorf("DB_HOST, DB_USER, and DB_NAME are required")
	}

	if c.AppEnv == "production" &&
		c.DBSSLMode != "require" &&
		c.DBSSLMode != "verify-ca" &&
		c.DBSSLMode != "verify-full" {
		return fmt.Errorf("production requires a secure DB_SSLMODE")
	}

	return nil
}