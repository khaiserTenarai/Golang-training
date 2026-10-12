package config

import (
 "fmt"
 "os"
 "strconv"
)

type Config struct {
 DBHost, DBPort, DBUser, DBPassword, DBName, ServerPort string
}

func Load() (Config, error) {
 c := Config{
  DBHost: env("DB_HOST", "localhost"), DBPort: env("DB_PORT", "5432"),
  DBUser: env("DB_USER", "postgres"), DBPassword: os.Getenv("DB_PASSWORD"),
  DBName: env("DB_NAME", "employee_db"), ServerPort: env("SERVER_PORT", "10015"),
 }
 if c.DBPassword == "" { return Config{}, fmt.Errorf("DB_PASSWORD is required in .env") }
 if _, err := strconv.Atoi(c.DBPort); err != nil { return Config{}, fmt.Errorf("DB_PORT must be numeric") }
 if _, err := strconv.Atoi(c.ServerPort); err != nil { return Config{}, fmt.Errorf("SERVER_PORT must be numeric") }
 return c, nil
}
func env(k, d string) string { if v := os.Getenv(k); v != "" { return v }; return d }
