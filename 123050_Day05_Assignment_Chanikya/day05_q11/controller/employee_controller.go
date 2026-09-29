package controller

import (
	"fmt"

	"employee-management/model"
	"employee-management/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) AddEmployee(employee model.Employee) {
	c.service.AddEmployee(employee)
	fmt.Println("Employee added")
}

func (c *EmployeeController) GetEmployee(id int) {
	employee, found := c.service.GetEmployee(id)

	if found {
		fmt.Println(employee)
	} else {
		fmt.Println("Employee not found")
	}
}

func (c *EmployeeController) GetAllEmployees() {
	employees := c.service.GetAllEmployees()

	for _, employee := range employees {
		fmt.Println(employee)
	}
}

func (c *EmployeeController) DeleteEmployee(id int) {
	c.service.DeleteEmployee(id)
	fmt.Println("Employee deleted")
}
