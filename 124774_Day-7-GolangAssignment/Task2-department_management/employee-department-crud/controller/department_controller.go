package controller

import (
	"fmt"

	"employee-management-app/service"
	"employee-management-app/view"
)

type DepartmentController struct {
	service service.DepartmentService
	view    *view.DepartmentView
}

// Constructor
func NewDepartmentController(
	service service.DepartmentService,
	view *view.DepartmentView,
) *DepartmentController {

	return &DepartmentController{
		service: service,
		view:    view,
	}
}

// Starts department management.
func (c *DepartmentController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.AddDepartment()

		case 2:
			c.FindDepartmentByID()

		case 3:
			c.FindAllDepartments()

		case 4:
			c.UpdateDepartment()

		case 5:
			c.DeleteDepartment()

		case 6:
			fmt.Println("Returning to main menu...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// CREATE
func (c *DepartmentController) AddDepartment() {

	department := c.view.ReadDepartment()

	err := c.service.AddDepartment(department)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Department added successfully.",
	)
}

// READ ONE
func (c *DepartmentController) FindDepartmentByID() {

	id := c.view.ReadInt(
		"Enter Department ID: ",
	)

	department, err :=
		c.service.FindDepartmentByID(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowDepartment(department)
}

// READ ALL
func (c *DepartmentController) FindAllDepartments() {

	departments :=
		c.service.FindAllDepartments()

	c.view.ShowDepartments(departments)
}

// UPDATE
func (c *DepartmentController) UpdateDepartment() {

	department :=
		c.view.ReadDepartmentForUpdate()

	err :=
		c.service.UpdateDepartment(department)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Department updated successfully.",
	)
}

// DELETE
func (c *DepartmentController) DeleteDepartment() {

	id := c.view.ReadInt(
		"Enter Department ID: ",
	)

	err :=
		c.service.DeleteDepartment(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Department deleted successfully.",
	)
}
