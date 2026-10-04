package model

type Account struct {
	ID            int
	AccountNumber string
	Name          string
	Balance       float64
}

type Transaction struct {
	ID           int
	AccountID    int
	Type         string
	Amount       float64
	BalanceAfter float64
}
