package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config contains runtime settings. Values are externalized so URLs/ports
// can change without editing application code. Missing .env is allowed in deployment.
type Config struct {
	Port                    string
	DBHost                  string
	DBPort                  string
	DBUser                  string
	DBPassword              string
	DBName                  string
	DBSSLMode               string
	DepartmentServiceURL    string
	HTTPTimeoutSeconds      int
	RetryCount              int
	CircuitFailureThreshold int
	CircuitOpenSeconds      int
}

func Load() (Config, error) {
	_ = godotenv.Load() // Local development convenience; environment variables still work.
	c := Config{
		Port: getEnv("PORT", "8080"), DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"), DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD"), DBName: getEnv("DB_NAME", "microservices_db"),
		DBSSLMode:            getEnv("DB_SSLMODE", "disable"),
		DepartmentServiceURL: getEnv("DEPARTMENT_SERVICE_URL", "http://localhost:8082"),
	}
	var err error
	if c.HTTPTimeoutSeconds, err = getInt("HTTP_TIMEOUT_SECONDS", 2); err != nil {
		return c, err
	}
	if c.RetryCount, err = getInt("RETRY_COUNT", 2); err != nil {
		return c, err
	}
	if c.CircuitFailureThreshold, err = getInt("CIRCUIT_FAILURE_THRESHOLD", 3); err != nil {
		return c, err
	}
	if c.CircuitOpenSeconds, err = getInt("CIRCUIT_OPEN_SECONDS", 10); err != nil {
		return c, err
	}
	if c.Port == "" || c.DBHost == "" || c.DBName == "" || c.DBUser == "" {
		return c, fmt.Errorf("PORT, DB_HOST, DB_USER and DB_NAME must not be empty")
	}
	if c.HTTPTimeoutSeconds < 1 || c.RetryCount < 0 || c.CircuitFailureThreshold < 1 || c.CircuitOpenSeconds < 1 {
		return c, fmt.Errorf("timeout/open seconds/threshold must be positive and retries cannot be negative")
	}
	return c, nil
}
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
func getInt(key string, fallback int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid integer for %s: %w", key, err)
	}
	return n, nil
}
func (c Config) DSN() string {
	// PostgreSQL connection string used by GORM's postgres driver.
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}
