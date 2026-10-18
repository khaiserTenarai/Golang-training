package config

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func lookupFrom(m map[string]string) LookupFunc {
	return func(k string) (string, bool) { v, ok := m[k]; return v, ok }
}

func TestDevelopmentDefaults(t *testing.T) {
	cfg, err := LoadFrom(lookupFrom(map[string]string{}))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.App.Env != EnvDevelopment || cfg.Server.Port != 8080 || cfg.Log.Level != "debug" || !cfg.Features.SeedData {
		t.Fatalf("unexpected development defaults: %+v", cfg)
	}
}

func TestTestProfileAndOverride(t *testing.T) {
	cfg, err := LoadFrom(lookupFrom(map[string]string{"APP_ENV": "test"}))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Server.Port != 8081 || cfg.Log.Level != "error" || cfg.Features.SeedData {
		t.Fatalf("unexpected test defaults: %+v", cfg)
	}

	cfg, err = LoadFrom(lookupFrom(map[string]string{"APP_ENV": "test", "SERVER_PORT": "9999"}))
	if err != nil || cfg.Server.Port != 9999 {
		t.Fatalf("override failed: %v %+v", err, cfg)
	}
}

func TestProductionNeedsASecret(t *testing.T) {
	_, err := LoadFrom(lookupFrom(map[string]string{"APP_ENV": "production"}))
	if err == nil || !strings.Contains(err.Error(), "API_KEY") {
		t.Fatalf("expected an API_KEY error, got %v", err)
	}

	_, err = LoadFrom(lookupFrom(map[string]string{"APP_ENV": "production", "API_KEY": DevAPIKey}))
	if err == nil {
		t.Fatal("placeholder key must be rejected in production")
	}
}

func TestProductionRules(t *testing.T) {
	good := map[string]string{"APP_ENV": "production", "API_KEY": "a-long-production-key-123"}
	cfg, err := LoadFrom(lookupFrom(good))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Log.Format != "json" || cfg.Features.ExposeSalary {
		t.Fatalf("unexpected production defaults: %+v", cfg)
	}

	good["LOG_LEVEL"] = "debug"
	if _, err := LoadFrom(lookupFrom(good)); err == nil {
		t.Fatal("debug logging must be rejected in production")
	}
}

func TestInvalidValuesAreAllReported(t *testing.T) {
	_, err := LoadFrom(lookupFrom(map[string]string{"SERVER_PORT": "abc", "FEATURE_SEARCH": "maybe", "SHUTDOWN_TIMEOUT": "soon"}))
	if err == nil {
		t.Fatal("expected errors")
	}
	for _, want := range []string{"SERVER_PORT", "FEATURE_SEARCH", "SHUTDOWN_TIMEOUT"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should mention %s: %v", want, err)
		}
	}

	if _, err := LoadFrom(lookupFrom(map[string]string{"SERVER_PORT": "70000"})); err == nil {
		t.Fatal("port out of range must be rejected")
	}
	if _, err := LoadFrom(lookupFrom(map[string]string{"APP_ENV": "staging"})); err == nil {
		t.Fatal("unknown environment must be rejected")
	}
}

func TestSecretsNeverPrint(t *testing.T) {
	cfg, err := LoadFrom(lookupFrom(map[string]string{"API_KEY": "super-secret-value"}))
	if err != nil {
		t.Fatal(err)
	}
	asJSON, _ := json.Marshal(cfg)
	outputs := []string{fmt.Sprintf("%v", cfg), fmt.Sprintf("%+v", cfg.Secrets), string(asJSON), fmt.Sprint(cfg.Summary())}
	for _, out := range outputs {
		if strings.Contains(out, "super-secret-value") {
			t.Fatalf("secret leaked: %s", out)
		}
	}
	if cfg.Secrets.APIKey.Reveal() != "super-secret-value" {
		t.Fatal("Reveal must return the real value")
	}
}
