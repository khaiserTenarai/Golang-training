
package model

// Employee represents an employee stored in PostgreSQL.
type Employee struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// EmployeeInput contains data received when creating or updating
// an employee.
type EmployeeInput struct {
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	Salary     float64 `json:"salary" binding:"required,gt=0"`
}

// LoginInput contains the login credentials.
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
