package config

import (
	"os"

	"github.com/joho/godotenv"
)

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

func Load() Config {
	_ = godotenv.Load()

	return Config{
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBName:           getEnv("DB_NAME", "employee_db"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       os.Getenv("DB_PASSWORD"),
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		JWTExpiryMinutes: getEnv("JWT_EXPIRY_MINUTES", "60"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
