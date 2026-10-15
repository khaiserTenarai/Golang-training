package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	DBURL            string
	JWTSecret        string
	JWTExpiryMinutes int
	AdminEmail       string
	AllowedOrigins   []string
	UploadDir        string
}

func Load() *Config {
	_ = godotenv.Load() 

	expiry, err := strconv.Atoi(get("JWT_EXPIRY_MINUTES", "60"))
	if err != nil || expiry <= 0 {
		expiry = 60
	}

	return &Config{
		Port:             get("PORT", "8080"),
		DBURL:            get("DB_URL", ""),
		JWTSecret:        get("JWT_SECRET", ""),
		JWTExpiryMinutes: expiry,
		AdminEmail:       strings.ToLower(get("ADMIN_EMAIL", "")),
		AllowedOrigins:   strings.Split(get("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
		UploadDir:        get("UPLOAD_DIR", "uploads"),
	}
}

func get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
