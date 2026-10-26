package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort              string
	EmployeeServiceURL   string
	DepartmentServiceURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	return &Config{
		AppPort:              getEnv("APP_PORT", "8080"),
		EmployeeServiceURL:   getEnv("EMPLOYEE_SERVICE_URL", "http://localhost:8081"),
		DepartmentServiceURL: getEnv("DEPARTMENT_SERVICE_URL", "http://localhost:8082"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
