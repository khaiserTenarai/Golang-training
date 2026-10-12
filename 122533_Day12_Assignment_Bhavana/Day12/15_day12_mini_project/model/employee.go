package model

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// ListFilter carries the optional filtering, sorting and paging
// options for the list endpoint down to the repository.
type ListFilter struct {
	Department string
	Search     string
	MinSalary  float64
	MaxSalary  float64
	SortBy     string
	Order      string
	Page       int
	Limit      int
}
