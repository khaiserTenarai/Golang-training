package controller

import (
	"fmt"
	"task01_employee_crud/repository"
	"task01_employee_crud/view"
)

type EmployeeController struct {
	repo *repository.EmployeeRepository
	view *view.EmployeeView
}

func NewEmployeeController(repo *repository.EmployeeRepository, v *view.EmployeeView) *EmployeeController {
	return &EmployeeController{repo: repo, view: v}
}

func (c *EmployeeController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.CreateEmployee()
		case 2:
			c.ListEmployees()
		case 3:
			c.GetEmployee()
		case 4:
			c.UpdateEmployee()
		case 5:
			c.DeleteEmployee()
		case 6:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func (c *EmployeeController) CreateEmployee() {
	emp := c.view.GetEmployeeInput()
	id, err := c.repo.Create(emp)
	if err != nil {
		c.view.ShowError("creating employee", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Employee created with ID: %d", id))
}

func (c *EmployeeController) ListEmployees() {
	employees, err := c.repo.GetAll()
	if err != nil {
		c.view.ShowError("fetching employees", err)
		return
	}
	c.view.ShowEmployees(employees)
}

func (c *EmployeeController) GetEmployee() {
	id := c.view.GetID()
	emp, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching employee", err)
		return
	}
	c.view.ShowEmployee(emp)
}

func (c *EmployeeController) UpdateEmployee() {
	id := c.view.GetID()
	existing, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching employee", err)
		return
	}
	updated := c.view.GetUpdateInput(existing)
	if err := c.repo.Update(updated); err != nil {
		c.view.ShowError("updating employee", err)
		return
	}
	c.view.ShowSuccess("Employee updated successfully.")
}

func (c *EmployeeController) DeleteEmployee() {
	id := c.view.GetID()
	if err := c.repo.Delete(id); err != nil {
		c.view.ShowError("deleting employee", err)
		return
	}
	c.view.ShowSuccess("Employee deleted successfully.")
}
