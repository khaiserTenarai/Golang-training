// Package config contains application configuration loading logic.
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config contains all application configuration values.
type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	ServerPort string
	LogLevel   string
	JWTSecret  string
}

// Load loads configuration from the .env file and environment variables.
func Load() Config {
	_ = godotenv.Load()

	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "employee_db"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "pgadmin"),
		ServerPort: getEnv("SERVER_PORT", "10005"),
		LogLevel:   getEnv("LOG_LEVEL", "INFO"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey123"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}