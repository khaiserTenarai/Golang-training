package controller

import (
	"fmt"

	"Day5Piyush/employee-management/model"
	"Day5Piyush/employee-management/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(
	service service.EmployeeService,
) *EmployeeController {

	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) AddEmployee(
	employee model.Employee,
) {

	err := c.service.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully")
}

func (c *EmployeeController) GetEmployee(id int) {

	employee, err := c.service.GetEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	employee.Display()
}

func (c *EmployeeController) GetAllEmployees() {

	employees := c.service.GetEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}

	for _, employee := range employees {
		employee.Display()
		fmt.Println("-------------------------")
	}
}

func (c *EmployeeController) DeleteEmployee(id int) {

	err := c.service.DeleteEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully")
}
