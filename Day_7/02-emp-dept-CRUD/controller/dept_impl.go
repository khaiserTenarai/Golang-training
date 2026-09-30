package controller

import (
	"fmt"

	"department-management/service"
	"department-management/view"
)

type DepartmentControllerImpl struct {
	view    view.DepartmentView
	service service.DepartmentService
}

func NewDepartmentController(
	view view.DepartmentView,
	service service.DepartmentService,
) DepartmentController {

	return &DepartmentControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *DepartmentControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 6 {
			return
		}

		c.Process(choice)
	}
}

func (c *DepartmentControllerImpl) Process(choice int) {

	switch choice {

	case 1:

		department := c.view.ReadDepartment()

		err := c.service.Save(department)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Department saved successfully.")

	case 2:

		id := c.view.ReadID()

		department, err := c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayDepartment(department)

	case 3:

		departments, err := c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayDepartments(departments)

	case 4:

		department := c.view.ReadDepartmentForUpdate()

		err := c.service.Update(department)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Department updated successfully.")

	case 5:

		id := c.view.ReadID()

		err := c.service.Delete(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Department deleted successfully.")

	default:

		fmt.Println("Invalid choice.")
	}
}
