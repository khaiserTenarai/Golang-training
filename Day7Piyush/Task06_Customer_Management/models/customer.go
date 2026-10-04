package models

import "time"

type Customer struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	City      string
	CreatedAt time.Time
}
