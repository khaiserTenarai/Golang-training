package utility

import "errors"

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

func ValidateDifferentAccounts(
	fromID int,
	toID int,
) error {

	if fromID == toID {
		return errors.New(
			"source and destination accounts must be different",
		)
	}

	return nil
}
