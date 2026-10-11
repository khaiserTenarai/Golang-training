package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppEnv         string
	Port           int
	ShutdownTimeout time.Duration
	SecretKey      string
	EnableNewFeature bool // Configuration-based feature flag
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		Port:            getEnvAsInt("PORT", 4444),
		ShutdownTimeout: time.Duration(getEnvAsInt("SHUTDOWN_TIMEOUT_SECONDS", 5)) * time.Second,
		SecretKey:       os.Getenv("SECRET_KEY"),
		EnableNewFeature: getEnvAsBool("ENABLE_NEW_FEATURE", false),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.AppEnv == "production" && c.SecretKey == "" {
		return errors.New("SECRET_KEY environment variable is required in production environment")
	}
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port number: %d", c.Port)
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valStr, exists := os.LookupEnv(key); exists {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if valStr, exists := os.LookupEnv(key); exists {
		if val, err := strconv.ParseBool(valStr); err == nil {
			return val
		}
	}
	return defaultValue
}