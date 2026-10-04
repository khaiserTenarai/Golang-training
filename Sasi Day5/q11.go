package main

import (
	"errors"
	"fmt"
)

// --- model ---

type Employee struct {
	ID   int
	Name string
	Role string
}

// --- repository (DAO) ---

type EmployeeRepository interface {
	Save(emp Employee) error
	FindByID(id int) (Employee, error)
}

type InMemoryRepo struct {
	data map[int]Employee
}

func NewInMemoryRepo() EmployeeRepository {
	return &InMemoryRepo{
		data: make(map[int]Employee),
	}
}

func (r *InMemoryRepo) Save(emp Employee) error {
	if _, exists := r.data[emp.ID]; exists {
		return errors.New("employee already exists")
	}
	r.data[emp.ID] = emp
	return nil
}

func (r *InMemoryRepo) FindByID(id int) (Employee, error) {
	emp, exists := r.data[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	return emp, nil
}

// --- service ---

type EmployeeService interface {
	OnboardEmployee(id int, name, role string) error
	FetchEmployee(id int) (Employee, error)
}

type employeeServiceImpl struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{
		repo: repo,
	}
}

func (s *employeeServiceImpl) OnboardEmployee(id int, name, role string) error {
	if name == "" || role == "" {
		return errors.New("name and role cannot be empty")
	}
	
	emp := Employee{
		ID:   id,
		Name: name,
		Role: role,
	}
	
	return s.repo.Save(emp)
}

func (s *employeeServiceImpl) FetchEmployee(id int) (Employee, error) {
	return s.repo.FindByID(id)
}

// --- controller ---

type EmployeeController struct {
	service EmployeeService
}

func NewEmployeeController(service EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) HandleOnboard(id int, name, role string) {
	err := c.service.OnboardEmployee(id, name, role)
	if err != nil {
		fmt.Printf("Controller Error: Failed to onboard - %v\n", err)
		return
	}
	fmt.Printf("Controller Success: Onboarded %s\n", name)
}

func (c *EmployeeController) HandleFetch(id int) {
	emp, err := c.service.FetchEmployee(id)
	if err != nil {
		fmt.Printf("Controller Error: Failed to fetch - %v\n", err)
		return
	}
	fmt.Printf("Controller Success: Found Employee -> ID: %d, Name: %s, Role: %s\n", emp.ID, emp.Name, emp.Role)
}

// --- main ---

func main() {
	repo := NewInMemoryRepo()
	service := NewEmployeeService(repo)
	controller := NewEmployeeController(service)

	controller.HandleOnboard(101, "Sasi", "Backend Engineer")
	controller.HandleOnboard(102, "Rekha", "Product Manager")
	controller.HandleOnboard(101, "Duplicate", "Tester") 
	
	controller.HandleFetch(101)
	controller.HandleFetch(999) 
}