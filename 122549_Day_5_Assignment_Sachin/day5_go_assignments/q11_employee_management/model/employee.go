package model

import "fmt"

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func (e Employee) String() string {
	return fmt.Sprintf("[%d] %s - %s ($%.2f)", e.ID, e.Name, e.Department, e.Salary)
}
