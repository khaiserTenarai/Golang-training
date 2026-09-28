package controller

import (
	"employee-management/model"
	"employee-management/service"
)

type EmployeeControllerImpl struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		service: service,
	}
}
func (c *EmployeeControllerImpl) Add(employee model.Employee) error {
	return c.service.AddEmployee(employee)
}
func (c *EmployeeControllerImpl) GetEmployeeById(id int) (model.Employee, error) {
	return c.service.GetById(id)
}
func (c *EmployeeControllerImpl) GetEmployee() []model.Employee {
	return c.service.GetAll()
}
func (c *EmployeeControllerImpl) UpdateEmp(employee model.Employee) error {
	return c.service.UpdateEmployee(employee)
}
func (c *EmployeeControllerImpl) DeleteEmp(id int) error {
	return c.service.DeleteEmployee(id)
}
