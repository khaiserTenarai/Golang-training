package controller
import (
	"fmt"

	"employee-management/service"
	"employee-management/view"
)

type EmployeeControllerImpl struct {
	view    view.EmployeeView
	service service.EmployeeService
}

func NewEmployeeController(
	view view.EmployeeView,
	service service.EmployeeService,
) EmployeeController {

	return &EmployeeControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *EmployeeControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 7 {
			return
		}

		c.Process(choice)
	}
}

func (c *EmployeeControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		employee := c.view.ReadEmployee()

		err := c.service.Save(employee)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Employee saved successfully.")

	case 2:

		id := c.view.ReadID()

		employee, err :=
			c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayEmployee(employee)

	case 3:

		employees, err :=
			c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayEmployees(employees)

	case 4:

		employee := c.view.ReadEmployee()

		err := c.service.Update(employee)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Employee updated successfully.")

	case 5:

		id := c.view.ReadID()

		err := c.service.Delete(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Employee deleted successfully.")

	case 6:

		search := c.view.ReadSearch()

		employees, total, err :=
			c.service.Search(search)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplaySearchResult(
			employees,
			total,
			search,
		)

	default:
		fmt.Println("Invalid choice.")
	}
}
