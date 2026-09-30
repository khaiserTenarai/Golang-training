package utility
import (
	"errors"
	"strings"

	"bank_account/model"
)

var (
	InvalidID      = errors.New("invalid account ID")
	InvalidName    = errors.New("invalid account name")
	InvalidEmail   = errors.New("invalid email")
	InvalidAmount  = errors.New("amount must be greater than 0")
)

func ValidateAccount(account model.Account) error {

	if strings.TrimSpace(account.Name) == "" {
		return InvalidName
	}

	if !strings.Contains(account.Email, "@") {
		return InvalidEmail
	}

	return nil
}

func ValidateID(id int) error {

	if id <= 0 {
		return InvalidID
	}

	return nil
}

func ValidateAmount(amount float64) error {

	if amount <= 0 {
		return InvalidAmount
	}

	return nil
}
