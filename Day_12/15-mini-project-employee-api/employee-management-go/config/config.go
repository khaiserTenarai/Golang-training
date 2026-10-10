// Package config contains application configuration loading logic.
package config

// Import os to read environment variables.
import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
) // Import strconv to convert numeric configuration values.

// Import godotenv to load values from the .env file.

// Config contains application configuration values.
type Config struct {
	DBHost            string
	DBPort            string
	DBName            string
	DBUser            string
	DBPassword        string
	ServerPort        string
	LogLevel          string
	JWTSecret         string
	AdminUsername     string
	AdminPasswordHash string
	UserUsername      string
	UserPasswordHash  string
	JWTExpiryMinutes  int
}

// Load loads configuration from .env and environment variables.
func Load() Config {
	_ = godotenv.Load()

	expiry, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTES"))

	if expiry <= 0 {
		expiry = 60
	}
	return Config{
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBName:            os.Getenv("DB_NAME"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		LogLevel:          os.Getenv("LOG_LEVEL"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		AdminUsername:     os.Getenv("ADMIN_USERNAME"),
		AdminPasswordHash: os.Getenv("ADMIN_PASSWORD_HASH"),
		UserUsername:      os.Getenv("USER_USERNAME"),
		UserPasswordHash:  os.Getenv("USER_PASSWORD_HASH"),
		JWTExpiryMinutes:  expiry,
	}
}
