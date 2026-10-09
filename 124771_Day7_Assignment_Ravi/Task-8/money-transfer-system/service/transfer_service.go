package service

import "money-transfer-system/model"

type TransferService interface {
	TransferMoney(transfer model.Transfer) error
	GetAccount(accountNumber string) (model.Account, error)
}
