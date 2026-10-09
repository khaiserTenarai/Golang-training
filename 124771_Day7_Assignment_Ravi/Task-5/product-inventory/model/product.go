package model

import "time"

type Product struct {
	ID                int64
	Name              string
	Description       string
	Price             float64
	StockQuantity     int
	LowStockThreshold int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
