package models

import "time"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Department string
	Salary     float64
	CreatedAt  time.Time
}
