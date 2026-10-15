package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores normal application and database settings.
type Config struct {
	AppEnv  string
	AppPort string

	DBHost    string
	DBPort    string
	DBUser    string
	DBName    string
	DBSSLMode string
}

// Load reads configuration from the .env file and environment variables.
func Load() (*Config, error) {
	// Load .env values without overwriting existing environment variables.
	// The file is optional when settings are supplied by the environment.
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	// Read settings and use defaults when optional values are missing.
	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBName:    getEnv("DB_NAME", "employee_db"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}

	// Validate settings before starting the application.
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// getEnv returns an environment variable or its fallback value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

// Validate checks whether the configuration values are valid.
func (c *Config) Validate() error {
	// The application port must be between 1 and 65535.
	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	// The database port must also be valid.
	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	// These database settings are required.
	if c.DBHost == "" || c.DBUser == "" || c.DBName == "" {
		return fmt.Errorf("DB_HOST, DB_USER, and DB_NAME are required")
	}

	// Production must use a secure PostgreSQL SSL mode.
	if c.AppEnv == "production" &&
		c.DBSSLMode != "require" &&
		c.DBSSLMode != "verify-ca" &&
		c.DBSSLMode != "verify-full" {
		return fmt.Errorf("production requires a secure DB_SSLMODE")
	}

	return nil
}