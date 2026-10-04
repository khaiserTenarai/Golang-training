package model

import "time"

type Customer struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type Product struct {
	ID            int
	Name          string
	Price         float64
	Stock         int
	LowStockLimit int
}

type Order struct {
	ID         int
	CustomerID int
	OrderDate  time.Time
}

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}
