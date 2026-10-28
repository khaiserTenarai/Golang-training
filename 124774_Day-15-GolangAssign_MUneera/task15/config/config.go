package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBName            string
	DBUser            string
	DBPassword        string
	ServerPort        string
	LogLevel          string
	JWTSecret         string
	JWTExpiryMinutes  string
	AdminUsername     string
	AdminPasswordHash string
	UserUsername      string
	UserPasswordHash  string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBName:            getEnv("DB_NAME", "employee_db"),
		DBUser:            getEnv("DB_USER", "postgres"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		ServerPort:        getEnv("SERVER_PORT", "8080"),
		LogLevel:          getEnv("LOG_LEVEL", "debug"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTExpiryMinutes:  getEnv("JWT_EXPIRY_MINUTES", "60"),
		AdminUsername:     getEnv("ADMIN_USERNAME", "admin"),
		AdminPasswordHash: os.Getenv("ADMIN_PASSWORD_HASH"),
		UserUsername:      getEnv("USER_USERNAME", "user"),
		UserPasswordHash:  os.Getenv("USER_PASSWORD_HASH"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
