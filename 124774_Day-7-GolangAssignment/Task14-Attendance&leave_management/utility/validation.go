package utility

import (
	"errors"
	"strings"
)

func ValidateName(name string) error {

	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func ValidateID(id int) error {

	if id <= 0 {
		return errors.New("employee ID must be greater than 0")
	}

	return nil
}

func ValidateReason(reason string) error {

	if strings.TrimSpace(reason) == "" {
		return errors.New("leave reason cannot be empty")
	}

	return nil
}
