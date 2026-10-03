package repository

import "bank-account-system/model"

type AccountRepository interface {
	CreateAccount(account model.Account) error
	GetAccount(id int) (*model.Account, error)
	Deposit(id int, amount float64) error
	Withdraw(id int, amount float64) error
	GetTransactionHistory(accountID int) ([]model.Transaction, error)
}
