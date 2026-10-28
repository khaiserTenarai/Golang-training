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
	// Read the password without printing or logging it.
	password := os.Getenv("DB_PASSWORD")

	// Stop startup if the password has not been configured.
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return &Secrets{
		DBPassword: password,
	}, nil
}
