// Models — shared data structures used across all layers.
// Covers: Tasks 1-4 (part of layered architecture)

package models

import "time"

// ---- Employee ----

type Employee struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ---- API Response Types ----

type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	RequestID  string      `json:"request_id,omitempty"`
}

// ---- Query Parameters (Tasks 10, 11, 12) ----

type QueryParams struct {
	// Task 11: Filtering
	Name       string
	Department string
	MinSalary  float64
	MaxSalary  float64

	// Task 12: Sorting
	SortBy    string // name, salary, age, department
	SortOrder string // asc, desc

	// Task 10: Pagination
	Page     int
	PageSize int
}

// ---- Validation Error (Task 9) ----

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
