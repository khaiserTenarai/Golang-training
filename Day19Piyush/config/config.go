// Task 7: Externalize Service URLs
// Service URLs are read from environment variables with defaults.

package config

import "os"

type ServiceConfig struct {
	EmployeeServiceURL   string
	DepartmentServiceURL string
	HTTPTimeoutSecs      int
	RetryMaxAttempts     int
	CircuitBreakerThreshold int
}

func Load() ServiceConfig {
	return ServiceConfig{
		EmployeeServiceURL:      getEnv("EMPLOYEE_SERVICE_URL", "http://localhost:8081"),
		DepartmentServiceURL:    getEnv("DEPARTMENT_SERVICE_URL", "http://localhost:8082"),
		HTTPTimeoutSecs:         5,
		RetryMaxAttempts:        3,
		CircuitBreakerThreshold: 3,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
