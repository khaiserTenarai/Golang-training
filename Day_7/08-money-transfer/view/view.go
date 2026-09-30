package view
import "money_transfer/model"

type TransferView interface {

	ShowMenu() int

	ReadTransfer() model.Transfer
}
