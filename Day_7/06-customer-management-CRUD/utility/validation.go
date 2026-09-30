package utility
import (
	"errors"
	"strings"

	"cms/model"
)

var (
	InvalidCustomerID   = errors.New("invalid customer ID")
	InvalidCustomerName = errors.New("invalid customer name")
	InvalidCustomerAge  = errors.New("invalid customer age")
	InvalidCustomerEmail = errors.New("invalid customer email")
	InvalidCustomerPhone = errors.New("invalid customer phone")
	InvalidPage         = errors.New("invalid page number")
	InvalidLimit        = errors.New("invalid limit")
	InvalidKeyword      = errors.New("search keyword cannot be empty")
)

func ValidateCustomer(customer model.Customer) error {

	if strings.TrimSpace(customer.Name) == "" {
		return InvalidCustomerName
	}

	if customer.Age < 18 || customer.Age > 100 {
		return InvalidCustomerAge
	}

	return nil
}

func ValidateCustomerID(id int) error {

	if id <= 0 {
		return InvalidCustomerID
	}

	return nil
}

func ValidateEmail(email string) error {

	if !strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {
		return InvalidCustomerEmail
	}

	return nil
}

func ValidatePhone(phone string) error {

	if len(phone) < 10 || len(phone) > 15 {
		return InvalidCustomerPhone
	}

	return nil
}

func ValidatePagination(page int, limit int) error {

	if page <= 0 {
		return InvalidPage
	}

	if limit <= 0 {
		return InvalidLimit
	}

	return nil
}

func ValidateSearch(keyword string) error {

	if strings.TrimSpace(keyword) == "" {
		return InvalidKeyword
	}

	return nil
}
