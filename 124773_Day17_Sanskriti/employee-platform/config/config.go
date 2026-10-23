package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port string
	DatabaseURL string
	JWTSecret string
	JWTExpiresMinutes int
}

func Load() (Config, error) {
	c := Config{
		Port: os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiresMinutes: 60,
	}
	if c.Port == "" { c.Port = "8080" }
	if v := os.Getenv("JWT_EXPIRES_MINUTES"); v != "" {
		n, err := strconv.Atoi(v); if err != nil || n < 1 { return c, fmt.Errorf("JWT_EXPIRES_MINUTES must be positive") }; c.JWTExpiresMinutes = n
	}
	if c.DatabaseURL == "" { return c, fmt.Errorf("DATABASE_URL is required") }
	if len(c.JWTSecret) < 32 { return c, fmt.Errorf("JWT_SECRET must be at least 32 characters") }
	return c, nil
}
