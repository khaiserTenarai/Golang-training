package service

import (
    "errors"
    "strings"

    "example.com/employee-management/models"
    "example.com/employee-management/repository"
)

type employeeServiceImpl struct {
    repository repository.EmployeeRepository
}

func NewEmployeeService(
    repository repository.EmployeeRepository,
) EmployeeService {

    return &employeeServiceImpl{
        repository: repository,
    }
}

func (s *employeeServiceImpl) validate(
    employee models.Employee,
) error {

    if strings.TrimSpace(employee.Name) == "" {
        return errors.New("employee name is required")
    }

    if strings.TrimSpace(employee.Email) == "" {
        return errors.New("employee email is required")
    }

    if employee.Age < 18 {
        return errors.New("employee must be at least 18 years old")
    }

    if employee.Salary < 0 {
        return errors.New("salary cannot be negative")
    }

    if strings.TrimSpace(employee.City) == "" {
        return errors.New("city is required")
    }

    if strings.TrimSpace(employee.State) == "" {
        return errors.New("state is required")
    }

    if strings.TrimSpace(employee.Pincode) == "" {
        return errors.New("pincode is required")
    }

    return nil
}

func (s *employeeServiceImpl) AddEmployee(
    employee models.Employee,
) error {

    if err := s.validate(employee); err != nil {
        return err
    }

    return s.repository.Save(employee)
}

func (s *employeeServiceImpl) GetEmployee(
    id int64,
) (models.Employee, error) {

    if id <= 0 {
        return models.Employee{}, errors.New("invalid employee id")
    }

    return s.repository.FindByID(id)
}

func (s *employeeServiceImpl) GetAllEmployees() (
    []models.Employee,
    error,
) {

    return s.repository.FindAll()
}

func (s *employeeServiceImpl) UpdateEmployee(
    employee models.Employee,
) error {

    if employee.ID <= 0 {
        return errors.New("invalid employee id")
    }

    if err := s.validate(employee); err != nil {
        return err
    }

    return s.repository.Update(employee)
}

func (s *employeeServiceImpl) DeleteEmployee(
    id int64,
) error {

    if id <= 0 {
        return errors.New("invalid employee id")
    }

    return s.repository.Delete(id)
}
func (s *employeeServiceImpl) AssignDepartment(employeeID int64, departmentID int) error {
	if employeeID <= 0 {
		return errors.New("invalid employee id")
	}
	if departmentID <= 0 {
		return errors.New("invalid department id")
	}
	return s.repository.AssignDepartment(employeeID, departmentID)
}

func (s *employeeServiceImpl) GetEmployeeWithDepartment(id int64) (models.EmployeeWithDepartment, error) {
	if id <= 0 {
		return models.EmployeeWithDepartment{}, errors.New("invalid employee id")
	}
	return s.repository.FindWithDepartment(id)
}
