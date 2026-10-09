package model

import "time"

type Transaction struct {
    ID              int
    AccountID       int
    TransactionType string
    Amount          float64
    CreatedAt       time.Time
}
