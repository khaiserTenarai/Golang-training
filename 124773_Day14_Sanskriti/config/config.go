package config

import (
	"fmt"
	"net"
	"net/url"
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
	AdminUsername     string
	AdminPasswordHash string
	UserUsername      string
	UserPasswordHash  string
	JWTExpiryMinutes  int
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("could not load .env file: %w", err)
	}

	cfg := Config{
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBName:            getEnv("DB_NAME", "emp_db_soap"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		ServerPort:        getEnv("SERVER_PORT", "10006"),
		LogLevel:          getEnv("LOG_LEVEL", "INFO"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		AdminUsername:     os.Getenv("ADMIN_USERNAME"),
		AdminPasswordHash: os.Getenv("ADMIN_PASSWORD_HASH"),
		UserUsername:      os.Getenv("USER_USERNAME"),
		UserPasswordHash:  os.Getenv("USER_PASSWORD_HASH"),
		JWTExpiryMinutes:  60,
	}

	if cfg.DBUser == "" || cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("DB_USER and JWT_SECRET must be configured")
	}

	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBName == "" {
		return Config{}, fmt.Errorf("database host, port, and name are required")
	}

	if cfg.AdminUsername == "" || cfg.UserUsername == "" {
		return Config{}, fmt.Errorf("admin and user usernames must be configured")
	}

	return cfg, nil
}

func (c Config) DatabaseURL() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.DBUser, c.DBPassword),
		Host:   net.JoinHostPort(c.DBHost, c.DBPort),
		Path:   c.DBName,
	}

	query := u.Query()
	query.Set("sslmode", "disable")
	u.RawQuery = query.Encode()

	return u.String()
}