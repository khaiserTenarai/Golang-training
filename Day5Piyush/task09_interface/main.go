package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

type EmployeeRepository interface {
	Add(employee Employee)
	GetByID(id int) (Employee, bool)
	Delete(id int)
}

type InMemoryEmployeeRepository struct {
	employees map[int]Employee
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make(map[int]Employee),
	}
}

func (r *InMemoryEmployeeRepository) Add(employee Employee) {
	r.employees[employee.ID] = employee
}

func (r *InMemoryEmployeeRepository) GetByID(id int) (Employee, bool) {
	employee, exists := r.employees[id]
	return employee, exists
}

func (r *InMemoryEmployeeRepository) Delete(id int) {
	delete(r.employees, id)
}

func main() {

	repository := NewInMemoryEmployeeRepository()

	repository.Add(Employee{
		ID:   101,
		Name: "Piyush",
	})

	employee, exists := repository.GetByID(101)

	if exists {
		fmt.Println(employee.Name)
	}

	repository.Delete(101)
}