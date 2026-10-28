package service

import (
	"employee-management/model"
	"employee-management/repository"
	"employee-management/utility"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(
	repo *repository.EmployeeRepository,
) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetEmployees(
	page int,
	limit int,
	department string,
	sortBy string,
	sortOrder string,
) ([]model.Employee, error) {
	return s.repo.GetEmployees(
		page,
		limit,
		department,
		sortBy,
		sortOrder,
	)
}

func (s *EmployeeService) GetEmployee(id uint) (model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) CreateEmployee(
	employee model.Employee,
) (model.Employee, error) {

	if err := utility.ValidateEmployee(employee); err != nil {
		return employee, err
	}

	return s.repo.Create(employee)
}

func (s *EmployeeService) UpdateEmployee(
	id uint,
	employee model.Employee,
) (model.Employee, error) {

	if err := utility.ValidateEmployee(employee); err != nil {
		return employee, err
	}

	return s.repo.Update(id, employee)
}

func (s *EmployeeService) DeleteEmployee(id uint) (bool, error) {
	err := s.repo.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
