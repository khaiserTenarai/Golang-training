package repository

import (
	"errors"
	"sync"

	"employee-api/models"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository interface {
	GetAll(department string) ([]models.Employee, error)
	GetByID(id int) (*models.Employee, error)
	Create(emp *models.Employee) error
	Update(id int, emp *models.Employee) error
	Patch(id int, update *models.UpdateEmployee) (*models.Employee, error)
	Delete(id int) error
}

type memoryRepository struct {
	mu        sync.RWMutex
	employees map[int]models.Employee
	nextID    int
}

func NewMemoryRepository() EmployeeRepository {
	repo := &memoryRepository{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}

	repo.employees[1] = models.Employee{
		ID:         1,
		Name:       "Alice Smith",
		Role:       "Backend Engineer",
		Department: "Engineering",
		Salary:     95000,
	}
	repo.nextID = 2

	return repo
}

func (r *memoryRepository) GetAll(department string) ([]models.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]models.Employee, 0)
	for _, emp := range r.employees {
		if department == "" || emp.Department == department {
			list = append(list, emp)
		}
	}
	return list, nil
}

func (r *memoryRepository) GetByID(id int) (*models.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	emp, exists := r.employees[id]
	if !exists {
		return nil, ErrNotFound
	}
	return &emp, nil
}

func (r *memoryRepository) Create(emp *models.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	emp.ID = r.nextID
	r.nextID++
	r.employees[emp.ID] = *emp
	return nil
}

func (r *memoryRepository) Update(id int, emp *models.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employees[id]; !exists {
		return ErrNotFound
	}

	emp.ID = id
	r.employees[id] = *emp
	return nil
}

func (r *memoryRepository) Patch(id int, update *models.UpdateEmployee) (*models.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	emp, exists := r.employees[id]
	if !exists {
		return nil, ErrNotFound
	}

	if update.Name != nil {
		emp.Name = *update.Name
	}
	if update.Role != nil {
		emp.Role = *update.Role
	}
	if update.Department != nil {
		emp.Department = *update.Department
	}
	if update.Salary != nil {
		emp.Salary = *update.Salary
	}

	r.employees[id] = emp
	return &emp, nil
}

func (r *memoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employees[id]; !exists {
		return ErrNotFound
	}

	delete(r.employees, id)
	return nil
}
