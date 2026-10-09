package model

import "time"

type Order struct {
    ID         int
    CustomerID int
    OrderDate  time.Time
    Status     string
}
