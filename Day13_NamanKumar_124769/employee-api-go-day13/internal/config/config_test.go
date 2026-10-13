package config

import (
	"testing"
	"time"
)

func TestValidateAcceptsDevelopmentDefaults(t *testing.T) {
	cfg := Config{
		AppEnv: "development", Port: "8080",
		ReadTimeout: time.Second, ReadHeaderTimeout: time.Second,
		WriteTimeout: time.Second, IdleTimeout: time.Second,
		ShutdownTimeout: time.Second, TokenTTL: time.Hour, LogLevel: "debug",
	}
	if err := Validate(cfg, Secrets{JWTSecret: "development-only"}); err != nil {
		t.Fatalf("expected valid config, got %v", err)
	}
}

func TestValidateRejectsWeakProductionSecret(t *testing.T) {
	cfg := Config{
		AppEnv: "production", Port: "8080",
		ReadTimeout: time.Second, ReadHeaderTimeout: time.Second,
		WriteTimeout: time.Second, IdleTimeout: time.Second,
		ShutdownTimeout: time.Second, TokenTTL: time.Hour, LogLevel: "info",
	}
	if err := Validate(cfg, Secrets{JWTSecret: "short"}); err == nil {
		t.Fatal("expected weak production secret to be rejected")
	}
}
