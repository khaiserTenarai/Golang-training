package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() (Config, error) {
	// .env is optional; environment variables can also be used.
	_ = godotenv.Load()
	get := func(key, fallback string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return fallback
	}
	host := get("DB_HOST", "localhost")
	port := get("DB_PORT", "5432")
	user := get("DB_USER", "postgres")
	password := get("DB_PASSWORD", "postgres")
	name := get("DB_NAME", "studentdb")
	sslmode := get("DB_SSLMODE", "disable")
	return Config{
		DatabaseURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, name, sslmode),
		Port:        get("PORT", "8080"),
	}, nil
}
