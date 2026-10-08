package controller

import "bank-account-system/model"

type AccountController interface {
	CreateAccount(account model.Account) error
	GetAccount(id int) (*model.Account, error)
	Deposit(id int, amount float64) error
	Withdraw(id int, amount float64) error
	GetTransactionHistory(accountID int) ([]model.Transaction, error)
}
