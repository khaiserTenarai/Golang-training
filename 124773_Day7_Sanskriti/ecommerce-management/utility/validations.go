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

func ValidateEmail(email string) error {

	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	return nil
}

func ValidatePrice(price float64) error {

	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}

func ValidateStock(stock int) error {

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return nil
}

func ValidateQuantity(quantity int) error {

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return nil
}
