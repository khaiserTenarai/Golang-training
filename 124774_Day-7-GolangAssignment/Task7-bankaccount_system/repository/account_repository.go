package repository

import "bankaccount/model"

type AccountRepository interface {
	Create(account model.Account) error

	Deposit(id int, amount float64) error

	Withdraw(id int, amount float64) error

	FindByID(id int) (model.Account, error)

	FindTransactions(id int) ([]model.Transaction, error)
}
