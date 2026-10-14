package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds non-secret settings and secret settings separately.
// Tasks 2, 3: configuration struct and validation.
// Task 8: Secrets are grouped separately from ordinary application settings.
type Config struct {
	AppEnv   string
	AppPort  string
	LogLevel string

	FeatureFlags FeatureFlags
	Secrets      Secrets
	Database     DatabaseConfig
}

type FeatureFlags struct {
	CreateEmployee bool
}

// Secrets contains credentials that must not be logged.
// Task 8: database credentials are kept separate from ordinary config.
type Secrets struct {
	DBUser     string
	DBPassword string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	SSLMode  string
	MaxConns int32
}

// Load reads .env if present, then reads process environment variables.
// Existing process environment values override .env values.
// Task 1: read configuration from environment variables.
// Task 4: provide defaults.
// Tasks 5–7: support development, test, and production modes.
func Load() (Config, error) {
	_ = loadDotEnv(".env") // Optional local development file; OS environment wins.

	env := envOrDefault("APP_ENV", "development")
	cfg := Config{
		AppEnv:   env,
		AppPort:  envOrDefault("APP_PORT", "8080"),
		LogLevel: envOrDefault("LOG_LEVEL", defaultLogLevel(env)),
		FeatureFlags: FeatureFlags{
			CreateEmployee: envBoolOrDefault("FEATURE_CREATE_EMPLOYEE", true),
		},
		Secrets: Secrets{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
		},
		Database: DatabaseConfig{
			Host:     envOrDefault("DB_HOST", "localhost"),
			Port:     envOrDefault("DB_PORT", "5432"),
			Name:     envOrDefault("DB_NAME", "employee_api"),
			SSLMode:  envOrDefault("DB_SSLMODE", defaultSSLMode(env)),
			MaxConns: int32(envIntOrDefault("DB_MAX_CONNS", 5)),
		},
	}

	// Task 3: validate configuration before starting the app.
	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) Validate() error {
	if c.AppEnv != "development" && c.AppEnv != "test" && c.AppEnv != "production" {
		return fmt.Errorf("APP_ENV must be development, test, or production")
	}
	port, err := strconv.Atoi(c.AppPort)
	if err != nil || port < 1 || port > 65535 {
		return errors.New("APP_PORT must be a valid port between 1 and 65535")
	}
	dbPort, err := strconv.Atoi(c.Database.Port)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return errors.New("DB_PORT must be a valid port between 1 and 65535")
	}
	if strings.TrimSpace(c.Database.Host) == "" || strings.TrimSpace(c.Database.Name) == "" {
		return errors.New("DB_HOST and DB_NAME are required")
	}
	if c.Database.MaxConns < 1 {
		return errors.New("DB_MAX_CONNS must be at least 1")
	}
	if strings.TrimSpace(c.Secrets.DBUser) == "" {
		return errors.New("DB_USER is required; set it in your environment or .env")
	}
	// Task 7: production must never fall back to a placeholder password.
	if c.AppEnv == "production" && strings.TrimSpace(c.Secrets.DBPassword) == "" {
		return errors.New("DB_PASSWORD is required in production")
	}
	if c.AppEnv != "production" && strings.TrimSpace(c.Secrets.DBPassword) == "" {
		return errors.New("DB_PASSWORD is required; set it in your environment or .env.example copy")
	}
	if c.LogLevel != "debug" && c.LogLevel != "info" && c.LogLevel != "warn" && c.LogLevel != "error" {
		return errors.New("LOG_LEVEL must be debug, info, warn, or error")
	}
	return nil
}

func defaultLogLevel(env string) string {
	if env == "production" {
		return "info"
	}
	return "debug"
}

func defaultSSLMode(env string) string {
	if env == "production" {
		return "require"
	}
	return "disable"
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func envBoolOrDefault(key string, fallback bool) bool {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback // Validate feature flag format separately below in production code if needed.
	}
	return parsed
}

func envIntOrDefault(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

// loadDotEnv is a small .env reader to keep this project dependency-light.
// It does not override values already set in the process environment.
func loadDotEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key == "" {
			continue
		}
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}
	return scanner.Err()
}
