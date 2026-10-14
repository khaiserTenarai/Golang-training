package config

import (
	"os"
	"testing"
)

// Task 6: test configuration can be selected through APP_ENV=test.
// This unit test checks config validation without requiring a live database.
func TestConfigValidationRejectsInvalidPort(t *testing.T) {
	cfg := Config{
		AppEnv:   "test",
		AppPort:  "70000",
		LogLevel: "info",
		Secrets:  Secrets{DBUser: "test_user", DBPassword: "test_password"},
		Database: DatabaseConfig{
			Host: "localhost", Port: "5432", Name: "employee_api_test",
			SSLMode: "disable", MaxConns: 2,
		},
	}
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected invalid APP_PORT to fail validation")
	}
}

// Task 4: ensure empty values use documented defaults.
func TestDefaults(t *testing.T) {
	t.Setenv("APP_ENV", "development")
	t.Setenv("APP_PORT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("FEATURE_CREATE_EMPLOYEE", "true")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_NAME", "employee_api_test")
	t.Setenv("DB_USER", "test_user")
	t.Setenv("DB_PASSWORD", "test_password")
	t.Setenv("DB_SSLMODE", "disable")
	t.Setenv("DB_MAX_CONNS", "2")
	_ = os.Unsetenv("APP_PORT")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.AppPort != "8080" {
		t.Fatalf("expected default port 8080, got %s", cfg.AppPort)
	}
}
