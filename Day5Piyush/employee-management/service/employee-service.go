package service

import (
	"errors"
	"strings"

	"Day5Piyush/employee-management/model"
	"Day5Piyush/employee-management/repository"
)

type employeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	repository repository.EmployeeRepository,
) EmployeeService {

	return &employeeService{
		repository: repository,
	}
}

func (s *employeeService) AddEmployee(
	employee model.Employee,
) error {

	employee.Name = strings.TrimSpace(employee.Name)
	employee.Email = strings.TrimSpace(employee.Email)

	if employee.Name == "" {
		return errors.New("employee name cannot be empty")
	}

	if employee.Email == "" {
		return errors.New("employee email cannot be empty")
	}

	if employee.Age <= 0 {
		return errors.New("employee age must be greater than zero")
	}

	if employee.Salary < 0 {
		return errors.New("employee salary cannot be negative")
	}

	return s.repository.Add(employee)
}

func (s *employeeService) GetEmployee(
	id int,
) (model.Employee, error) {

	return s.repository.GetByID(id)
}

func (s *employeeService) GetEmployees() []model.Employee {

	return s.repository.GetAll()
}

func (s *employeeService) DeleteEmployee(id int) error {

	return s.repository.Delete(id)
}
