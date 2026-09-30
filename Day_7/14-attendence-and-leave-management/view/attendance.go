package view

import (
	"fmt"

	"attendance_leave/model"
)

type AttendanceView interface {
	ShowMenu() int
	ReadEmployeeID() int
	DisplayReports(reports []model.AttendanceReport)
}

type AttendanceViewImpl struct {
}

func NewAttendanceView() AttendanceView {
	return &AttendanceViewImpl{}
}

func (v *AttendanceViewImpl) ShowMenu() int {

	fmt.Println("\n========== Attendance Management ==========")
	fmt.Println("1. Check In")
	fmt.Println("2. Check Out")
	fmt.Println("3. Employee Attendance Report")
	fmt.Println("4. All Attendance Report")
	fmt.Println("5. Back")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *AttendanceViewImpl) ReadEmployeeID() int {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	return id
}

func (v *AttendanceViewImpl) DisplayReports(
	reports []model.AttendanceReport,
) {

	if len(reports) == 0 {
		fmt.Println("No attendance records found.")
		return
	}

	fmt.Println("\n---------- Attendance Report ----------")

	for _, report := range reports {

		fmt.Println(
			"Employee ID   :",
			report.EmployeeID,
		)

		fmt.Println(
			"Employee Name :",
			report.EmployeeName,
		)

		fmt.Println(
			"Date          :",
			report.AttendanceDate.Format("2006-01-02"),
		)

		if report.CheckIn != nil {

			fmt.Println(
				"Check In      :",
				report.CheckIn.Format(
					"2006-01-02 15:04:05",
				),
			)

		} else {

			fmt.Println("Check In      : -")
		}

		if report.CheckOut != nil {

			fmt.Println(
				"Check Out     :",
				report.CheckOut.Format(
					"2006-01-02 15:04:05",
				),
			)

		} else {

			fmt.Println("Check Out     : -")
		}

		fmt.Println("---------------------------------------")
	}
}
