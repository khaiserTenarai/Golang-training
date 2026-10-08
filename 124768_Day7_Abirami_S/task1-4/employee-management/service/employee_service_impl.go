package service

import (
	"employee-management/model"
	"employee-management/repository"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{
		repository: repository,
	}
}

func (s *EmployeeServiceImpl) CreateEmployee(employee model.Employee) error {
	return s.repository.CreateEmployee(employee)
}

func (s *EmployeeServiceImpl) GetEmployee(id int) (*model.Employee, error) {
	return s.repository.GetEmployee(id)
}

func (s *EmployeeServiceImpl) GetAllEmployees() ([]model.Employee, error) {
	return s.repository.GetAllEmployees()
}

func (s *EmployeeServiceImpl) UpdateEmployee(employee model.Employee) error {
	return s.repository.UpdateEmployee(employee)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) error {
	return s.repository.DeleteEmployee(id)
}

func (s *EmployeeServiceImpl) SearchEmployees(name string, departmentID int, salary float64, page int, size int, sortBy string, sortOrder string) ([]model.Employee, error) {
	return s.repository.SearchEmployees(name, departmentID, salary, page, size, sortBy, sortOrder)
}
