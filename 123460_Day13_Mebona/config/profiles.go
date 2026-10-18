package config

import "time"

// Placeholder keys that are accepted in development/test but REJECTED in production.
const (
	DevAPIKey  = "dev-api-key-change-me"
	TestAPIKey = "test-api-key-not-secret"
)

// defaultsFor returns the built-in configuration of one environment.
// This is where "development", "test" and "production" configurations differ.
// Anything set in a .env file or the real environment overrides these values.
func defaultsFor(env string) Config {
	cfg := Config{
		App: AppConfig{Name: "employee-api", Env: env},
		Server: ServerConfig{
			Host:            "localhost",
			Port:            8080,
			ReadTimeout:     15 * time.Second,
			WriteTimeout:    15 * time.Second,
			IdleTimeout:     60 * time.Second,
			RequestTimeout:  5 * time.Second,
			StartupTimeout:  10 * time.Second,
			ShutdownTimeout: 15 * time.Second,
		},
		Log:        LogConfig{Level: "info", Format: "text"},
		Pagination: PaginationConfig{DefaultPageSize: 10, MaxPageSize: 100},
		Features:   Features{SearchEnabled: true, ExposeSalary: true, SeedData: false},
	}

	switch env {
	case EnvTest:
		cfg.Server.Host = "127.0.0.1"
		cfg.Server.Port = 8081
		cfg.Server.WriteTimeout = 5 * time.Second
		cfg.Server.RequestTimeout = 2 * time.Second
		cfg.Server.ShutdownTimeout = 1 * time.Second
		cfg.Log = LogConfig{Level: "error", Format: "text"}
		cfg.Secrets.APIKey = Secret(TestAPIKey)

	case EnvProduction:
		cfg.Server.Host = "0.0.0.0"
		cfg.Server.ShutdownTimeout = 30 * time.Second
		cfg.Log = LogConfig{Level: "info", Format: "json"}
		cfg.Features.ExposeSalary = false // safer default
		// No default API key: it MUST come from .env.secrets or the real environment.

	default: // development
		cfg.Log = LogConfig{Level: "debug", Format: "text"}
		cfg.Features.SeedData = true
		cfg.Secrets.APIKey = Secret(DevAPIKey)
	}
	return cfg
}
