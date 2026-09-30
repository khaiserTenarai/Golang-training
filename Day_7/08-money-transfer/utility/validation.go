package utility
import (
	"errors"

	"money_transfer/model"
)

var (
	InvalidAccountID = errors.New("invalid account ID")
	InvalidAmount    = errors.New("amount must be greater than 0")
	SameAccount      = errors.New("sender and receiver cannot be same")
)

func ValidateTransfer(
	transfer model.Transfer,
) error {

	if transfer.FromAccountID <= 0 ||
		transfer.ToAccountID <= 0 {

		return InvalidAccountID
	}

	if transfer.Amount <= 0 {
		return InvalidAmount
	}

	if transfer.FromAccountID ==
		transfer.ToAccountID {

		return SameAccount
	}

	return nil
}
