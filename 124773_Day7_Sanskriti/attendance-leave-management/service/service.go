package service

import (
	"errors"
	"time"

	"attendance_leave_management/model"
	"attendance_leave_management/repository"
	"attendance_leave_management/utility"
	"github.com/jackc/pgx/v5"
)

func AddEmployee(conn *pgx.Conn, e model.Employee) error {
	e.Name = utility.CleanString(e.Name)
	e.Email = utility.CleanString(e.Email)
	e.Department = utility.CleanString(e.Department)

	if err := utility.ValidateName(e.Name); err != nil {
		return err
	}
	if err := utility.ValidateEmail(e.Email); err != nil {
		return err
	}
	if e.Department == "" {
		return errors.New("department cannot be empty")
	}

	return repository.AddEmployee(conn, e)
}

func GetEmployees(conn *pgx.Conn) ([]model.Employee, error) {
	return repository.GetEmployees(conn)
}

func CheckIn(conn *pgx.Conn, id int) error {
	if id <= 0 {
		return utility.ErrEmployeeNotFound
	}
	return repository.CheckIn(conn, id)
}

func CheckOut(conn *pgx.Conn, id int) error {
	if id <= 0 {
		return utility.ErrEmployeeNotFound
	}
	return repository.CheckOut(conn, id)
}

func GetAttendanceReport(conn *pgx.Conn) ([]model.Attendance, error) {
	return repository.GetAttendanceReport(conn)
}

func ApplyLeave(conn *pgx.Conn, l model.Leave) error {
	if l.EmployeeID <= 0 {
		return utility.ErrEmployeeNotFound
	}
	l.LeaveType = utility.CleanString(l.LeaveType)
	l.Reason = utility.CleanString(l.Reason)

	if l.LeaveType == "" {
		return errors.New("leave type cannot be empty")
	}
	if l.Reason == "" {
		return errors.New("leave reason cannot be empty")
	}
	if l.EndDate.Before(l.StartDate) {
		return errors.New("end date cannot be before start date")
	}

	return repository.ApplyLeave(conn, l)
}

func GetLeaves(conn *pgx.Conn) ([]model.Leave, error) {
	return repository.GetLeaves(conn)
}

func ApproveLeave(conn *pgx.Conn, id int) error {
	return repository.UpdateLeaveStatus(conn, id, "Approved")
}

func RejectLeave(conn *pgx.Conn, id int) error {
	return repository.UpdateLeaveStatus(conn, id, "Rejected")
}

func ParseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}
