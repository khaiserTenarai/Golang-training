
package employee

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func NewEmployee(id int, name string, salary float64) Employee {
	return Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}
}

func (e Employee) GetDetails() string {
	return "ID: " + string(rune(e.ID)) + ", Name: " + e.Name
}
