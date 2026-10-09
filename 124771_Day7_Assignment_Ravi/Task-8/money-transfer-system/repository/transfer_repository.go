package repository

import "money-transfer-system/model"

type TransferRepository interface {
	TransferMoney(transfer model.Transfer) error
	GetAccount(accountNumber string) (model.Account, error)
}
