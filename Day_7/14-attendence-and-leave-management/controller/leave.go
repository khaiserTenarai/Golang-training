package controller
import (
	"fmt"

	"attendance_leave/service"
	"attendance_leave/view"
)

type LeaveController interface {
	Start()
	Process(choice int)
}

type LeaveControllerImpl struct {
	view    view.LeaveView
	service service.LeaveService
}

func NewLeaveController(
	view view.LeaveView,
	service service.LeaveService,
) LeaveController {

	return &LeaveControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *LeaveControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 6 {
			return
		}

		c.Process(choice)
	}
}

func (c *LeaveControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		leave := c.view.ReadLeave()

		err := c.service.Apply(leave)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Leave applied successfully.",
		)

	case 2:

		id := c.view.ReadID()

		err := c.service.Approve(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Leave approved successfully.",
		)

	case 3:

		id := c.view.ReadID()

		err := c.service.Reject(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Leave rejected successfully.",
		)

	case 4:

		employeeID :=
			c.view.ReadEmployeeID()

		leaves, err :=
			c.service.FindByEmployee(
				employeeID,
			)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayLeaves(leaves)

	case 5:

		leaves, err :=
			c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayLeaves(leaves)

	default:

		fmt.Println("Invalid choice.")
	}
}
