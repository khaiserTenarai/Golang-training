package controller
import (
	"fmt"

	"attendance_leave/service"
	"attendance_leave/view"
)

type EmployeeController interface {
	Start()
	Process(choice int)
}

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

		if choice == 4 {
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

	default:

		fmt.Println("Invalid choice.")
	}
}
