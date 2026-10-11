// Task 8: Separate Secrets from Configuration
// Secrets are loaded from dedicated env vars, never logged in plain text.

package config

import "os"

type SecretsConfig struct {
	AdminToken string
	APIKey     string
	JWTSecret  string
	DBPassword string
}

func LoadSecrets() *SecretsConfig {
	return &SecretsConfig{
		AdminToken: getSecretEnv("ADMIN_TOKEN", "admin-secret-token-123"),
		APIKey:     getSecretEnv("API_KEY", "dev-api-key-456"),
		JWTSecret:  getSecretEnv("JWT_SECRET", "dev-jwt-secret-789"),
		DBPassword: getSecretEnv("DB_PASSWORD", ""),
	}
}

func getSecretEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
