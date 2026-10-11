package model

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	Department string `json:"department" binding:"required"`
	Role       string `json:"role" binding:"required"`
}