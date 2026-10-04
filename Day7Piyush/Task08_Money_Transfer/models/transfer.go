package models

import "time"

type Account struct {
	ID         int
	HolderName string
	Balance    float64
	CreatedAt  time.Time
}

type TransferLog struct {
	ID            int
	FromAccountID int
	FromName      string
	ToAccountID   int
	ToName        string
	Amount        float64
	Status        string
	Description   string
	CreatedAt     time.Time
}
