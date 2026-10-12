package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
	UserToken  string
	AdminToken string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		Port:       get("SERVER_PORT", "8080"),
		DBHost:     get("DB_HOST", "localhost"),
		DBPort:     get("DB_PORT", "5432"),
		DBName:     get("DB_NAME", "go_training"),
		DBUser:     get("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBSSLMode:  get("DB_SSLMODE", "disable"),
		UserToken:  os.Getenv("USER_TOKEN"),
		AdminToken: os.Getenv("ADMIN_TOKEN"),
	}
}

func get(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
