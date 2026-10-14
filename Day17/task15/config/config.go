package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port      int
	Env       string
	JWTSecret string
}

func LoadConfig() Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "super-secret-jwt-key"
	}

	return Config{
		Port:      port,
		Env:       env,
		JWTSecret: secret,
	}
}
