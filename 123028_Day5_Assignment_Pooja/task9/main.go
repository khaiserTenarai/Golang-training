package main

type Employee struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	JobTitle  string    `json:"job_title"`
	Salary    float64   `json:"salary"`
	IsActive  bool      `json:"is_active"`
}

type EmployeeRepository interface {
	Create(emp *Employee) error
	GetByID(id int) (*Employee, error)
	GetAll() ([]Employee, error)
	Update(emp *Employee) error
	Delete(id int) error
}