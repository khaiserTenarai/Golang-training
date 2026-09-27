package service

import (
	"emp_simple/dao"
	"emp_simple/model"
	"emp_simple/utility"
)

type Emp_service interface {
	AddEmp(employee *model.Employee) error
	DeleteEmployee(id int) bool
	GetEmployee() []model.Employee
}

type EmpService struct{}

func (s *EmpService) AddEmp(employee *model.Employee) error {

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}
	if err := utility.ValidateEntry(employee.Name); err != nil {
		return err
	}
	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}
	dao.AddEmp(employee)

	return nil
}

func (s *EmpService) DeleteEmployee(id int) bool {
	return dao.DeleteEmployee(id)
}

func (s *EmpService) GetEmployee() []model.Employee {
	return dao.GetEmployee()
}
