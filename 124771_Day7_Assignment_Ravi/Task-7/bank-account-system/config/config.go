package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	return Config{
		DBHost:     get("DB_HOST", ""),
		DBPort:     get("DB_PORT", ""),
		DBUser:     get("DB_USER", ""),
		DBPassword: get("DB_PASSWORD", ""),
		DBName:     get("DB_NAME", ""),
		DBSSLMode:  get("DB_SSLMODE", ""),
	}
}

func (c Config) DatabaseURL() string {
	return "postgres://" + c.DBUser + ":" + c.DBPassword +
		"@" + c.DBHost + ":" + c.DBPort + "/" +
		c.DBName + "?sslmode=" + c.DBSSLMode
}

func get(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
