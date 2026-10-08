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

func (c *EmployeeControllerImpl) CreateEmployee(employee model.Employee) error {
	return c.service.CreateEmployee(employee)
}

func (c *EmployeeControllerImpl) GetEmployee(id int) (*model.Employee, error) {
	return c.service.GetEmployee(id)
}

func (c *EmployeeControllerImpl) GetAllEmployees() ([]model.Employee, error) {
	return c.service.GetAllEmployees()
}

func (c *EmployeeControllerImpl) UpdateEmployee(employee model.Employee) error {
	return c.service.UpdateEmployee(employee)
}

func (c *EmployeeControllerImpl) DeleteEmployee(id int) error {
	return c.service.DeleteEmployee(id)
}

func (c *EmployeeControllerImpl) SearchEmployees(name string, departmentID int, salary float64, page int, size int, sortBy string, sortOrder string) ([]model.Employee, error) {
	return c.service.SearchEmployees(name, departmentID, salary, page, size, sortBy, sortOrder)
}
