package models

type Employee struct {
	ID       string  `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Role     string  `json:"role" binding:"required"`
	Department string `json:"department"`
	Salary   float64 `json:"salary"`
}

type UpdateEmployee struct {
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}