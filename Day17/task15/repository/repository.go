package repository

import (
	"errors"
	"sort"

	"employee-platform/models"
)

var ErrNotFound = errors.New("resource not found")

type Repository interface {
	GetEmployees(filter models.EmployeeFilter, page models.Pagination, sortBy string, sortOrder string) ([]models.Employee, int, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	CreateEmployee(emp *models.Employee) error
	GetDepartmentByID(id string) (*models.Department, error)
	GetDepartmentsByIDs(ids []string) (map[string]models.Department, error)
}

type memoryRepo struct {
	employees   map[string]models.Employee
	departments map[string]models.Department
}

func NewMemoryRepository() Repository {
	depts := map[string]models.Department{
		"dept-1": {ID: "dept-1", Name: "Engineering", Location: "Building A"},
		"dept-2": {ID: "dept-2", Name: "Human Resources", Location: "Building B"},
	}

	emps := map[string]models.Employee{
		"emp-1": {ID: "emp-1", Name: "Alice Smith", Role: "Software Engineer", DepartmentID: "dept-1", Salary: 95000},
		"emp-2": {ID: "emp-2", Name: "Bob Jones", Role: "HR Specialist", DepartmentID: "dept-2", Salary: 65000},
		"emp-3": {ID: "emp-3", Name: "Charlie Brown", Role: "DevOps Engineer", DepartmentID: "dept-1", Salary: 105000},
	}

	return &memoryRepo{employees: emps, departments: depts}
}

func (r *memoryRepo) GetEmployees(filter models.EmployeeFilter, page models.Pagination, sortBy string, sortOrder string) ([]models.Employee, int, error) {
	var list []models.Employee

	for _, emp := range r.employees {
		if filter.DepartmentID != nil && emp.DepartmentID != *filter.DepartmentID {
			continue
		}
		if filter.MinSalary != nil && emp.Salary < *filter.MinSalary {
			continue
		}
		list = append(list, emp)
	}

	total := len(list)

	// Sorting
	sort.Slice(list, func(i, j int) bool {
		if sortBy == "salary" {
			if sortOrder == "desc" {
				return list[i].Salary > list[j].Salary
			}
			return list[i].Salary < list[j].Salary
		}
		if sortOrder == "desc" {
			return list[i].Name > list[j].Name
		}
		return list[i].Name < list[j].Name
	})

	// Pagination
	start := page.Offset
	if start > total {
		start = total
	}
	end := start + page.Limit
	if end > total || page.Limit == 0 {
		end = total
	}

	return list[start:end], total, nil
}

func (r *memoryRepo) GetEmployeeByID(id string) (*models.Employee, error) {
	emp, ok := r.employees[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &emp, nil
}

func (r *memoryRepo) CreateEmployee(emp *models.Employee) error {
	r.employees[emp.ID] = *emp
	return nil
}

func (r *memoryRepo) GetDepartmentByID(id string) (*models.Department, error) {
	dept, ok := r.departments[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &dept, nil
}

// Batch loader optimization for nested Employee -> Department queries
func (r *memoryRepo) GetDepartmentsByIDs(ids []string) (map[string]models.Department, error) {
	result := make(map[string]models.Department)
	for _, id := range ids {
		if dept, ok := r.departments[id]; ok {
			result[id] = dept
		}
	}
	return result, nil
}
