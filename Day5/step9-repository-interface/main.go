package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

type EmployeeRepository interface {
	Save(e Employee) Employee
	FindByID(id int) (Employee, bool)
	FindAll() []Employee
}

type InMemoryEmployeeRepository struct {
	employees []Employee
	nextID    int
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{nextID: 1}
}

func (r *InMemoryEmployeeRepository) Save(e Employee) Employee {
	e.ID = r.nextID
	r.nextID++
	r.employees = append(r.employees, e)
	return e
}

func (r *InMemoryEmployeeRepository) FindByID(id int) (Employee, bool) {
	for _, e := range r.employees {
		if e.ID == id {
			return e, true
		}
	}
	return Employee{}, false
}

func (r *InMemoryEmployeeRepository) FindAll() []Employee {
	return r.employees
}

func useRepository(repo EmployeeRepository) {
	repo.Save(Employee{Name: "Ray"})
	repo.Save(Employee{Name: "Meera"})

	all := repo.FindAll()
	fmt.Println("All employees via the interface:")
	for _, e := range all {
		fmt.Printf("  [%d] %s\n", e.ID, e.Name)
	}
}

func main() {
	repo := NewInMemoryEmployeeRepository()
	useRepository(repo)
}
