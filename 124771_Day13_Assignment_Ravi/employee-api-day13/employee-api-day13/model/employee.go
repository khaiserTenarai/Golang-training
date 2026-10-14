package model

import "time"

type Employee struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}

type EmployeeInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}
