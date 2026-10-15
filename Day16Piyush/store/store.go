// In-memory data store used by resolvers. No PostgreSQL.

package store

import (
	"fmt"
	"sync"
	"time"
)

type Employee struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	DepartmentID int       `json:"department_id"`
	Salary       float64   `json:"salary"`
	IsActive     bool      `json:"is_active"`
	JoinDate     string    `json:"join_date"`
	CreatedAt    time.Time `json:"created_at"`
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Budget   float64 `json:"budget"`
}

type Store struct {
	mu          sync.RWMutex
	Employees   map[int]Employee
	Departments map[int]Department
	NextEmpID   int
	NextDeptID  int
}

func NewStore() *Store {
	s := &Store{
		Employees:   make(map[int]Employee),
		Departments: make(map[int]Department),
		NextEmpID:   1,
		NextDeptID:  1,
	}

	// Seed departments
	depts := []Department{
		{Name: "Engineering", Location: "Building A", Budget: 500000},
		{Name: "Design", Location: "Building B", Budget: 300000},
		{Name: "Marketing", Location: "Building C", Budget: 250000},
		{Name: "HR", Location: "Building D", Budget: 200000},
	}
	for _, d := range depts {
		d.ID = s.NextDeptID
		s.Departments[d.ID] = d
		s.NextDeptID++
	}

	// Seed employees
	emps := []Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, DepartmentID: 1, Salary: 60000, IsActive: true, JoinDate: "2023-01-15"},
		{Name: "Alice", Email: "alice@test.com", Age: 30, DepartmentID: 2, Salary: 55000, IsActive: true, JoinDate: "2022-06-01"},
		{Name: "Bob", Email: "bob@test.com", Age: 28, DepartmentID: 3, Salary: 45000, IsActive: true, JoinDate: "2023-03-20"},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, DepartmentID: 1, Salary: 80000, IsActive: false, JoinDate: "2020-11-10"},
		{Name: "Diana", Email: "diana@test.com", Age: 27, DepartmentID: 4, Salary: 50000, IsActive: true, JoinDate: "2024-01-05"},
	}
	for _, e := range emps {
		e.ID = s.NextEmpID
		e.CreatedAt = time.Now()
		s.Employees[e.ID] = e
		s.NextEmpID++
	}

	return s
}

func (s *Store) GetEmployee(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.Employees[id]
	if !ok {
		return Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return e, nil
}

func (s *Store) GetAllEmployees() []Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]Employee, 0, len(s.Employees))
	for _, e := range s.Employees {
		list = append(list, e)
	}
	return list
}

func (s *Store) CreateEmployee(e Employee) (Employee, error) {
	if e.Name == "" {
		return Employee{}, fmt.Errorf("name is required")
	}
	if e.Email == "" {
		return Employee{}, fmt.Errorf("email is required")
	}
	if e.Age < 18 || e.Age > 65 {
		return Employee{}, fmt.Errorf("age must be between 18 and 65")
	}
	if e.Salary < 0 {
		return Employee{}, fmt.Errorf("salary cannot be negative")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, ex := range s.Employees {
		if ex.Email == e.Email {
			return Employee{}, fmt.Errorf("email '%s' already exists", e.Email)
		}
	}
	e.ID = s.NextEmpID
	s.NextEmpID++
	e.CreatedAt = time.Now()
	if e.JoinDate == "" {
		e.JoinDate = time.Now().Format("2006-01-02")
	}
	s.Employees[e.ID] = e
	return e, nil
}

func (s *Store) UpdateEmployee(id int, e Employee) (Employee, error) {
	if e.Name == "" {
		return Employee{}, fmt.Errorf("name is required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	old, ok := s.Employees[id]
	if !ok {
		return Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	e.ID = id
	e.CreatedAt = old.CreatedAt
	if e.JoinDate == "" {
		e.JoinDate = old.JoinDate
	}
	s.Employees[id] = e
	return e, nil
}

func (s *Store) DeleteEmployee(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Employees[id]; !ok {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	delete(s.Employees, id)
	return nil
}

func (s *Store) GetDepartment(id int) (Department, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d, ok := s.Departments[id]
	if !ok {
		return Department{}, fmt.Errorf("department with ID %d not found", id)
	}
	return d, nil
}

func (s *Store) GetAllDepartments() []Department {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]Department, 0, len(s.Departments))
	for _, d := range s.Departments {
		list = append(list, d)
	}
	return list
}
