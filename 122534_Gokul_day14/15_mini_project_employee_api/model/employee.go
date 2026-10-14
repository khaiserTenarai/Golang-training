package model

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type EmployeeRequest struct {
	Name       string  `json:"name" binding:"required,min=2"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"gte=0"`
}

type EmployeePatch struct {
	Name       *string  `json:"name" binding:"omitempty,min=2"`
	Email      *string  `json:"email" binding:"omitempty,email"`
	Department *string  `json:"department"`
	Salary     *float64 `json:"salary" binding:"omitempty,gte=0"`
}
