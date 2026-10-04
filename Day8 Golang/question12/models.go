package main

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type PaginatedResponse struct {
	Data       []User `json:"data"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
}