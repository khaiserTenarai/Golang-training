package model

import "time"

// Product mirrors one row in the "product" table.
type Product struct {
	ID        int
	Name      string
	SKU       string
	Category  string
	Price     float64
	Quantity  int
	CreatedAt time.Time
}
