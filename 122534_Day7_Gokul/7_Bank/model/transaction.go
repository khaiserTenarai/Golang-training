package model

import "time"

// Transaction mirrors one row in the "transaction" table - a single
// deposit or withdrawal.
type Transaction struct {
	ID           int
	AccountID    int
	Type         string // "DEPOSIT" or "WITHDRAWAL"
	Amount       float64
	BalanceAfter float64
	CreatedAt    time.Time
}
