package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	AppEnv               string
	Port                 string
	ReadTimeout          time.Duration
	ReadHeaderTimeout    time.Duration
	WriteTimeout         time.Duration
	IdleTimeout          time.Duration
	ShutdownTimeout      time.Duration
	TokenTTL             time.Duration
	EnableEmployeeDelete bool
	LogLevel             string
}

type Secrets struct {
	JWTSecret string
}

func Load() (Config, Secrets, error) {
	env := strings.ToLower(strings.TrimSpace(getEnv("APP_ENV", "development")))
	cfg := Config{
		AppEnv:               env,
		Port:                 getEnv("PORT", "8080"),
		ReadTimeout:          10 * time.Second,
		ReadHeaderTimeout:    5 * time.Second,
		WriteTimeout:         15 * time.Second,
		IdleTimeout:          60 * time.Second,
		ShutdownTimeout:      10 * time.Second,
		TokenTTL:             8 * time.Hour,
		EnableEmployeeDelete: true,
		LogLevel:             getEnv("LOG_LEVEL", "info"),
	}
	switch env {
	case "development":
		cfg.LogLevel = getEnv("LOG_LEVEL", "debug")
	case "test":
		cfg.Port = getEnv("PORT", "8081")
		cfg.ReadTimeout = 3 * time.Second
		cfg.ReadHeaderTimeout = 2 * time.Second
		cfg.WriteTimeout = 3 * time.Second
		cfg.IdleTimeout = 5 * time.Second
		cfg.ShutdownTimeout = 2 * time.Second
		cfg.TokenTTL = 15 * time.Minute
		cfg.EnableEmployeeDelete = false
		cfg.LogLevel = getEnv("LOG_LEVEL", "error")
	case "production":
		cfg.EnableEmployeeDelete = parseBoolEnv("FEATURE_EMPLOYEE_DELETE", false)
		cfg.LogLevel = getEnv("LOG_LEVEL", "info")
	default:
		return Config{}, Secrets{}, fmt.Errorf("APP_ENV must be development, test, or production")
	}
	if value := os.Getenv("READ_TIMEOUT"); value != "" {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return Config{}, Secrets{}, errors.New("READ_TIMEOUT must be a valid duration")
		}
		cfg.ReadTimeout = duration
	}
	if value := os.Getenv("SHUTDOWN_TIMEOUT"); value != "" {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return Config{}, Secrets{}, errors.New("SHUTDOWN_TIMEOUT must be a valid duration")
		}
		cfg.ShutdownTimeout = duration
	}
	if value := os.Getenv("TOKEN_TTL"); value != "" {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return Config{}, Secrets{}, errors.New("TOKEN_TTL must be a valid duration")
		}
		cfg.TokenTTL = duration
	}
	secrets := Secrets{JWTSecret: os.Getenv("JWT_SECRET")}
	if env == "development" && secrets.JWTSecret == "" {
		secrets.JWTSecret = "development-only-change-me"
	}
	if err := Validate(cfg, secrets); err != nil {
		return Config{}, Secrets{}, err
	}
	return cfg, secrets, nil
}

func Validate(cfg Config, secrets Secrets) error {
	if cfg.Port == "" {
		return errors.New("PORT cannot be empty")
	}
	port, err := strconv.Atoi(cfg.Port)
	if err != nil || port < 1 || port > 65535 {
		return errors.New("PORT must be an integer between 1 and 65535")
	}
	if cfg.ReadTimeout <= 0 || cfg.ReadHeaderTimeout <= 0 || cfg.WriteTimeout <= 0 || cfg.IdleTimeout <= 0 || cfg.ShutdownTimeout <= 0 || cfg.TokenTTL <= 0 {
		return errors.New("timeouts and token TTL must be positive")
	}
	if cfg.LogLevel != "debug" && cfg.LogLevel != "info" && cfg.LogLevel != "warn" && cfg.LogLevel != "error" {
		return errors.New("LOG_LEVEL must be debug, info, warn, or error")
	}
	if cfg.AppEnv == "production" && len(secrets.JWTSecret) < 32 {
		return errors.New("JWT_SECRET must contain at least 32 characters in production")
	}
	if secrets.JWTSecret == "" {
		return errors.New("JWT_SECRET is required outside development")
	}
	return nil
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func parseBoolEnv(key string, fallback bool) bool {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}
