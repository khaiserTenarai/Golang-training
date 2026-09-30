package main

import (
	"fmt"
	"os"

	"attendance_leave/controller"
	"attendance_leave/database"
	"attendance_leave/repository"
	"attendance_leave/service"
	"attendance_leave/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	defer db.Close()

	// Employee

	employeeRepository :=
		repository.NewEmployeeRepository(db)

	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	employeeView :=
		view.NewEmployeeView()

	employeeController :=
		controller.NewEmployeeController(
			employeeView,
			employeeService,
		)

	// Attendance

	attendanceRepository :=
		repository.NewAttendanceRepository(db)

	attendanceService :=
		service.NewAttendanceService(
			attendanceRepository,
		)

	attendanceView :=
		view.NewAttendanceView()

	attendanceController :=
		controller.NewAttendanceController(
			attendanceView,
			attendanceService,
		)

	// Leave

	leaveRepository :=
		repository.NewLeaveRepository(db)

	leaveService :=
		service.NewLeaveService(
			leaveRepository,
		)

	leaveView :=
		view.NewLeaveView()

	leaveController :=
		controller.NewLeaveController(
			leaveView,
			leaveService,
		)

	// Main View

	mainView :=
		view.NewMainView()

	for {

		choice := mainView.ShowMenu()

		switch choice {

		case 1:

			employeeController.Start()

		case 2:

			attendanceController.Start()

		case 3:

			leaveController.Start()

		case 4:

			fmt.Println("Thank you..")
			return

		default:

			fmt.Println("Invalid choice.")
		}
	}
}
