package model

type Employee struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
	Salary     float64    `json:"salary"`
}