// Tasks 1, 2, 3, 4: Configuration
// Task 1 — Read configuration from environment variables
// Task 2 — Create a configuration struct
// Task 3 — Validate configuration
// Task 4 — Implement configuration defaults

package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task 2: Configuration struct
type Config struct {
	// Server
	ServerPort         string
	ServerHost         string
	Environment        string // development, test, production
	ReadTimeoutSecs    int
	WriteTimeoutSecs   int
	ShutdownTimeoutSecs int

	// Logging
	LogLevel  string // debug, info, warn, error
	LogFormat string // text, json

	// Feature Flags (Task 14)
	EnablePagination bool
	EnableFiltering  bool
	EnableSorting    bool
	EnableAuth       bool
	MaxPageSize      int

	// Secrets (Task 8 — loaded separately)
	Secrets *SecretsConfig
}

// Task 1: Read from environment variables with Task 4 defaults
func Load() *Config {
	cfg := &Config{
		ServerPort:          getEnv("SERVER_PORT", "8080"),
		ServerHost:          getEnv("SERVER_HOST", "0.0.0.0"),
		Environment:         getEnv("APP_ENV", "development"),
		ReadTimeoutSecs:     getEnvInt("READ_TIMEOUT_SECS", 10),
		WriteTimeoutSecs:    getEnvInt("WRITE_TIMEOUT_SECS", 10),
		ShutdownTimeoutSecs: getEnvInt("SHUTDOWN_TIMEOUT_SECS", 5),
		LogLevel:            getEnv("LOG_LEVEL", "debug"),
		LogFormat:           getEnv("LOG_FORMAT", "text"),
		EnablePagination:    getEnvBool("FEATURE_PAGINATION", true),
		EnableFiltering:     getEnvBool("FEATURE_FILTERING", true),
		EnableSorting:       getEnvBool("FEATURE_SORTING", true),
		EnableAuth:          getEnvBool("FEATURE_AUTH", false),
		MaxPageSize:         getEnvInt("MAX_PAGE_SIZE", 50),
		Secrets:             LoadSecrets(),
	}

	// Apply environment-specific overrides (Tasks 5, 6, 7)
	applyEnvironmentOverrides(cfg)

	return cfg
}

// Task 3: Validate configuration
func (c *Config) Validate() error {
	var errs []string

	if c.ServerPort == "" {
		errs = append(errs, "SERVER_PORT is required")
	}
	port, err := strconv.Atoi(c.ServerPort)
	if err != nil || port < 1 || port > 65535 {
		errs = append(errs, "SERVER_PORT must be 1-65535")
	}

	validEnvs := map[string]bool{"development": true, "test": true, "production": true}
	if !validEnvs[c.Environment] {
		errs = append(errs, "APP_ENV must be development, test, or production")
	}

	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLevels[c.LogLevel] {
		errs = append(errs, "LOG_LEVEL must be debug, info, warn, or error")
	}

	if c.ReadTimeoutSecs <= 0 {
		errs = append(errs, "READ_TIMEOUT_SECS must be positive")
	}
	if c.WriteTimeoutSecs <= 0 {
		errs = append(errs, "WRITE_TIMEOUT_SECS must be positive")
	}
	if c.MaxPageSize <= 0 || c.MaxPageSize > 1000 {
		errs = append(errs, "MAX_PAGE_SIZE must be 1-1000")
	}

	if c.Secrets != nil {
		if c.EnableAuth && c.Secrets.AdminToken == "" {
			errs = append(errs, "ADMIN_TOKEN required when auth is enabled")
		}
	}

	if len(errs) > 0 {
		return errors.New("config validation failed: " + strings.Join(errs, "; "))
	}
	return nil
}

func (c *Config) Print() {
	fmt.Println("  ┌─── Configuration ───────────────────────")
	fmt.Printf("  │ Environment:  %s\n", c.Environment)
	fmt.Printf("  │ Server:       %s:%s\n", c.ServerHost, c.ServerPort)
	fmt.Printf("  │ Log Level:    %s\n", c.LogLevel)
	fmt.Printf("  │ Log Format:   %s\n", c.LogFormat)
	fmt.Printf("  │ Timeouts:     read=%ds, write=%ds, shutdown=%ds\n",
		c.ReadTimeoutSecs, c.WriteTimeoutSecs, c.ShutdownTimeoutSecs)
	fmt.Printf("  │ Features:     pagination=%v, filtering=%v, sorting=%v, auth=%v\n",
		c.EnablePagination, c.EnableFiltering, c.EnableSorting, c.EnableAuth)
	fmt.Printf("  │ Max Page:     %d\n", c.MaxPageSize)
	fmt.Printf("  │ Secrets:      admin_token=%s, api_key=%s\n",
		maskSecret(c.Secrets.AdminToken), maskSecret(c.Secrets.APIKey))
	fmt.Println("  └──────────────────────────────────────────")
}

func maskSecret(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:4] + "****"
}

// ---- Env helpers (Task 1, 4) ----

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defaultVal
}
