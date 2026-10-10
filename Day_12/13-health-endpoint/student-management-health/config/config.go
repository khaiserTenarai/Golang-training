package config

import (
 "fmt"
 "os"

 "github.com/joho/godotenv"
)

type Config struct {
 DatabaseURL string
 ServerPort  string
}

func Load() (Config, error) {
 // Loading .env is optional; environment variables can also be supplied directly.
 _ = godotenv.Load()

 host := get("DB_HOST", "localhost")
 port := get("DB_PORT", "5432")
 user := get("DB_USER", "postgres")
 password := get("DB_PASSWORD", "postgres")
 dbName := get("DB_NAME", "studentdb")
 sslMode := get("DB_SSLMODE", "disable")

 return Config{
  DatabaseURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbName, sslMode),
  ServerPort:  get("SERVER_PORT", "8080"),
 }, nil
}

func get(key, fallback string) string {
 if value := os.Getenv(key); value != "" {
  return value
 }
 return fallback
}
