package models

import "time"

type Customer struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

type Product struct {
	ID        int
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
}

type Order struct {
	ID           int
	CustomerID   int
	CustomerName string
	TotalAmount  float64
	Status       string
	CreatedAt    time.Time
	Items        []OrderItem
}

type OrderItem struct {
	ID          int
	OrderID     int
	ProductID   int
	ProductName string
	Quantity    int
	UnitPrice   float64
	Subtotal    float64
}
