package utility

import (
	"errors"
	"strings"
)

func ValidateUsername(username string) error {

	if strings.TrimSpace(username) == "" {
		return errors.New("username cannot be empty")
	}

	if len(username) < 3 {
		return errors.New("username must contain at least 3 characters")
	}

	return nil
}

func ValidatePassword(password string) error {

	if len(password) < 6 {
		return errors.New("password must contain at least 6 characters")
	}

	return nil
}

func ValidateRole(role string) error {

	if role != "admin" && role != "user" {
		return errors.New("role must be admin or user")
	}

	return nil
}
