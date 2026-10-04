package utility

import "errors"

func ValidateName(name string) error {

	if name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func ValidateEmail(email string) error {

	if email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}

func ValidatePrice(price float64) error {

	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}

func ValidateQuantity(quantity int) error {

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return nil
}
