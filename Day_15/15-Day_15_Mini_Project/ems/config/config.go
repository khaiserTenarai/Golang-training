package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	DBSSLMode        string
	ServerPort       string
	JWTSecret        string
	JWTExpiryMinutes int
	AdminUsername    string
	AdminPassword    string
	UserUsername     string
	UserPassword     string
}

func Load() (Config, error) {
	// Load values from .env.
	// Existing environment variables take precedence.
	_ = godotenv.Load()

	cfg := Config{
		DBHost:           os.Getenv("DB_HOST"),
		DBPort:           os.Getenv("DB_PORT"),
		DBName:           os.Getenv("DB_NAME"),
		DBUser:           os.Getenv("DB_USER"),
		DBPassword:       os.Getenv("DB_PASSWORD"),
		DBSSLMode:        os.Getenv("DB_SSLMODE"),
		ServerPort:       os.Getenv("SERVER_PORT"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		AdminUsername:    os.Getenv("ADMIN_USERNAME"),
		AdminPassword:    os.Getenv("ADMIN_PASSWORD"),
		UserUsername:     os.Getenv("USER_USERNAME"),
		UserPassword:     os.Getenv("USER_PASSWORD"),
		JWTExpiryMinutes: 60,
	}

	// Set simple defaults.
	if cfg.DBHost == "" {
		cfg.DBHost = "localhost"
	}
	if cfg.DBPort == "" {
		cfg.DBPort = "5432"
	}
	if cfg.DBName == "" {
		cfg.DBName = "employee_db"
	}
	if cfg.DBUser == "" {
		cfg.DBUser = "postgres"
	}
	if cfg.DBSSLMode == "" {
		cfg.DBSSLMode = "disable"
	}
	if cfg.ServerPort == "" {
		cfg.ServerPort = "10005"
	}

	if value := os.Getenv("JWT_EXPIRY_MINUTES"); value != "" {
		n, err := strconv.Atoi(value)
		if err != nil || n <= 0 {
			return Config{}, fmt.Errorf(
				"JWT_EXPIRY_MINUTES must be positive",
			)
		}
		cfg.JWTExpiryMinutes = n
	}

	if cfg.DBPassword == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is required")
	}
	if len(cfg.JWTSecret) < 32 {
		return Config{}, fmt.Errorf(
			"JWT_SECRET must contain at least 32 bytes",
		)
	}
	if cfg.AdminUsername == "" || cfg.AdminPassword == "" ||
		cfg.UserUsername == "" || cfg.UserPassword == "" {
		return Config{}, fmt.Errorf("admin and user credentials are required")
	}
	if cfg.AdminUsername == cfg.UserUsername {
		return Config{}, fmt.Errorf("admin and user usernames must differ")
	}

	return cfg, nil
}
