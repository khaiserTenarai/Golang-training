package model

import "time"

type Account struct {
    ID            int
    AccountNumber string
    CustomerName  string
    Balance       float64
    CreatedAt     time.Time
}
