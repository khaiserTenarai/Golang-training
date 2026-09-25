package employee

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func New(id int, name string, salary float64) Employee {
	return Employee{ID: id, Name: name, Salary: salary}
}
