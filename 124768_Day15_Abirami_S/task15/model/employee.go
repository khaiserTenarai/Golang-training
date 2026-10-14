package model

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Department string `json:"department" binding:"required"`
	Salary     int    `json:"salary" binding:"required,min=1"`
}
