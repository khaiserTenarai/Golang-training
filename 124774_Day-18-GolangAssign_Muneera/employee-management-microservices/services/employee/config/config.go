package config

import (
	"fmt"
	"os"
)

// Config stores Employee Service settings.
type Config struct {
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	Port                 string
	DepartmentServiceURL string
}

func Load() Config {
	return Config{
		DBHost:               getEnv("EMPLOYEE_DB_HOST", "localhost"),
		DBPort:               getEnv("EMPLOYEE_DB_PORT", "5432"),
		DBUser:               getEnv("EMPLOYEE_DB_USER", "postgres"),
		DBPassword:           os.Getenv("EMPLOYEE_DB_PASSWORD"),
		DBName:               getEnv("EMPLOYEE_DB_NAME", "employee_db"),
		Port:                 getEnv("EMPLOYEE_SERVICE_PORT", "8081"),
		DepartmentServiceURL: getEnv("DEPARTMENT_SERVICE_URL", "http://localhost:8082"),
	}
}

// DSN returns the PostgreSQL connection string.
func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
