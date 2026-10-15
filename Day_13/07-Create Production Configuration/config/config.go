package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config stores application and database settings.
type Config struct {
	AppEnv  string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// Load reads the selected environment's configuration.
func Load() *Config {
	// Read the base .env file to discover APP_ENV.
	// Missing .env files are allowed.
	baseValues, _ := godotenv.Read(".env")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = baseValues["APP_ENV"]
	}
	if env == "" {
		env = "development"
	}

	// Read only the configuration file for the selected environment.
	fileValues, _ := godotenv.Read(".env." + env)

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

	cfg.AppEnv = env
	cfg.AppPort = getValue("APP_PORT", fileValues, cfg.AppPort)
	cfg.DBHost = getValue("DB_HOST", fileValues, cfg.DBHost)
	cfg.DBPort = getValue("DB_PORT", fileValues, cfg.DBPort)
	cfg.DBUser = getValue("DB_USER", fileValues, cfg.DBUser)
	cfg.DBPassword = getValue("DB_PASSWORD", fileValues, cfg.DBPassword)
	cfg.DBName = getValue("DB_NAME", fileValues, cfg.DBName)
	cfg.DBSSLMode = getValue("DB_SSLMODE", fileValues, cfg.DBSSLMode)

	return cfg
}

// getValue uses system environment variables first,
// then the selected .env file, then the fallback value.
func getValue(key string, fileValues map[string]string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	if value, exists := fileValues[key]; exists && value != "" {
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
	if c.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if c.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	appPort, err := strconv.Atoi(c.AppPort)
	if err != nil || appPort < 1 || appPort > 65535 {
		return fmt.Errorf("APP_PORT must be between 1 and 65535")
	}

	dbPort, err := strconv.Atoi(c.DBPort)
	if err != nil || dbPort < 1 || dbPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535")
	}

	if c.AppEnv == "production" && c.DBSSLMode != "require" &&
		c.DBSSLMode != "verify-ca" && c.DBSSLMode != "verify-full" {
		return fmt.Errorf("production requires DB_SSLMODE=require, verify-ca, or verify-full")
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
