package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost, DBPort, DBUser, DBPassword, DBName, DBSSLMode string
	ServerPort, JWTSecret                                 string
	JWTHours                                              int
}

func Load() Config {
	hours, _ := strconv.Atoi(getenv("JWT_HOURS", "24"))
	return Config{
		DBHost:     getenv("DB_HOST", "localhost"),
		DBPort:     getenv("DB_PORT", "5432"),
		DBUser:     getenv("DB_USER", "postgres"),
		DBPassword: getenv("DB_PASSWORD", ""),
		DBName:     getenv("DB_NAME", "employee_platform"),
		DBSSLMode:  getenv("DB_SSLMODE", "disable"),
		ServerPort: getenv("SERVER_PORT", "8080"),
		JWTSecret:  getenv("JWT_SECRET", ""),
		JWTHours:   hours,
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
