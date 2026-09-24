package employee

// Employee represents an employee in the company.
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// CreateEmployee creates a new employee.
func CreateEmployee(id int, name string, salary float64) Employee {
	return Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}
}