package models

type Employee struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Role       string  `json:"role"` // e.g., "admin", "user"
	Salary     float64 `json:"salary"`
}

type QueryParams struct {
	Page       int
	Limit      int
	Department string
	SortBy     string
	Order      string
}