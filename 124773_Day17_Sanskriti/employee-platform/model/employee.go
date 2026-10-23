package model

type Department struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Salary float64 `json:"salary"`
	DepartmentID int64 `json:"department_id"`
	Department *Department `json:"department,omitempty"`
}

type EmployeeInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Salary float64 `json:"salary"`
	DepartmentID int64 `json:"department_id"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string
	Role string
}
