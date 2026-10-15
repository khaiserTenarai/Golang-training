package service

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"graphql_employee/models"
)

type EmployeeService struct {
	mu   sync.RWMutex
	data map[int]models.Employee
	next int
}

func NewEmployeeService() *EmployeeService {
	s := &EmployeeService{data: make(map[int]models.Employee), next: 1}
	seed := []models.Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, DepartmentID: 1, Salary: 60000, Role: "developer"},
		{Name: "Alice", Email: "alice@test.com", Age: 30, DepartmentID: 2, Salary: 55000, Role: "designer"},
		{Name: "Bob", Email: "bob@test.com", Age: 28, DepartmentID: 3, Salary: 45000, Role: "marketer"},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, DepartmentID: 1, Salary: 80000, Role: "lead"},
		{Name: "Diana", Email: "diana@test.com", Age: 27, DepartmentID: 4, Salary: 50000, Role: "hr"},
		{Name: "Eve", Email: "eve@test.com", Age: 29, DepartmentID: 5, Salary: 65000, Role: "analyst"},
	}
	for _, e := range seed {
		e.ID = s.next
		e.CreatedAt = time.Now()
		s.data[e.ID] = e
		s.next++
	}
	return s
}

// Task 8: Validate input
func (s *EmployeeService) Validate(name, email string, age int, salary float64) error {
	if strings.TrimSpace(name) == "" || len(name) < 2 {
		return fmt.Errorf("name must be at least 2 characters")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("invalid email format")
	}
	if age < 18 || age > 65 {
		return fmt.Errorf("age must be 18-65")
	}
	if salary < 0 {
		return fmt.Errorf("salary cannot be negative")
	}
	return nil
}

func (s *EmployeeService) Create(emp models.Employee) (models.Employee, error) {
	if err := s.Validate(emp.Name, emp.Email, emp.Age, emp.Salary); err != nil {
		return models.Employee{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, e := range s.data {
		if strings.EqualFold(e.Email, emp.Email) {
			return models.Employee{}, fmt.Errorf("email already exists")
		}
	}
	emp.ID = s.next
	s.next++
	emp.CreatedAt = time.Now()
	s.data[emp.ID] = emp
	return emp, nil
}

func (s *EmployeeService) GetByID(id int) (models.Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.data[id]
	if !ok {
		return models.Employee{}, fmt.Errorf("employee %d not found", id)
	}
	return e, nil
}

// Tasks 9, 10, 11: Pagination, filtering, sorting
func (s *EmployeeService) GetAll(filter models.EmployeeFilter, page models.PaginationInput) ([]models.Employee, int) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []models.Employee
	for _, e := range s.data {
		if filter.Name != "" && !strings.Contains(strings.ToLower(e.Name), strings.ToLower(filter.Name)) {
			continue
		}
		if filter.MinSalary > 0 && e.Salary < filter.MinSalary {
			continue
		}
		if filter.MaxSalary > 0 && e.Salary > filter.MaxSalary {
			continue
		}
		if filter.Department != "" {
			// Filter by department name would need dept service; skip here, use dept_id
			continue
		}
		list = append(list, e)
	}

	// Task 11: Sorting
	desc := strings.ToLower(filter.SortOrder) == "desc"
	sort.Slice(list, func(i, j int) bool {
		var less bool
		switch strings.ToLower(filter.SortBy) {
		case "name":
			less = list[i].Name < list[j].Name
		case "salary":
			less = list[i].Salary < list[j].Salary
		case "age":
			less = list[i].Age < list[j].Age
		default:
			less = list[i].ID < list[j].ID
		}
		if desc {
			return !less
		}
		return less
	})

	// Task 9: Pagination
	total := len(list)
	if page.Page < 1 {
		page.Page = 1
	}
	if page.PageSize < 1 {
		page.PageSize = 10
	}
	start := (page.Page - 1) * page.PageSize
	end := start + page.PageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	return list[start:end], total
}

func (s *EmployeeService) Update(id int, emp models.Employee) (models.Employee, error) {
	if err := s.Validate(emp.Name, emp.Email, emp.Age, emp.Salary); err != nil {
		return models.Employee{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	old, ok := s.data[id]
	if !ok {
		return models.Employee{}, fmt.Errorf("employee %d not found", id)
	}
	emp.ID = id
	emp.CreatedAt = old.CreatedAt
	s.data[id] = emp
	return emp, nil
}

func (s *EmployeeService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return fmt.Errorf("employee %d not found", id)
	}
	delete(s.data, id)
	return nil
}

func (s *EmployeeService) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func TotalPages(total, pageSize int) int {
	return int(math.Ceil(float64(total) / float64(pageSize)))
}
