package main

import (
	"fmt"
	"log"
	"net/http"
	"task14_attendance_leave_management/config"
	"task14_attendance_leave_management/controller"
	"task14_attendance_leave_management/repository"
	"task14_attendance_leave_management/routes"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTables(db)

	attRepo := repository.NewAttendanceRepository(db)
	leaveRepo := repository.NewLeaveRepository(db)
	ctrl := controller.NewAttendanceController(attRepo, leaveRepo)
	router := routes.SetupRoutes(ctrl)

	fmt.Println("Attendance & Leave Management Server on http://localhost:8083")
	fmt.Println("Endpoints:")
	fmt.Println("  POST /api/employees                  - Add employee")
	fmt.Println("  GET  /api/employees                  - List employees")
	fmt.Println("  POST /api/attendance/checkin/{id}     - Check in")
	fmt.Println("  POST /api/attendance/checkout/{id}    - Check out")
	fmt.Println("  GET  /api/attendance/report/{id}      - Attendance report")
	fmt.Println("  POST /api/leave/apply                 - Apply for leave")
	fmt.Println("  GET  /api/leave/employee/{id}         - My leaves")
	fmt.Println("  GET  /api/leave/pending               - Pending leaves")
	fmt.Println("  PUT  /api/leave/approve/{id}          - Approve leave")
	fmt.Println("  PUT  /api/leave/reject/{id}           - Reject leave")
	log.Fatal(http.ListenAndServe(":8083", router))
}
