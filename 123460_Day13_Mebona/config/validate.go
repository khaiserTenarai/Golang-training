package config

import (
	"errors"
	"fmt"
	"time"
)

// Validate checks the finished configuration and reports EVERY problem, not just the first.
// The application refuses to start if this fails (fail fast).
func (c *Config) Validate() error {
	var errs []error
	add := func(format string, args ...any) { errs = append(errs, fmt.Errorf(format, args...)) }

	switch c.App.Env {
	case EnvDevelopment, EnvTest, EnvProduction:
	default:
		add("APP_ENV must be development, test or production (got %q)", c.App.Env)
	}
	if c.App.Name == "" {
		add("APP_NAME must not be empty")
	}

	if c.Server.Host == "" {
		add("SERVER_HOST must not be empty")
	}
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		add("SERVER_PORT must be between 1 and 65535 (got %d)", c.Server.Port)
	}
	positive := func(name string, d time.Duration) {
		if d <= 0 {
			add("%s must be greater than 0 (got %s)", name, d)
		}
	}
	positive("SERVER_READ_TIMEOUT", c.Server.ReadTimeout)
	positive("SERVER_WRITE_TIMEOUT", c.Server.WriteTimeout)
	positive("SERVER_IDLE_TIMEOUT", c.Server.IdleTimeout)
	positive("REQUEST_TIMEOUT", c.Server.RequestTimeout)
	positive("STARTUP_TIMEOUT", c.Server.StartupTimeout)
	positive("SHUTDOWN_TIMEOUT", c.Server.ShutdownTimeout)
	if c.Server.RequestTimeout > 0 && c.Server.WriteTimeout > 0 && c.Server.RequestTimeout >= c.Server.WriteTimeout {
		add("REQUEST_TIMEOUT (%s) must be shorter than SERVER_WRITE_TIMEOUT (%s)", c.Server.RequestTimeout, c.Server.WriteTimeout)
	}

	switch c.Log.Level {
	case "debug", "info", "warn", "error":
	default:
		add("LOG_LEVEL must be debug, info, warn or error (got %q)", c.Log.Level)
	}
	switch c.Log.Format {
	case "text", "json":
	default:
		add("LOG_FORMAT must be text or json (got %q)", c.Log.Format)
	}

	if c.Pagination.DefaultPageSize < 1 {
		add("PAGE_SIZE_DEFAULT must be at least 1")
	}
	if c.Pagination.MaxPageSize < c.Pagination.DefaultPageSize || c.Pagination.MaxPageSize > 1000 {
		add("PAGE_SIZE_MAX must be between PAGE_SIZE_DEFAULT and 1000")
	}

	// secrets
	key := c.Secrets.APIKey.Reveal()
	if key == "" {
		add("API_KEY is required (put it in .env.secrets or in the real environment)")
	}

	// stricter rules for production
	if c.IsProduction() {
		if key != "" && len(key) < 16 {
			add("API_KEY must be at least 16 characters in production")
		}
		if key == DevAPIKey || key == TestAPIKey {
			add("API_KEY must not be the development/test placeholder in production")
		}
		if c.Log.Level == "debug" {
			add("LOG_LEVEL=debug is not allowed in production")
		}
		if c.Features.SeedData {
			add("FEATURE_SEED_DATA must be false in production")
		}
	}

	return errors.Join(errs...)
}
