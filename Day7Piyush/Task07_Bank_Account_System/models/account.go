package models

import "time"

type Account struct {
	ID          int
	HolderName  string
	AccountType string
	Balance     float64
	CreatedAt   time.Time
}

type Transaction struct {
	ID           int
	AccountID    int
	TxnType      string
	Amount       float64
	BalanceAfter float64
	Description  string
	CreatedAt    time.Time
}
