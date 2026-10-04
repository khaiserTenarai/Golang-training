package utility

import "errors"

func ValidateCustomerName(name string) error {

	if name == "" {
		return errors.New("customer name cannot be empty")
	}

	return nil
}

func ValidateProductID(id int) error {

	if id <= 0 {
		return errors.New("product ID must be greater than 0")
	}

	return nil
}

func ValidateQuantity(quantity int) error {

	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return nil
}
