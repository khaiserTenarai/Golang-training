package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"attendance-leave/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewAttendanceRepository(
	db *pgxpool.Pool,
) AttendanceRepository {

	return &AttendanceRepositoryImpl{
		db: db,
	}
}

// Add Employee

func (r *AttendanceRepositoryImpl) AddEmployee(
	employee model.Employee,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO employees (name)
		 VALUES ($1)`,
		employee.Name,
	)

	return err
}

// Check In

func (r *AttendanceRepositoryImpl) CheckIn(
	employeeID int,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO attendance
		 (employee_id, check_in)
		 VALUES ($1, $2)`,
		employeeID,
		time.Now(),
	)

	return err
}

// Check Out

func (r *AttendanceRepositoryImpl) CheckOut(
	employeeID int,
) error {

	result, err := r.db.Exec(
		context.Background(),
		`UPDATE attendance
		 SET check_out = $1
		 WHERE employee_id = $2
		 AND check_out IS NULL`,
		time.Now(),
		employeeID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no active check-in found")
	}

	return nil
}

// Attendance Report

func (r *AttendanceRepositoryImpl) GetAttendance() []model.Attendance {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, employee_id, check_in, check_out
		 FROM attendance
		 ORDER BY id`,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer rows.Close()

	var list []model.Attendance

	for rows.Next() {

		var a model.Attendance

		err := rows.Scan(
			&a.ID,
			&a.EmployeeID,
			&a.CheckIn,
			&a.CheckOut,
		)

		if err != nil {
			return nil
		}

		list = append(list, a)
	}

	return list
}

// Apply Leave

func (r *AttendanceRepositoryImpl) ApplyLeave(
	leave model.Leave,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO leaves
		 (employee_id, reason)
		 VALUES ($1, $2)`,
		leave.EmployeeID,
		leave.Reason,
	)

	return err
}

// Approve / Reject Leave

func (r *AttendanceRepositoryImpl) UpdateLeaveStatus(
	leaveID int,
	status string,
) error {

	result, err := r.db.Exec(
		context.Background(),
		`UPDATE leaves
		 SET status = $1
		 WHERE id = $2`,
		status,
		leaveID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("leave not found")
	}

	return nil
}

// Leave Report

func (r *AttendanceRepositoryImpl) GetLeaves() []model.Leave {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, employee_id, reason, status
		 FROM leaves
		 ORDER BY id`,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer rows.Close()

	var list []model.Leave

	for rows.Next() {

		var leave model.Leave

		err := rows.Scan(
			&leave.ID,
			&leave.EmployeeID,
			&leave.Reason,
			&leave.Status,
		)

		if err != nil {
			return nil
		}

		list = append(list, leave)
	}

	return list
}

// Avoid unused pgx import issue if database driver changes
var _ = pgx.ErrNoRows
