package controller

import (
	"fmt"

	"employee-management-app/service"
	"employee-management-app/view"
)

type EmployeeController struct {
	service service.EmployeeService

	view *view.EmployeeView
}

// Creates Controller.
func NewEmployeeController(
	service service.EmployeeService,
	view *view.EmployeeView,
) *EmployeeController {

	return &EmployeeController{
		service: service,
		view:    view,
	}
}

// Starts application.
func (c *EmployeeController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.AddEmployee()

		case 2:
			c.DeleteEmployee()

		case 3:
			c.UpdateEmployee()

		case 4:
			c.FindEmployeeByID()

		case 5:
			c.FindAllEmployees()

		case 6:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// Add employee.
func (c *EmployeeController) AddEmployee() {

	employee := c.view.ReadEmployee()

	err := c.service.AddEmployee(employee)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowMessage(
		"Employee added successfully.",
	)
}

// Delete employee.
func (c *EmployeeController) DeleteEmployee() {

	id := c.view.ReadInt(
		"Enter Employee ID: ",
	)

	err := c.service.DeleteEmployee(id)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowMessage(
		"Employee deleted successfully.",
	)
}

func (c *EmployeeController) UpdateEmployee() {

	employee := c.view.ReadEmployee()

	err := c.service.UpdateEmployee(employee)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowMessage(
		"Employee updated successfully.",
	)
}

// Find employee by ID.
func (c *EmployeeController) FindEmployeeByID() {

	id := c.view.ReadInt(
		"Enter Employee ID: ",
	)

	employee, err := c.service.FindEmployeeByID(id)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowEmployee(employee)
}

// Find all employees.
func (c *EmployeeController) FindAllEmployees() {

	employees := c.service.FindAllEmployees()

	c.view.ShowEmployees(employees)
}
