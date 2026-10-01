package config

import (
    "os"
    "github.com/joho/godotenv"
    "log"
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
		log.Println(".env file not found, using default values")
	}

	return Config{
		DBHost:     get("DB_HOST", "localhost"),
		DBPort:     get("DB_PORT", "5432"),
		DBUser:     get("DB_USER", "postgres"),
		DBPassword: get("DB_PASSWORD", "Info%40131"),
		DBName:     get("DB_NAME", "go_training"),
		DBSSLMode:  get("DB_SSLMODE", "disable"),
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
