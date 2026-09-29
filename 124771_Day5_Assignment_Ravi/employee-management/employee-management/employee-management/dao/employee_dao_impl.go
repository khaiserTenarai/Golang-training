package dao

import (
	"errors"
	"strings"

	"employee-management/model"
)

type EmployeeDAOImpl struct {
	employees []model.Employee
}

func NewEmployeeDAO() *EmployeeDAOImpl {
	return &EmployeeDAOImpl{
		employees: make([]model.Employee, 0),
	}
}

func (dao *EmployeeDAOImpl) AddEmployee(employee model.Employee) error {

	for _, emp := range dao.employees {

		if emp.ID == employee.ID {
			return errors.New("employee ID already exists")
		}
	}

	dao.employees = append(dao.employees, employee)

	return nil
}

func (dao *EmployeeDAOImpl) GetEmployeeByID(id int) (model.Employee, error) {

	for _, emp := range dao.employees {

		if emp.ID == id {
			return emp, nil
		}
	}

	return model.Employee{}, errors.New("employee not found")
}

func (dao *EmployeeDAOImpl) GetAllEmployees() []model.Employee {

	return dao.employees
}

func (dao *EmployeeDAOImpl) UpdateEmployee(employee model.Employee) error {

	for i := range dao.employees {

		if dao.employees[i].ID == employee.ID {

			dao.employees[i] = employee

			return nil
		}
	}

	return errors.New("employee not found")
}

func (dao *EmployeeDAOImpl) DeleteEmployee(id int) error {

	for i := range dao.employees {

		if dao.employees[i].ID == id {

			dao.employees = append(
				dao.employees[:i],
				dao.employees[i+1:]...,
			)

			return nil
		}
	}

	return errors.New("employee not found")
}

func (dao *EmployeeDAOImpl) SearchEmployee(name string) []model.Employee {

	var result []model.Employee

	for _, emp := range dao.employees {

		if strings.Contains(
			strings.ToLower(emp.Name),
			strings.ToLower(name),
		) {
			result = append(result, emp)
		}
	}

	return result
}