package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv               string
	AppPort              string
	DBHost               string
	DBPort               string
	DBUser               string
	DBName               string
	DBSSLMode            string
	EnableEmployeeCreate bool
}

// Load reads configuration from .env and environment variables.
func Load() (*Config, error) {
	// The .env file is optional.
	_ = godotenv.Load(".env")

	appPort := getEnv("APP_PORT", "8080")
	dbPort := getEnv("DB_PORT", "5432")
	appEnv := getEnv("APP_ENV", "development")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	// Validate the application port.
	if err := validatePort(appPort); err != nil {
		return nil, fmt.Errorf("invalid APP_PORT: %w", err)
	}

	// Validate the database port.
	if err := validatePort(dbPort); err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	// Require SSL when running in production.
	if appEnv == "production" && dbSSLMode == "disable" {
		return nil, fmt.Errorf("DB_SSLMODE must enable SSL in production")
	}

	// Employee creation is enabled by default.
	enableEmployeeCreate := true

	// Parse the feature flag if it is configured.
	if value := os.Getenv("ENABLE_EMPLOYEE_CREATE"); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid ENABLE_EMPLOYEE_CREATE value: %w", err,
			)
		}

		enableEmployeeCreate = parsed
	}

	cfg := &Config{
		AppEnv:               appEnv,
		AppPort:              appPort,
		DBHost:               getEnv("DB_HOST", "localhost"),
		DBPort:               dbPort,
		DBUser:               getEnv("DB_USER", "postgres"),
		DBName:               getEnv("DB_NAME", "employee_db"),
		DBSSLMode:            dbSSLMode,
		EnableEmployeeCreate: enableEmployeeCreate,
	}

	return cfg, nil
}

// getEnv returns an environment variable or its default value.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

// validatePort checks whether a port is between 1 and 65535.
func validatePort(value string) error {
	port, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("port must be a number")
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}

	return nil
}