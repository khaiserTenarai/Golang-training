package service

import (
	"errors"

	"employee-managementPostgre/model"
	"employee-managementPostgre/repository"
)

type EmployeeService interface {
	AddEmployee(employee *model.Employee) error
	GetEmployees() ([]model.Employee, error)
	GetEmployee(id int) (model.Employee, error)
	UpdateEmployee(employee *model.Employee) error
	DeleteEmployee(id int) error
}

type EmployeeServiceImpl struct {
	Repository repository.EmployeeRepository
}

func (s *EmployeeServiceImpl) AddEmployee(employee *model.Employee) error {

	if employee.Name == "" {
		return errors.New("employee name cannot be empty")
	}

	if employee.Email == "" {
		return errors.New("employee email cannot be empty")
	}

	if employee.Age <= 0 {
		return errors.New("employee age must be greater than 0")
	}

	if employee.Salary < 0 {
		return errors.New("employee salary cannot be negative")
	}

	if employee.DepartmentID <= 0 {
		return errors.New("department ID must be greater than 0")
	}

	return s.Repository.CreateEmployee(employee)
}

//gives all the emp data
func (s *EmployeeServiceImpl) GetEmployees() ([]model.Employee, error) {

	return s.Repository.GetAllEmployees()
}

//give emp data based on the id input
func (s *EmployeeServiceImpl) GetEmployee(id int) (model.Employee, error) {

	if id <= 0 {
		return model.Employee{}, errors.New("employee ID must be greater than 0")
	}

	return s.Repository.GetEmployeeByID(id)
}

func (s *EmployeeServiceImpl) UpdateEmployee(employee *model.Employee) error {

	if employee.ID <= 0 {
		return errors.New("employee ID must be greater than 0")
	}

	if employee.Name == "" {
		return errors.New("employee name cannot be empty")
	}

	if employee.Email == "" {
		return errors.New("employee email cannot be empty")
	}

	if employee.Age <= 0 {
		return errors.New("employee age must be greater than 0")
	}

	if employee.Salary < 0 {
		return errors.New("employee salary cannot be negative")
	}

	if employee.DepartmentID <= 0 {
		return errors.New("department ID must be greater than 0")
	}

	return s.Repository.UpdateEmployee(employee)
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) error {

	if id <= 0 {
		return errors.New("employee ID must be greater than 0")
	}

	return s.Repository.DeleteEmployee(id)
}