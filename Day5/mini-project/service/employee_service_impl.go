package service

import (
	"fmt"

	"example.com/employee-management/model"
	"example.com/employee-management/repository"
	"example.com/employee-management/util"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repository: repository}
}

func validate(employee model.Employee) error {
	if err := util.ValidateName(employee.Name); err != nil {
		return err
	}
	if err := util.ValidateAge(employee.Age); err != nil {
		return err
	}
	if !util.IsValidEmail(employee.Email) {
		return fmt.Errorf("invalid email: %s", employee.Email)
	}
	return nil
}

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) (model.Employee, error) {
	fmt.Println("Hello from Service - Add Employee")
	if err := validate(employee); err != nil {
		return model.Employee{}, err
	}
	return s.repository.Save(employee), nil
}

func (s *EmployeeServiceImpl) GetEmployee(id int) (model.Employee, error) {
	fmt.Println("Hello from Service - Get Employee")
	employee, found := s.repository.FindByID(id)
	if !found {
		return model.Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return employee, nil
}

func (s *EmployeeServiceImpl) GetAllEmployees() []model.Employee {
	fmt.Println("Hello from Service - Get All Employees")
	return s.repository.FindAll()
}

func (s *EmployeeServiceImpl) UpdateEmployee(employee model.Employee) (model.Employee, error) {
	fmt.Println("Hello from Service - Update Employee")
	if err := validate(employee); err != nil {
		return model.Employee{}, err
	}
	updated, found := s.repository.Update(employee)
	if !found {
		return model.Employee{}, fmt.Errorf("employee with ID %d not found", employee.ID)
	}
	return updated, nil
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) error {
	fmt.Println("Hello from Service - Delete Employee")
	if !s.repository.Delete(id) {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	return nil
}
