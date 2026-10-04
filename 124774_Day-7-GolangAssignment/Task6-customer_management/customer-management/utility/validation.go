package utility

import (
	"errors"
	"strings"
)

func ValidateName(name string) error {

	if strings.TrimSpace(name) == "" {
		return errors.New("customer name cannot be empty")
	}

	return nil
}

func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("email cannot be empty")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	return nil
}

func ValidatePhone(phone string) error {

	if strings.TrimSpace(phone) == "" {
		return errors.New("phone cannot be empty")
	}

	return nil
}

func ValidateCity(city string) error {

	if strings.TrimSpace(city) == "" {
		return errors.New("city cannot be empty")
	}

	return nil
}

func ValidateID(id int) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	return nil
}

func ValidatePage(page int) error {

	if page <= 0 {
		return errors.New("page must be greater than 0")
	}

	return nil
}

func ValidatePageSize(pageSize int) error {

	if pageSize <= 0 {
		return errors.New("page size must be greater than 0")
	}

	return nil
}
