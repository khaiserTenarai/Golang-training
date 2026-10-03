package view

import "bank-account-system/model"

type AccountView interface {
	Start()
	CreateAccount()
	Deposit()
	Withdraw()
	BalanceEnquiry()
	TransactionHistory()
	DisplayAccount(account *model.Account)
	DisplayTransactions(transactions []model.Transaction)
}
