package config

import (
	"log"
	"os"

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

// Load reads environment variables and returns the configuration.
func Load() *Config {
	// Load .env for local development.
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found; using system environment")
	}

	// Create and return the configuration object.
	return &Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "employee_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// getEnv returns the environment value or a fallback.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
