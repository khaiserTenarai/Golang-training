package controller

import (
	"fmt"

	"employee-management/model"
	"employee-management/service"
)

// EmployeeController handles user interaction
type EmployeeController struct {
	service *service.EmployeeService
}

// Create controller
func NewEmployeeController(
	service *service.EmployeeService,
) *EmployeeController {

	return &EmployeeController{
		service: service,
	}
}

// Add employee
func (c *EmployeeController) AddEmployee(employee model.Employee) {

	c.service.AddEmployee(employee)

	fmt.Println("Employee added successfully")
}

// View employee
func (c *EmployeeController) ViewEmployee(id int) {

	employee, err := c.service.GetEmployee(id)

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nEmployee Details")
	fmt.Println("----------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Salary:", employee.Salary)
}

// Delete employee
func (c *EmployeeController) DeleteEmployee(id int) {

	err := c.service.DeleteEmployee(id)

	if err != nil {

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully")
}

// View all employees
func (c *EmployeeController) ViewAllEmployees() {

	employees := c.service.GetAllEmployees()

	fmt.Println("\nAll Employees")
	fmt.Println("-------------")

	for _, employee := range employees {

		fmt.Println(
			employee.ID,
			employee.Name,
			employee.Department,
		)
	}
}