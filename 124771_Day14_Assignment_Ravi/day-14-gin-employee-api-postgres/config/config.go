package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func Load() Config {
	// .env is optional if environment variables are already configured.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; reading environment variables")
	}

	cfg := Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME", "employee_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	if cfg.DBPassword == "" {
		log.Println("Warning: DB_PASSWORD is empty; check your .env file")
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
