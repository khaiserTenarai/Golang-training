package service
import "money_transfer/model"

type TransferService interface {
	TransferMoney(transfer model.Transfer) error
}
