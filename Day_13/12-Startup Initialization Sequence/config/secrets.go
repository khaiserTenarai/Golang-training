package config

import (
	"fmt"
	"os"
)

// Secrets stores sensitive values separately from normal configuration.
type Secrets struct {
	DBPassword string
}

// LoadSecrets reads the database password from the environment.
func LoadSecrets() (*Secrets, error) {
	password := os.Getenv("DB_PASSWORD")

	// Do not start without a database password.
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return &Secrets{
		DBPassword: password,
	}, nil
}