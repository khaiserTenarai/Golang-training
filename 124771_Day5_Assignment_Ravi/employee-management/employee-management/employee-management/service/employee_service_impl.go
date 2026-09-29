package service

import (
	"errors"
	"strings"

	"employee-management/dao"
	"employee-management/model"
	"employee-management/utility"
)

type EmployeeServiceImpl struct {
	dao dao.EmployeeDAO
}

func NewEmployeeService(employeeDAO dao.EmployeeDAO) *EmployeeServiceImpl {

	return &EmployeeServiceImpl{
		dao: employeeDAO,
	}
}

func (service *EmployeeServiceImpl) AddEmployee(
	employee model.Employee,
) error {

	if employee.ID <= 0 {
		return errors.New("employee ID must be greater than 0")
	}

	if !utility.IsValidName(employee.Name) {
		return errors.New("employee name cannot be empty")
	}

	if !utility.IsValidAge(employee.Age) {
		return errors.New("invalid employee age")
	}

	if !utility.IsValidSalary(employee.Salary) {
		return errors.New("salary cannot be negative")
	}

	if !utility.IsValidEmail(employee.Email) {
		return errors.New("invalid email")
	}

	return service.dao.AddEmployee(employee)
}

func (service *EmployeeServiceImpl) GetEmployee(
	id int,
) (model.Employee, error) {

	if id <= 0 {
		return model.Employee{}, errors.New("invalid employee ID")
	}

	return service.dao.GetEmployeeByID(id)
}

func (service *EmployeeServiceImpl) GetAllEmployees() []model.Employee {

	return service.dao.GetAllEmployees()
}

func (service *EmployeeServiceImpl) UpdateEmployee(
	employee model.Employee,
) error {

	if employee.ID <= 0 {
		return errors.New("invalid employee ID")
	}

	if !utility.IsValidName(employee.Name) {
		return errors.New("employee name cannot be empty")
	}

	if !utility.IsValidAge(employee.Age) {
		return errors.New("invalid employee age")
	}

	if !utility.IsValidSalary(employee.Salary) {
		return errors.New("salary cannot be negative")
	}

	if !utility.IsValidEmail(employee.Email) {
		return errors.New("invalid email")
	}

	return service.dao.UpdateEmployee(employee)
}

func (service *EmployeeServiceImpl) DeleteEmployee(id int) error {

	if id <= 0 {
		return errors.New("invalid employee ID")
	}

	return service.dao.DeleteEmployee(id)
}

func (service *EmployeeServiceImpl) SearchEmployee(
	name string,
) []model.Employee {

	name = strings.TrimSpace(name)

	if name == "" {
		return []model.Employee{}
	}

	return service.dao.SearchEmployee(name)
}