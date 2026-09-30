package repository

import "money_transfer/model"

type TransferRepository interface {
	TransferMoney(transfer model.Transfer) error
}
