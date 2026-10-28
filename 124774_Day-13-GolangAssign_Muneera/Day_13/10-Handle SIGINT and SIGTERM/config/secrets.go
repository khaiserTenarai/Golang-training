package config

import (
	"fmt"
	"os"
)

// Secrets stores sensitive application values.
type Secrets struct {
	DBPassword string
}

// LoadSecrets reads secrets from process environment variables.
func LoadSecrets() (*Secrets, error) {
	password := os.Getenv("DB_PASSWORD")

	// Stop startup if the database password is missing.
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return &Secrets{
		DBPassword: password,
	}, nil
}