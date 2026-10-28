package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port               string
	DepartmentURL      string
	HTTPTimeout        time.Duration
	RetryCount         int
	BreakerThreshold   int
	BreakerOpenTimeout time.Duration
}

func Load() Config {
	return Config{
		Port:               env("EMPLOYEE_PORT", "8081"),
		DepartmentURL:      env("DEPARTMENT_SERVICE_URL", "http://localhost:8082"),
		HTTPTimeout:        durationEnv("HTTP_TIMEOUT", 3*time.Second),
		RetryCount:         intEnv("RETRY_COUNT", 3),
		BreakerThreshold:   intEnv("CIRCUIT_BREAKER_FAILURE_THRESHOLD", 5),
		BreakerOpenTimeout: durationEnv("CIRCUIT_BREAKER_OPEN_TIMEOUT", 10*time.Second),
	}
}
func env(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}
func intEnv(k string, fallback int) int {
	v, err := strconv.Atoi(os.Getenv(k))
	if err != nil || v < 1 {
		return fallback
	}
	return v
}
func durationEnv(k string, fallback time.Duration) time.Duration {
	v, err := time.ParseDuration(os.Getenv(k))
	if err != nil || v <= 0 {
		return fallback
	}
	return v
}
