package model

import "time"

type OrderDetails struct {
    OrderID       int
    CustomerName  string
    CustomerEmail string
    OrderDate     time.Time
    Status        string
    ProductName   string
    Quantity      int
    Price         float64
    ItemTotal     float64
}
