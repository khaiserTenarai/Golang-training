package model

import "time"

// Account mirrors one row in the "account" table. Its ID doubles as
// the account number in this simplified system.
type Account struct {
	ID                int
	AccountHolderName string
	Balance           float64
	CreatedAt         time.Time
}
