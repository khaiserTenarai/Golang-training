package config

import (
	"os"
	"strings"
)

type Config struct {
	Address        string
	JWTSecret      string
	TokenTTLHours  int
	AllowedOrigins []string
	Mode           string
}

func Load() Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "local-development-secret-change-me"
	}

	origins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	cleanOrigins := make([]string, 0, len(origins))
	for _, origin := range origins {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			cleanOrigins = append(cleanOrigins, origin)
		}
	}
	if len(cleanOrigins) == 0 {
		cleanOrigins = []string{"http://localhost:3000"}
	}

	return Config{
		Address:        envOr("ADDRESS", ":8080"),
		JWTSecret:      secret,
		TokenTTLHours:  2,
		AllowedOrigins: cleanOrigins,
		Mode:           envOr("GIN_MODE", "debug"),
	}
}

func envOr(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
