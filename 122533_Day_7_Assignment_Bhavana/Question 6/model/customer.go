package model

import "time"

// Customer mirrors one row in the "customer" table.
type Customer struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time
}
