package controller
import (
	"fmt"

	"attendance_leave/service"
	"attendance_leave/view"
)

type AttendanceController interface {
	Start()
	Process(choice int)
}

type AttendanceControllerImpl struct {
	view    view.AttendanceView
	service service.AttendanceService
}

func NewAttendanceController(
	view view.AttendanceView,
	service service.AttendanceService,
) AttendanceController {

	return &AttendanceControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *AttendanceControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 5 {
			return
		}

		c.Process(choice)
	}
}

func (c *AttendanceControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		employeeID :=
			c.view.ReadEmployeeID()

		err := c.service.CheckIn(
			employeeID,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Employee checked in successfully.",
		)

	case 2:

		employeeID :=
			c.view.ReadEmployeeID()

		err := c.service.CheckOut(
			employeeID,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Employee checked out successfully.",
		)

	case 3:

		employeeID :=
			c.view.ReadEmployeeID()

		reports, err :=
			c.service.FindByEmployee(
				employeeID,
			)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayReports(reports)

	case 4:

		reports, err :=
			c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayReports(reports)

	default:

		fmt.Println("Invalid choice.")
	}
}
