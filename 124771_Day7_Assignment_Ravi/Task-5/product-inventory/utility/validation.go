package utility

import (
	"fmt"
	"strings"
)

func ValidateProductID(
	id int64,
) error {

	if id <= 0 {
		return fmt.Errorf(
			"product ID must be greater than 0",
		)
	}

	return nil
}

func ValidateProductName(
	name string,
) error {

	if strings.TrimSpace(name) == "" {
		return fmt.Errorf(
			"product name cannot be empty",
		)
	}

	return nil
}

func ValidatePrice(
	price float64,
) error {

	if price < 0 {
		return fmt.Errorf(
			"price cannot be negative",
		)
	}

	return nil
}

func ValidateStock(
	stock int,
) error {

	if stock < 0 {
		return fmt.Errorf(
			"stock quantity cannot be negative",
		)
	}

	return nil
}

func ValidateThreshold(
	threshold int,
) error {

	if threshold < 0 {
		return fmt.Errorf(
			"low stock threshold cannot be negative",
		)
	}

	return nil
}

func ValidateStockChange(
	quantity int,
) error {

	if quantity <= 0 {
		return fmt.Errorf(
			"stock change quantity must be greater than 0",
		)
	}

	return nil
}
