package models

import "time"

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name" binding:"required,min=2"`
	Email      string    `json:"email" binding:"required,email"`
	Age        int       `json:"age" binding:"required,min=18,max=65"`
	Department string    `json:"department" binding:"required"`
	Salary     float64   `json:"salary" binding:"min=0"`
	CreatedAt  time.Time `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
	Role     string `json:"role"` // admin, manager, user
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
}
