package employee

// Employee represents an employee in the employee management system.
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// NewEmployee creates and returns a new Employee.
func NewEmployee(id int, name string, salary float64) Employee {
	return Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}
}
