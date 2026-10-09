package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"attendance-leave-management/model"
)

type PostgresAttendanceRepository struct {
	db *pgxpool.Pool
}

func NewPostgresAttendanceRepository(
	db *pgxpool.Pool,
) *PostgresAttendanceRepository {
	return &PostgresAttendanceRepository{
		db: db,
	}
}

func (r *PostgresAttendanceRepository) GetEmployees() ([]model.Employee, error) {
	query := `
        SELECT id, name, email
        FROM employees
        ORDER BY id
    `

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var employee model.Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
		); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, rows.Err()
}

func (r *PostgresAttendanceRepository) CheckIn(
	employeeID int,
	date time.Time,
	checkIn time.Time,
) error {
	query := `
        INSERT INTO attendance
        (employee_id, attendance_date, check_in)
        VALUES ($1, $2, $3)
        ON CONFLICT (employee_id, attendance_date)
        DO UPDATE SET check_in = EXCLUDED.check_in
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		employeeID,
		date,
		checkIn,
	)

	return err
}

func (r *PostgresAttendanceRepository) CheckOut(
	employeeID int,
	date time.Time,
	checkOut time.Time,
) error {
	query := `
        UPDATE attendance
        SET check_out = $1
        WHERE employee_id = $2
          AND attendance_date = $3
    `

	result, err := r.db.Exec(
		context.Background(),
		query,
		checkOut,
		employeeID,
		date,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("employee has not checked in today")
	}

	return nil
}

func (r *PostgresAttendanceRepository) GetAttendanceReport(
	employeeID int,
	from time.Time,
	to time.Time,
) ([]model.Attendance, error) {
	query := `
        SELECT id, employee_id, attendance_date, check_in, check_out
        FROM attendance
        WHERE employee_id = $1
          AND attendance_date BETWEEN $2 AND $3
        ORDER BY attendance_date
    `

	rows, err := r.db.Query(
		context.Background(),
		query,
		employeeID,
		from,
		to,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.Attendance

	for rows.Next() {
		var attendance model.Attendance

		if err := rows.Scan(
			&attendance.ID,
			&attendance.EmployeeID,
			&attendance.AttendanceDate,
			&attendance.CheckIn,
			&attendance.CheckOut,
		); err != nil {
			return nil, err
		}

		records = append(records, attendance)
	}

	return records, rows.Err()
}

func (r *PostgresAttendanceRepository) ApplyLeave(
	leave model.Leave,
) error {
	query := `
        INSERT INTO leaves
        (employee_id, leave_date, reason, status)
        VALUES ($1, $2, $3, 'PENDING')
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		leave.EmployeeID,
		leave.LeaveDate,
		leave.Reason,
	)

	return err
}

func (r *PostgresAttendanceRepository) GetLeaves(
	employeeID int,
) ([]model.Leave, error) {
	query := `
        SELECT id, employee_id, leave_date, reason, status
        FROM leaves
        WHERE employee_id = $1
        ORDER BY leave_date
    `

	rows, err := r.db.Query(
		context.Background(),
		query,
		employeeID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaves []model.Leave

	for rows.Next() {
		var leave model.Leave

		if err := rows.Scan(
			&leave.ID,
			&leave.EmployeeID,
			&leave.LeaveDate,
			&leave.Reason,
			&leave.Status,
		); err != nil {
			return nil, err
		}

		leaves = append(leaves, leave)
	}

	return leaves, rows.Err()
}

func (r *PostgresAttendanceRepository) GetAllPendingLeaves() ([]model.Leave, error) {
	query := `
        SELECT id, employee_id, leave_date, reason, status
        FROM leaves
        WHERE status = 'PENDING'
        ORDER BY leave_date
    `

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaves []model.Leave

	for rows.Next() {
		var leave model.Leave

		if err := rows.Scan(
			&leave.ID,
			&leave.EmployeeID,
			&leave.LeaveDate,
			&leave.Reason,
			&leave.Status,
		); err != nil {
			return nil, err
		}

		leaves = append(leaves, leave)
	}

	return leaves, rows.Err()
}

func (r *PostgresAttendanceRepository) UpdateLeaveStatus(
	leaveID int,
	status string,
) error {
	query := `
        UPDATE leaves
        SET status = $1
        WHERE id = $2
          AND status = 'PENDING'
    `

	result, err := r.db.Exec(
		context.Background(),
		query,
		status,
		leaveID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf(
			"leave %d not found or already processed",
			leaveID,
		)
	}

	return nil
}

var _ = pgx.ErrNoRows
