package service

import (
	"bank-account-system/model"
)

type AccountService interface {
	CreateAccount(account model.Account) error

	Deposit(
		accountNumber string,
		amount float64,
	) error

	Withdraw(
		accountNumber string,
		amount float64,
	) error

	GetBalance(
		accountNumber string,
	) (float64, error)

	GetTransactionHistory(
		accountNumber string,
	) ([]model.Transaction, error)
}
