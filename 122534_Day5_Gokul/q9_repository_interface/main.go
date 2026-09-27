package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

type EmployeeRepository interface {
	Save(e Employee)
	FindByID(id int) (Employee, bool)
}

type InMemoryRepository struct {
	data map[int]Employee
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[int]Employee)}
}

func (r *InMemoryRepository) Save(e Employee) {
	r.data[e.ID] = e
}

func (r *InMemoryRepository) FindByID(id int) (Employee, bool) {
	e, ok := r.data[id]
	return e, ok
}

func main() {
	var repo EmployeeRepository = NewInMemoryRepository()

	repo.Save(Employee{ID: 1, Name: "Gokul"})

	e, found := repo.FindByID(1)
	if found {
		fmt.Println("Found:", e)
	}
}
