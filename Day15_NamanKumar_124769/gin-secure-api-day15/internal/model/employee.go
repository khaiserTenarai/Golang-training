package model

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required,min=2,max=80"`
	Email      string `json:"email" binding:"required,email"`
	Department string `json:"department" binding:"required,min=2,max=60"`
	Role       string `json:"role" binding:"required,oneof=admin manager employee"`
}
