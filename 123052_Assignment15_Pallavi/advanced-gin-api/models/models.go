package models

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateEmployeeRequest struct {
	Name       string  `json:"name" binding:"required,min=2"`
	Role       string  `json:"role" binding:"required"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"required,gt=0"`
}

type Employee struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}