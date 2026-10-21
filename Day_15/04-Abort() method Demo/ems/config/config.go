package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// Load reads the application configuration from .env.
func Load() (Config, error) {
	// Load .env if it exists; environment variables can also be supplied externally.
	_ = godotenv.Load()

	cfg := Config{
		DBHost:     get("DB_HOST", "localhost"),
		DBPort:     get("DB_PORT", "5432"),
		DBUser:     get("DB_USER", "postgres"),
		DBPassword: get("DB_PASSWORD", ""),
		DBName:     get("DB_NAME", "employee_db"),
		ServerPort: get("SERVER_PORT", "8081"),
	}

	if cfg.DBPassword == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is required")
	}

	return cfg, nil
}

// get returns an environment value or its default.
func get(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

// DSN builds the PostgreSQL connection string.
func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBPort,
	)
}