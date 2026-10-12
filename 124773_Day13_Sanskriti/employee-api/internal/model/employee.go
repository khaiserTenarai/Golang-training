package model

import "time"

type Employee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Position  string    `json:"position"`
	CreatedAt time.Time `json:"created_at"`
}
