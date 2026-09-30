package view
import (
	"bank_account/model"
)

type AccountView interface {

	ShowMenu() int

	ReadAccount() model.Account

	ReadID() int

	ReadAmount() float64

	DisplayAccount(account model.Account)

	DisplayTransactions(
		transactions []model.Transaction,
	)
}
