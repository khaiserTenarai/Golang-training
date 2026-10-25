package config

import (
	"fmt"
	"os"
)

// Config stores Department Service settings.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
}

// Load reads settings from environment variables.
func Load() Config {
	return Config{
		DBHost:     getEnv("DEPARTMENT_DB_HOST", "localhost"),
		DBPort:     getEnv("DEPARTMENT_DB_PORT", "5432"),
		DBUser:     getEnv("DEPARTMENT_DB_USER", "postgres"),
		DBPassword: os.Getenv("DEPARTMENT_DB_PASSWORD"),
		DBName:     getEnv("DEPARTMENT_DB_NAME", "department_db"),
		Port:       getEnv("DEPARTMENT_SERVICE_PORT", "8082"),
	}
}

// DSN creates the PostgreSQL connection string.
func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
