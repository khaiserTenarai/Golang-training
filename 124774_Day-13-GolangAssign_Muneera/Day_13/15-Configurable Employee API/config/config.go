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

// Load reads settings from environment variables and .env.
func Load() (*Config, error) {
	// The .env file is optional.
	_ = godotenv.Load(".env")

	appPort := getEnv("APP_PORT", "8080")
	dbPort := getEnv("DB_PORT", "5432")
	appEnv := getEnv("APP_ENV", "development")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	if err := validatePort(appPort); err != nil {
		return nil, fmt.Errorf("invalid APP_PORT: %w", err)
	}

	if err := validatePort(dbPort); err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	// Production must not use disabled database SSL.
	if appEnv == "production" && dbSSLMode == "disable" {
		return nil, fmt.Errorf(
			"DB_SSLMODE must not be disable in production",
		)
	}

	// Employee creation is enabled by default.
	enableCreate := true

	if value := os.Getenv("ENABLE_EMPLOYEE_CREATE"); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid ENABLE_EMPLOYEE_CREATE: %w", err,
			)
		}

		enableCreate = parsed
	}

	cfg := &Config{
		AppEnv:               appEnv,
		AppPort:              appPort,
		DBHost:               getEnv("DB_HOST", "localhost"),
		DBPort:               dbPort,
		DBUser:               getEnv("DB_USER", "postgres"),
		DBName:               getEnv("DB_NAME", "employee_db"),
		DBSSLMode:            dbSSLMode,
		EnableEmployeeCreate: enableCreate,
	}

	return cfg, nil
}

// getEnv returns a value or its default.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

// validatePort checks whether the port is valid.
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