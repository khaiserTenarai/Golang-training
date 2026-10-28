package config

import (
	"fmt"
	"os"
)

// Secrets stores sensitive values.
type Secrets struct {
	DBPassword string
}

// LoadSecrets reads sensitive values from the process environment.
func LoadSecrets() (*Secrets, error) {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return &Secrets{
		DBPassword: password,
	}, nil
}