package service

import (
	"errors"

	"employee-management-app/model"
	"employee-management-app/repository"
	"employee-management-app/utility"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewEmployeeService(
	repository repository.EmployeeRepository,
) EmployeeService {

	return &EmployeeServiceImpl{
		repository: repository,
	}
}

// ==================================================
// ADD EMPLOYEE
// ==================================================

func (s *EmployeeServiceImpl) AddEmployee(
	employee model.Employee,
) error {

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateEmail(employee.Email); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	return s.repository.Save(employee)
}

// ==================================================
// DELETE EMPLOYEE
// ==================================================

func (s *EmployeeServiceImpl) DeleteEmployee(
	id int,
) error {

	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}

	return s.repository.Delete(id)
}

// ==================================================
// UPDATE EMPLOYEE
// ==================================================

func (s *EmployeeServiceImpl) UpdateEmployee(
	employee model.Employee,
) error {

	if employee.ID <= 0 {
		return errors.New("ID must be greater than 0")
	}

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateEmail(employee.Email); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	return s.repository.Update(employee)
}

// ==================================================
// FIND BY ID
// ==================================================

func (s *EmployeeServiceImpl) FindEmployeeByID(
	id int,
) (model.Employee, error) {

	if id <= 0 {
		return model.Employee{}, errors.New(
			"ID must be greater than 0",
		)
	}

	return s.repository.FindByID(id)
}

// ==================================================
// FIND ALL
// ==================================================

func (s *EmployeeServiceImpl) FindAllEmployees() []model.Employee {

	return s.repository.FindAll()
}
