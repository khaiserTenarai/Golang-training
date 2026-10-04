package controller

import (
	"time"

	"attendance_leave_management/model"
	"attendance_leave_management/service"
	"github.com/jackc/pgx/v5"
)

func AddEmployee(conn *pgx.Conn, e model.Employee) error    { return service.AddEmployee(conn, e) }
func GetEmployees(conn *pgx.Conn) ([]model.Employee, error) { return service.GetEmployees(conn) }
func CheckIn(conn *pgx.Conn, id int) error                  { return service.CheckIn(conn, id) }
func CheckOut(conn *pgx.Conn, id int) error                 { return service.CheckOut(conn, id) }
func GetAttendanceReport(conn *pgx.Conn) ([]model.Attendance, error) {
	return service.GetAttendanceReport(conn)
}
func ApplyLeave(conn *pgx.Conn, l model.Leave) error  { return service.ApplyLeave(conn, l) }
func GetLeaves(conn *pgx.Conn) ([]model.Leave, error) { return service.GetLeaves(conn) }
func ApproveLeave(conn *pgx.Conn, id int) error       { return service.ApproveLeave(conn, id) }
func RejectLeave(conn *pgx.Conn, id int) error        { return service.RejectLeave(conn, id) }
func ParseDate(value string) (time.Time, error)       { return service.ParseDate(value) }
