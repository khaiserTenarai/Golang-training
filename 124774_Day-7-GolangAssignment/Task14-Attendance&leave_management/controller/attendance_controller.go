package controller

import (
	"fmt"

	"attendance-leave/model"
	"attendance-leave/service"
	"attendance-leave/view"
)

type AttendanceController struct {
	service service.AttendanceService
	view    *view.AttendanceView
}

func NewAttendanceController(
	service service.AttendanceService,
	view *view.AttendanceView,
) *AttendanceController {

	return &AttendanceController{
		service: service,
		view:    view,
	}
}

func (c *AttendanceController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.AddEmployee()

		case 2:
			c.CheckIn()

		case 3:
			c.CheckOut()

		case 4:
			c.AttendanceReport()

		case 5:
			c.ApplyLeave()

		case 6:
			c.ApproveLeave()

		case 7:
			c.RejectLeave()

		case 8:
			c.LeaveReport()

		case 9:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *AttendanceController) AddEmployee() {

	name := c.view.ReadString(
		"Enter employee name: ",
	)

	employee := model.Employee{
		Name: name,
	}

	err := c.service.AddEmployee(employee)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Employee added successfully.",
	)
}

func (c *AttendanceController) CheckIn() {

	id := c.view.ReadInt(
		"Enter employee ID: ",
	)

	err := c.service.CheckIn(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Employee checked in successfully.",
	)
}

func (c *AttendanceController) CheckOut() {

	id := c.view.ReadInt(
		"Enter employee ID: ",
	)

	err := c.service.CheckOut(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Employee checked out successfully.",
	)
}

func (c *AttendanceController) AttendanceReport() {

	list := c.service.GetAttendance()

	c.view.ShowAttendance(list)
}

func (c *AttendanceController) ApplyLeave() {

	id := c.view.ReadInt(
		"Enter employee ID: ",
	)

	reason := c.view.ReadString(
		"Enter leave reason: ",
	)

	leave := model.Leave{
		EmployeeID: id,
		Reason:     reason,
	}

	err := c.service.ApplyLeave(leave)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Leave applied successfully.",
	)
}

func (c *AttendanceController) ApproveLeave() {

	id := c.view.ReadInt(
		"Enter leave ID: ",
	)

	err := c.service.ApproveLeave(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Leave approved successfully.",
	)
}

func (c *AttendanceController) RejectLeave() {

	id := c.view.ReadInt(
		"Enter leave ID: ",
	)

	err := c.service.RejectLeave(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Leave rejected successfully.",
	)
}

func (c *AttendanceController) LeaveReport() {

	list := c.service.GetLeaves()

	c.view.ShowLeaves(list)
}
