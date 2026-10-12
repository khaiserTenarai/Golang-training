// Package config loads application configuration.
package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Config contains environment settings.
type Config struct {
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	ServerPort       string
	LogLevel         string
	JWTSecret        string
	JWTExpiryMinutes string
}

// Load reads .env and process environment variables.
func Load() Config {
	_ = godotenv.Load()
	return Config{
		DBHost: os.Getenv("DB_HOST"), DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"), DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"), ServerPort: os.Getenv("SERVER_PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"), JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiryMinutes: os.Getenv("JWT_EXPIRY_MINUTES"),
	}
}
