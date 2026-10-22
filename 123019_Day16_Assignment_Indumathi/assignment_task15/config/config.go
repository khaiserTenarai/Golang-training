// Package config loads application configuration from environment variables.
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config contains environment settings for PostgreSQL and the server.
type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	ServerPort string
	LogLevel   string
}

// Load reads .env and system environment variables.
func Load() Config {
	_ = godotenv.Load()
	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		ServerPort: os.Getenv("SERVER_PORT"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
	}
}