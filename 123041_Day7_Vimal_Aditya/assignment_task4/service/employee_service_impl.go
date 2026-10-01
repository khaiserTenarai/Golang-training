package service

import (
    "errors"
    "strings"

    "example.com/employee-management/model"
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
    employee model.Employee,
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

    if strings.TrimSpace(employee.Address.City) == "" {
        return errors.New("city is required")
    }

    if strings.TrimSpace(employee.Address.State) == "" {
        return errors.New("state is required")
    }

    if strings.TrimSpace(employee.Address.Pincode) == "" {
        return errors.New("pincode is required")
    }

    return nil
}

func (s *employeeServiceImpl) AddEmployee(
    employee model.Employee,
) error {

    if err := s.validate(employee); err != nil {
        return err
    }

    return s.repository.Save(employee)
}

func (s *employeeServiceImpl) GetEmployee(
    id int64,
) (model.Employee, error) {

    if id <= 0 {
        return model.Employee{}, errors.New("invalid employee id")
    }

    return s.repository.FindByID(id)
}

func (s *employeeServiceImpl) GetAllEmployees() (
    []model.Employee,
    error,
) {

    return s.repository.FindAll()
}

func (s *employeeServiceImpl) UpdateEmployee(
    employee model.Employee,
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

func (s *employeeServiceImpl) UpdateSalary(id int64, newSalary float64) error {
    if id <= 0 {
        return errors.New("invalid employee id")
    }
    if newSalary < 0 {
        return errors.New("salary cannot be negative")
    }
    return s.repository.UpdateSalaryWithHistory(id, newSalary)
}

func (s *employeeServiceImpl) GetSalaryHistory(id int64) ([]model.SalaryHistory, error) {
    if id <= 0 {
        return nil, errors.New("invalid employee id")
    }
    return s.repository.GetSalaryHistory(id)
}