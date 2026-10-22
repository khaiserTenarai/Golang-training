package models

import "time"

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
