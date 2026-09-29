// 9. Create EmployeeRepository interface.

package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID   int
	Name string
}

type EmployeeRepository interface {
	Add(emp Employee)
	Get(id int) (Employee, error)
}

type InMemoryRepository struct {
	employees []Employee
}

func (r *InMemoryRepository) Add(emp Employee) {
	r.employees = append(r.employees, emp)
}

func (r *InMemoryRepository) Get(id int) (Employee, error) {
	for _, emp := range r.employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee not found")
}

func main() {

	var repo EmployeeRepository = &InMemoryRepository{}

	repo.Add(Employee{ID: 1, Name: "Anita Sharma"})
	repo.Add(Employee{ID: 2, Name: "Ravi Kumar"})

	emp, err := repo.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", emp)
	}
}
