package model

type Employee struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" binding:"required,min=2,max=100"`
	Email       string  `json:"email" binding:"required,email"`
	Department  string  `json:"department" binding:"required"`
	Designation string  `json:"designation" binding:"required"`
	Salary      float64 `json:"salary" binding:"required,gt=0"`
	AvatarPath  string  `json:"avatar_path,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"` // ADMIN or USER
}