package controller

import (
	"fmt"

	"employee-app/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: svc}
}

func (c *EmployeeController) Register(id int, name, dept string, salary float64) {
	emp, err := c.service.CreateEmployee(id, name, dept, salary)
	if err != nil {
		fmt.Println("Controller Error:", err)
		return
	}
	fmt.Println("Registered successfully:", emp)
}

func (c *EmployeeController) Show(id int) {
	emp, err := c.service.FetchEmployee(id)
	if err != nil {
		fmt.Println("Controller Error:", err)
		return
	}
	fmt.Println("Employee Details:", emp)
}

func (c *EmployeeController) ListAll() {
	list := c.service.FetchAllEmployees()
	fmt.Println("--- All Employees ---")
	for _, emp := range list {
		fmt.Println(emp)
	}
}
