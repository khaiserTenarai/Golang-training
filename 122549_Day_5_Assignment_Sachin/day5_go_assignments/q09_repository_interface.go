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
	Save(emp Employee) error
	FindByID(id int) (Employee, error)
}

type MemoryRepo struct {
	store map[int]Employee
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{store: make(map[int]Employee)}
}

func (m *MemoryRepo) Save(emp Employee) error {
	m.store[emp.ID] = emp
	return nil
}

func (m *MemoryRepo) FindByID(id int) (Employee, error) {
	emp, ok := m.store[id]
	if !ok {
		return Employee{}, errors.New("not found")
	}
	return emp, nil
}

func main() {
	var repo EmployeeRepository = NewMemoryRepo()

	repo.Save(Employee{ID: 1, Name: "Hannah"})

	emp, err := repo.FindByID(1)
	if err == nil {
		fmt.Println("Found in Repo:", emp)
	}
}
