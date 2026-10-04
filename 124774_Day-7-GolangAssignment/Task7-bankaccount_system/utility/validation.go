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

func ValidateAccountID(id int) error {

	if id <= 0 {
		return errors.New("account ID must be greater than 0")
	}

	return nil
}

func ValidateAmount(amount float64) error {

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	return nil
}

func ValidateOpeningBalance(balance float64) error {

	if balance < 0 {
		return errors.New("opening balance cannot be negative")
	}

	return nil
}
