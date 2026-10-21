package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	// .env is optional in production if environment variables are injected.
	// Other environment variables take precedence over .env values.
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return Config{}, fmt.Errorf("load .env: %w", err)
	}

	cfg := Config{
		DBHost:           env("DB_HOST", "localhost"),
		DBPort:           env("DB_PORT", "5432"),
		DBName:           env("DB_NAME", "employee_db"),
		DBUser:           env("DB_USER", "postgres"),
		DBPassword:       os.Getenv("DB_PASSWORD"),
		DBSSLMode:        env("DB_SSLMODE", "disable"),
		ServerPort:       env("SERVER_PORT", "10005"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		AdminUsername:    os.Getenv("ADMIN_USERNAME"),
		AdminPassword:    os.Getenv("ADMIN_PASSWORD"),
		UserUsername:     os.Getenv("USER_USERNAME"),
		UserPassword:     os.Getenv("USER_PASSWORD"),
		JWTExpiryMinutes: 60,
	}

	if v := os.Getenv("JWT_EXPIRY_MINUTES"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n <= 0 {
			return Config{}, fmt.Errorf("JWT_EXPIRY_MINUTES must be positive")
		}
		cfg.JWTExpiryMinutes = n
	}

	required := map[string]string{
		"DB_PASSWORD":    cfg.DBPassword,
		"JWT_SECRET":     cfg.JWTSecret,
		"ADMIN_USERNAME": cfg.AdminUsername,
		"ADMIN_PASSWORD": cfg.AdminPassword,
		"USER_USERNAME":  cfg.UserUsername,
		"USER_PASSWORD":  cfg.UserPassword,
	}

	for key, value := range required {
		if strings.TrimSpace(value) == "" {
			return Config{}, fmt.Errorf("%s is required", key)
		}
	}

	if len(cfg.JWTSecret) < 32 {
		return Config{}, fmt.Errorf("JWT_SECRET must be at least 32 bytes")
	}
	if cfg.AdminUsername == cfg.UserUsername {
		return Config{}, fmt.Errorf("admin and user usernames must differ")
	}

	return cfg, nil
}

func env(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}
