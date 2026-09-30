package repository
import "bank_account/model"

type AccountRepository interface {

	Create(account model.Account) error

	FindByID(id int) (model.Account, error)

	Deposit(id int, amount float64) error

	Withdraw(id int, amount float64) error

	GetTransactions(id int) ([]model.Transaction, error)
}
