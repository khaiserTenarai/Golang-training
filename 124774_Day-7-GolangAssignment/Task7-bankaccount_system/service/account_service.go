package service

import "bankaccount/model"

type AccountService interface {
	CreateAccount(
		account model.Account,
	) error

	Deposit(
		id int,
		amount float64,
	) error

	Withdraw(
		id int,
		amount float64,
	) error

	FindAccount(
		id int,
	) (model.Account, error)

	FindTransactions(
		id int,
	) ([]model.Transaction, error)
}
