package employee

import (
	"company/utils"
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func New(id int, name string, salary float64) Employee {
	return Employee{
		ID:     id,
		Name:   utils.FormatName(name),
		Salary: salary,
	}
}

func (e Employee) Display() {
	fmt.Println("ID:", e.ID, "Name:", e.Name, "Salary:", e.Salary)
}

func (e Employee) GiveRaise(bonus float64) Employee {
	e.Salary = e.Salary + bonus
	return e
}
