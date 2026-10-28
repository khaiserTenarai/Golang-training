package config

import (
	"fmt"
	"os"
)

type Secrets struct {
	DBPassword string
}

// LoadSecrets reads sensitive values separately.
func LoadSecrets() (*Secrets, error) {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return &Secrets{
		DBPassword: password,
	}, nil
}