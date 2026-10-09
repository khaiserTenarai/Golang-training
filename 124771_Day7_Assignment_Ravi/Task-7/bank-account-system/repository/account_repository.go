package repository

import "bank-account-system/model"

type AccountRepository interface {
	CreateAccount(account model.Account) error
	GetAccountByNumber(accountNumber string) (model.Account, error)
	UpdateBalance(accountID int, balance float64) error
	AddTransaction(transaction model.Transaction) error
	GetTransactions(accountID int) ([]model.Transaction, error)
}
