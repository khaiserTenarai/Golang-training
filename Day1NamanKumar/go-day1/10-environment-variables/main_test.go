package main

import "testing"

func TestLoadConfigDefaults(t *testing.T) {
	t.Setenv("APP_NAME", "")
	t.Setenv("APP_PORT", "")
	cfg, err := LoadConfig()
	if err != nil || cfg.AppName != "EmployeeApp" || cfg.Port != 8080 {
		t.Fatalf("got %+v, %v", cfg, err)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	t.Setenv("APP_NAME", "Employee API")
	t.Setenv("APP_PORT", "9090")
	cfg, err := LoadConfig()
	if err != nil || cfg.AppName != "Employee API" || cfg.Port != 9090 {
		t.Fatalf("got %+v, %v", cfg, err)
	}
}

func TestLoadConfigBadPort(t *testing.T) {
	t.Setenv("APP_PORT", "abc")
	if _, err := LoadConfig(); err == nil {
		t.Fatal("expected error for invalid port")
	}
}
