package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores normal application settings.
// Passwords and other secrets are kept in a separate struct.
type Config struct {
	AppEnv  string
	AppPort string

	DBHost    string
	DBPort    string
	DBUser    string
	DBName    string
	DBSSLMode string

	// EnableEmployeeCreate controls the employee creation endpoint.
	EnableEmployeeCreate bool
}

// Load reads and validates the application configuration.
func Load() (*Config, error) {
	// Load .env for local development.
	// In production, environment variables can be supplied by the deployment.
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	// Read configuration and apply defaults.
	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBName:    getEnv("DB_NAME", "employee_db"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}

	// Read the feature flag as a boolean.
	enableCreate, err := strconv.ParseBool(
		getEnv("ENABLE_EMPLOYEE_CREATE", "true"),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"ENABLE_EMPLOYEE_CREATE must be true or false",
		)
	}
	cfg.EnableEmployeeCreate = enableCreate

	// Reject invalid configuration before starting the application.
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// getEnv returns an environment variable or a default value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

// Validate checks whether the configuration is valid.
func (c *Config) Validate() error {
	// Validate the HTTP port.
	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	// Validate the PostgreSQL port.
	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	// Require database connection details.
	if c.DBHost == "" || c.DBUser == "" || c.DBName == "" {
		return fmt.Errorf("DB_HOST, DB_USER, and DB_NAME are required")
	}

	// Require a secure SSL mode in production.
	if c.AppEnv == "production" &&
		c.DBSSLMode != "require" &&
		c.DBSSLMode != "verify-ca" &&
		c.DBSSLMode != "verify-full" {
		return fmt.Errorf("production requires a secure DB_SSLMODE")
	}

	return nil
}