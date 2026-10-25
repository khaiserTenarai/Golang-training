package config

import (
	"fmt"
	"os"
)

// Config stores authentication service settings.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	JWTSecret  string
}

func Load() Config {
	return Config{
		DBHost:     getEnv("AUTH_DB_HOST", "localhost"),
		DBPort:     getEnv("AUTH_DB_PORT", "5432"),
		DBUser:     getEnv("AUTH_DB_USER", "postgres"),
		DBPassword: os.Getenv("AUTH_DB_PASSWORD"),
		DBName:     getEnv("AUTH_DB_NAME", "auth_db"),
		Port:       getEnv("AUTH_SERVICE_PORT", "8083"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

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
