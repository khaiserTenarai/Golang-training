package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores normal application and database settings.
// Passwords and other secrets belong in the Secrets struct.
type Config struct {
	AppEnv  string
	AppPort string

	DBHost    string
	DBPort    string
	DBUser    string
	DBName    string
	DBSSLMode string
}

// Load reads settings for the selected environment.
func Load() *Config {
	// Read the base .env file to find APP_ENV.
	baseValues, _ := godotenv.Read(".env")

	// System environment variables have the highest priority.
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = baseValues["APP_ENV"]
	}
	if env == "" {
		env = "development"
	}

	// Read settings only from the selected environment file.
	fileValues, _ := godotenv.Read(".env." + env)

	// Start with default settings for that environment.
	var cfg *Config

	switch env {
	case "development":
		cfg = DevelopmentDefaults()
	case "test":
		cfg = TestDefaults()
	case "production":
		cfg = ProductionDefaults()
	default:
		cfg = &Config{
			AppEnv:  env,
			AppPort: "8080",
		}
	}

	// Override defaults with system or environment-file values.
	cfg.AppEnv = env
	cfg.AppPort = getValue("APP_PORT", fileValues, cfg.AppPort)
	cfg.DBHost = getValue("DB_HOST", fileValues, cfg.DBHost)
	cfg.DBPort = getValue("DB_PORT", fileValues, cfg.DBPort)
	cfg.DBUser = getValue("DB_USER", fileValues, cfg.DBUser)
	cfg.DBName = getValue("DB_NAME", fileValues, cfg.DBName)
	cfg.DBSSLMode = getValue("DB_SSLMODE", fileValues, cfg.DBSSLMode)

	return cfg
}

// getValue returns the system value, file value, or default.
func getValue(key string, values map[string]string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	if value, exists := values[key]; exists && value != "" {
		return value
	}

	return fallback
}

// Validate checks required settings and port numbers.
func (c *Config) Validate() error {
	if c.DBHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if c.DBUser == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if c.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	// Check that the application port is valid.
	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	// Check that the database port is valid.
	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	// Require SSL when using the production environment.
	if c.AppEnv == "production" &&
		c.DBSSLMode != "require" &&
		c.DBSSLMode != "verify-ca" &&
		c.DBSSLMode != "verify-full" {
		return fmt.Errorf(
			"production requires DB_SSLMODE=require, verify-ca, or verify-full",
		)
	}

	return nil
}

// DevelopmentDefaults returns local development settings.
func DevelopmentDefaults() *Config {
	return &Config{
		AppEnv:    "development",
		AppPort:   "8080",
		DBHost:    "localhost",
		DBPort:    "5432",
		DBUser:    "postgres",
		DBName:    "employee_db",
		DBSSLMode: "disable",
	}
}
