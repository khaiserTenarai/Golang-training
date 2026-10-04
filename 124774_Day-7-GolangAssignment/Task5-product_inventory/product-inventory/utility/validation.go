package utility

import (
	"errors"
	"strings"
)

func ValidateName(name string) error {

	if strings.TrimSpace(name) == "" {
		return errors.New("product name cannot be empty")
	}

	return nil
}

func ValidatePrice(price float64) error {

	if price < 0 {
		return errors.New("price cannot be negative")
	}

	return nil
}

func ValidateStock(stock int) error {

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return nil
}

func ValidateLowStockLimit(limit int) error {

	if limit < 0 {
		return errors.New("low stock limit cannot be negative")
	}

	return nil
}

func ValidateID(id int) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	return nil
}

func ValidateQuantity(quantity int) error {

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return nil
}
