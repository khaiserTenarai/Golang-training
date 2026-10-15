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

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// DevelopmentDefaults returns default development settings.
func DevelopmentDefaults() *Config {
	return &Config{
		AppEnv:     "development",
		AppPort:    "8080",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "",
		DBName:     "employee_db",
		DBSSLMode:  "disable",
	}
}

// Load reads environment variables and applies defaults.
func Load() *Config {
	// Load .env for local development.
	// Existing system environment variables take precedence.
	_ = godotenv.Load()

	// Determine the environment.
	env := getEnv("APP_ENV", "development")

	var cfg *Config

	// Select environment-specific defaults.
	switch env {
	case "development":
		cfg = DevelopmentDefaults()
	default:
		cfg = &Config{
			AppEnv:     env,
			AppPort:    "8080",
			DBHost:     "localhost",
			DBPort:     "5432",
			DBUser:     "postgres",
			DBPassword: "",
			DBName:     "employee_db",
			DBSSLMode:  "disable",
		}
	}

	// Override defaults with environment variables.
	cfg.AppEnv = getEnv("APP_ENV", cfg.AppEnv)
	cfg.AppPort = getEnv("APP_PORT", cfg.AppPort)
	cfg.DBHost = getEnv("DB_HOST", cfg.DBHost)
	cfg.DBPort = getEnv("DB_PORT", cfg.DBPort)
	cfg.DBUser = getEnv("DB_USER", cfg.DBUser)
	cfg.DBPassword = getEnv("DB_PASSWORD", cfg.DBPassword)
	cfg.DBName = getEnv("DB_NAME", cfg.DBName)
	cfg.DBSSLMode = getEnv("DB_SSLMODE", cfg.DBSSLMode)

	return cfg
}

// getEnv returns the environment value or a fallback.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

// Validate checks required settings and port numbers.
func (c *Config) Validate() error {
	// Validate required database settings.
	if c.DBHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	if c.DBUser == "" {
		return fmt.Errorf("DB_USER is required")
	}

	if c.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	// Validate application port.
	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	// Validate database port.
	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	// Require a database password.
	if c.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	return nil
}
