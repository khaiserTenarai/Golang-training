package main

import "fmt"

// Employee structure
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// EmployeeRepository interface
type EmployeeRepository interface {
	AddEmployee(employee Employee)
	GetEmployees() []Employee
	DeleteEmployee(id int)
}

// Repository implementation
type EmployeeRepositoryImpl struct {
	employees []Employee
}

// Add employee
func (r *EmployeeRepositoryImpl) AddEmployee(employee Employee) {
	r.employees = append(r.employees, employee)
}

// Get all employees
func (r *EmployeeRepositoryImpl) GetEmployees() []Employee {
	return r.employees
}

// Delete employee
func (r *EmployeeRepositoryImpl) DeleteEmployee(id int) {
	for i, employee := range r.employees {
		if employee.ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return
		}
	}
}

// EmployeeService uses composition
type EmployeeService struct {
	repository EmployeeRepository
}

// Add employee through service
func (s EmployeeService) AddEmployee(employee Employee) {
	s.repository.AddEmployee(employee)
}

// Display employees
func (s EmployeeService) DisplayEmployees() {
	employees := s.repository.GetEmployees()

	for _, employee := range employees {
		fmt.Println(employee.ID, employee.Name, employee.Salary)
	}
}

// Delete employee
func (s EmployeeService) DeleteEmployee(id int) {
	s.repository.DeleteEmployee(id)
}

func main() {

	// Create repository
	repo := &EmployeeRepositoryImpl{}

	// Composition:
	// EmployeeService HAS-A EmployeeRepository
	service := EmployeeService{
		repository: repo,
	}

	// Add employees
	service.AddEmployee(Employee{1, "Sanskriti", 50000})
	service.AddEmployee(Employee{2, "Rahul", 60000})

	// Display
	fmt.Println("Employees:")
	service.DisplayEmployees()

	// Delete employee
	service.DeleteEmployee(1)

	fmt.Println("\nAfter deleting employee 1:")
	service.DisplayEmployees()
}
