package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
}

type EmployeeRepository interface {
	Save(emp Employee)
	FindByID(id int) (Employee, bool)
	GetAll() []Employee
}

type InMemoryEmployeeRepo struct {
	data map[int]Employee
}

func NewInMemoryRepo() *InMemoryEmployeeRepo {
	return &InMemoryEmployeeRepo{
		data: make(map[int]Employee),
	}
}

func (r *InMemoryEmployeeRepo) Save(emp Employee) {
	r.data[emp.ID] = emp
}

func (r *InMemoryEmployeeRepo) FindByID(id int) (Employee, bool) {
	emp, exists := r.data[id]
	return emp, exists
}

func (r *InMemoryEmployeeRepo) GetAll() []Employee {
	var employees []Employee
	for _, emp := range r.data {
		employees = append(employees, emp)
	}
	return employees
}

func main() {
	var repo EmployeeRepository = NewInMemoryRepo()

	repo.Save(Employee{ID: 1, Name: "Ayush", Department: "Engineering"})
	repo.Save(Employee{ID: 2, Name: "Lakshmi", Department: "HR"})

	allEmployees := repo.GetAll()
	fmt.Println(allEmployees)

	emp, found := repo.FindByID(1)
	if found {
		fmt.Println(emp.Name)
	}

	_, found = repo.FindByID(99)
	fmt.Println(found)
}