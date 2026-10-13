package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv                 string
	Port                   string
	DatabaseURL            string
	LogLevel               string
	EmployeeExportEnabled  bool
	ShutdownTimeoutSeconds int
}

func loadConfig() (Config, error) {
	// .env is optional. Real OS environment variables take precedence.
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found; reading configuration from system environment")
	}

	cfg := Config{
		AppEnv:                 getEnv("APP_ENV", "development"),
		Port:                   getEnv("PORT", "8080"),
		DatabaseURL:            strings.TrimSpace(os.Getenv("DATABASE_URL")),
		LogLevel:               strings.ToLower(getEnv("LOG_LEVEL", defaultLogLevel())),
		EmployeeExportEnabled:  getBool("FEATURE_EMPLOYEE_EXPORT", false),
		ShutdownTimeoutSeconds: getInt("SHUTDOWN_TIMEOUT_SECONDS", 10),
	}
	if err := cfg.validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) validate() error {
	if c.AppEnv != "development" && c.AppEnv != "test" && c.AppEnv != "production" {
		return fmt.Errorf("APP_ENV must be development, test, or production")
	}
	port, err := strconv.Atoi(c.Port)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("PORT must be a number from 1 to 65535")
	}
	if c.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	switch c.LogLevel {
	case "debug", "info", "warn", "error":
	default:
		return fmt.Errorf("LOG_LEVEL must be debug, info, warn, or error")
	}
	if c.ShutdownTimeoutSeconds < 1 || c.ShutdownTimeoutSeconds > 120 {
		return fmt.Errorf("SHUTDOWN_TIMEOUT_SECONDS must be from 1 to 120")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func getInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return -1 // validation reports the bad value
	}
	return n
}

func getBool(key string, fallback bool) bool {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return b
}

func defaultLogLevel() string {
	if strings.EqualFold(os.Getenv("APP_ENV"), "production") {
		return "info"
	}
	return "debug"
}

// Environment presets for the assignment; environment variables can override them.
func developmentDefaults() Config {
	return Config{AppEnv: "development", Port: "8080", LogLevel: "debug", ShutdownTimeoutSeconds: 10}
}
func testDefaults() Config {
	return Config{AppEnv: "test", Port: "8081", LogLevel: "error", ShutdownTimeoutSeconds: 3}
}
func productionDefaults() Config {
	return Config{AppEnv: "production", Port: "8080", LogLevel: "info", ShutdownTimeoutSeconds: 20}
}
