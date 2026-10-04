package routes

import (
	"task14_attendance_leave_management/controller"

	"github.com/gorilla/mux"
)

func SetupRoutes(ctrl *controller.AttendanceController) *mux.Router {
	r := mux.NewRouter()

	// Employee
	r.HandleFunc("/api/employees", ctrl.CreateEmployee).Methods("POST")
	r.HandleFunc("/api/employees", ctrl.ListEmployees).Methods("GET")

	// Attendance
	r.HandleFunc("/api/attendance/checkin/{id}", ctrl.CheckIn).Methods("POST")
	r.HandleFunc("/api/attendance/checkout/{id}", ctrl.CheckOut).Methods("POST")
	r.HandleFunc("/api/attendance/report/{id}", ctrl.AttendanceReport).Methods("GET")

	// Leave
	r.HandleFunc("/api/leave/apply", ctrl.ApplyLeave).Methods("POST")
	r.HandleFunc("/api/leave/employee/{id}", ctrl.MyLeaves).Methods("GET")
	r.HandleFunc("/api/leave/pending", ctrl.PendingLeaves).Methods("GET")
	r.HandleFunc("/api/leave/approve/{id}", ctrl.ApproveLeave).Methods("PUT")
	r.HandleFunc("/api/leave/reject/{id}", ctrl.RejectLeave).Methods("PUT")

	return r
}
