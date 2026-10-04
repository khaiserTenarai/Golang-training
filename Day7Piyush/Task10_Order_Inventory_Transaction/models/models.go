package models

import "time"

type Product struct {
	ID        int
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
}

type Order struct {
	ID          int
	ProductID   int
	ProductName string
	Quantity    int
	TotalPrice  float64
	Status      string
	CreatedAt   time.Time
}
