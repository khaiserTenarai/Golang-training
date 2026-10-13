package main

import (
	"fmt"
	"os"
	"strings"
)

// Secrets are loaded separately from ordinary application configuration.
type Secrets struct {
	AdminToken string
	UserToken  string
}

func loadSecrets() (Secrets, error) {
	s := Secrets{
		AdminToken: strings.TrimSpace(os.Getenv("ADMIN_TOKEN")),
		UserToken:  strings.TrimSpace(os.Getenv("USER_TOKEN")),
	}
	if s.AdminToken == "" || s.UserToken == "" {
		return Secrets{}, fmt.Errorf("ADMIN_TOKEN and USER_TOKEN are required")
	}
	if s.AdminToken == s.UserToken {
		return Secrets{}, fmt.Errorf("ADMIN_TOKEN and USER_TOKEN must be different")
	}
	return s, nil
}
