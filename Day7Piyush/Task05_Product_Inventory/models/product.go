package models

import "time"

type Product struct {
	ID        int
	Name      string
	Category  string
	Price     float64
	Stock     int
	CreatedAt time.Time
}
