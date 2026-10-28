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

// Load loads configuration for the selected environment.
func Load() *Config {
	// Load .env without overriding existing system variables.
	_ = godotenv.Load()

	env := getEnv("APP_ENV", "development")

	var cfg *Config

	// Choose defaults based on the environment.
	switch env {
	case "development":
		cfg = DevelopmentDefaults()
	case "test":
		cfg = TestDefaults()
	default:
		cfg = &Config{
			AppEnv:  env,
			AppPort: "8080",
		}
	}

	// Environment variables override defaults.
	cfg.AppEnv = env
	cfg.AppPort = getEnv("APP_PORT", cfg.AppPort)
	cfg.DBHost = getEnv("DB_HOST", cfg.DBHost)
	cfg.DBPort = getEnv("DB_PORT", cfg.DBPort)
	cfg.DBUser = getEnv("DB_USER", cfg.DBUser)
	cfg.DBPassword = getEnv("DB_PASSWORD", cfg.DBPassword)
	cfg.DBName = getEnv("DB_NAME", cfg.DBName)
	cfg.DBSSLMode = getEnv("DB_SSLMODE", cfg.DBSSLMode)

	return cfg
}

// getEnv returns an environment value or its fallback.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Validate checks required settings and ports.
func (c *Config) Validate() error {
	if c.DBHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if c.DBUser == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if c.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if c.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	return nil
}

// DevelopmentDefaults returns development settings.
func DevelopmentDefaults() *Config {
	return &Config{
		AppEnv:     "development",
		AppPort:    "8080",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBName:     "employee_db",
		DBSSLMode:  "disable",
	}
}
