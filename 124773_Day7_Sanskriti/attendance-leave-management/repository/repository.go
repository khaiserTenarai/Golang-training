package repository

import (
	"context"

	"attendance_leave_management/model"
	"attendance_leave_management/utility"
	"github.com/jackc/pgx/v5"
)

func AddEmployee(conn *pgx.Conn, e model.Employee) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO employees (name, email, department) VALUES ($1, $2, $3)`,
		e.Name, e.Email, e.Department)
	return err
}

func GetEmployees(conn *pgx.Conn) ([]model.Employee, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT id, name, email, department FROM employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department); err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, rows.Err()
}

func employeeExists(conn *pgx.Conn, id int) (bool, error) {
	var exists bool
	err := conn.QueryRow(context.Background(),
		`SELECT EXISTS(SELECT 1 FROM employees WHERE id = $1)`, id).Scan(&exists)
	return exists, err
}

func CheckIn(conn *pgx.Conn, employeeID int) error {
	exists, err := employeeExists(conn, employeeID)
	if err != nil {
		return err
	}
	if !exists {
		return utility.ErrEmployeeNotFound
	}

	_, err = conn.Exec(context.Background(),
		`INSERT INTO attendance (employee_id, attendance_date, check_in)
         VALUES ($1, CURRENT_DATE, CURRENT_TIMESTAMP)`,
		employeeID)
	if err != nil {
		return utility.ErrAlreadyCheckedIn
	}
	return nil
}

func CheckOut(conn *pgx.Conn, employeeID int) error {
	exists, err := employeeExists(conn, employeeID)
	if err != nil {
		return err
	}
	if !exists {
		return utility.ErrEmployeeNotFound
	}

	var checkIn, checkOut interface{}
	err = conn.QueryRow(context.Background(),
		`SELECT check_in, check_out FROM attendance
         WHERE employee_id = $1 AND attendance_date = CURRENT_DATE`,
		employeeID).Scan(&checkIn, &checkOut)
	if err != nil {
		return utility.ErrNotCheckedIn
	}
	if checkIn == nil {
		return utility.ErrNotCheckedIn
	}
	if checkOut != nil {
		return utility.ErrAlreadyCheckedOut
	}

	_, err = conn.Exec(context.Background(),
		`UPDATE attendance SET check_out = CURRENT_TIMESTAMP
         WHERE employee_id = $1 AND attendance_date = CURRENT_DATE`, employeeID)
	return err
}

func GetAttendanceReport(conn *pgx.Conn) ([]model.Attendance, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT a.id, a.employee_id, e.name, a.attendance_date,
                a.check_in, a.check_out
         FROM attendance a
         JOIN employees e ON e.id = a.employee_id
         ORDER BY a.attendance_date DESC, a.employee_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Attendance
	for rows.Next() {
		var a model.Attendance
		if err := rows.Scan(&a.ID, &a.EmployeeID, &a.EmployeeName,
			&a.AttendanceDate, &a.CheckIn, &a.CheckOut); err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, rows.Err()
}

func ApplyLeave(conn *pgx.Conn, l model.Leave) error {
	exists, err := employeeExists(conn, l.EmployeeID)
	if err != nil {
		return err
	}
	if !exists {
		return utility.ErrEmployeeNotFound
	}

	_, err = conn.Exec(context.Background(),
		`INSERT INTO leaves
         (employee_id, leave_type, start_date, end_date, reason, status)
         VALUES ($1, $2, $3, $4, $5, 'Pending')`,
		l.EmployeeID, l.LeaveType, l.StartDate, l.EndDate, l.Reason)
	return err
}

func GetLeaves(conn *pgx.Conn) ([]model.Leave, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT l.id, l.employee_id, e.name, l.leave_type,
                l.start_date, l.end_date, l.reason, l.status, l.applied_at
         FROM leaves l JOIN employees e ON e.id = l.employee_id
         ORDER BY l.id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Leave
	for rows.Next() {
		var l model.Leave
		if err := rows.Scan(&l.ID, &l.EmployeeID, &l.EmployeeName,
			&l.LeaveType, &l.StartDate, &l.EndDate, &l.Reason,
			&l.Status, &l.AppliedAt); err != nil {
			return nil, err
		}
		result = append(result, l)
	}
	return result, rows.Err()
}

func UpdateLeaveStatus(conn *pgx.Conn, id int, status string) error {
	var current string
	err := conn.QueryRow(context.Background(),
		`SELECT status FROM leaves WHERE id = $1`, id).Scan(&current)
	if err != nil {
		return utility.ErrLeaveNotFound
	}
	if current != "Pending" {
		return utility.ErrInvalidLeaveStatus
	}

	_, err = conn.Exec(context.Background(),
		`UPDATE leaves SET status = $1 WHERE id = $2`, status, id)
	return err
}
