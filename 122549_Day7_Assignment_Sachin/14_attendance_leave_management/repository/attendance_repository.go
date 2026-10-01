package repository

import (
	"context"
	"example.com/q14-attendance-leave/model"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepository interface {
	AddEmployee(context.Context, string) error
	CheckIn(context.Context, int64) error
	CheckOut(context.Context, int64) error
	ShowAttendance(context.Context) error
	ApplyLeave(context.Context, int64, string, string) error
	UpdateLeave(context.Context, int64, string) error
	ShowLeaves(context.Context) error
}
type PostgresAttendanceRepository struct{ db *pgxpool.Pool }

func NewPostgresAttendanceRepository(db *pgxpool.Pool) *PostgresAttendanceRepository {
	return &PostgresAttendanceRepository{db: db}
}
func (r *PostgresAttendanceRepository) AddEmployee(c context.Context, n string) error {
	_, e := r.db.Exec(c, `INSERT INTO employees(name) VALUES($1)`, n)
	return e
}
func (r *PostgresAttendanceRepository) CheckIn(c context.Context, id int64) error {
	_, e := r.db.Exec(c, `INSERT INTO attendance(employee_id,check_in) VALUES($1,CURRENT_TIMESTAMP)`, id)
	return e
}
func (r *PostgresAttendanceRepository) CheckOut(c context.Context, id int64) error {
	res, e := r.db.Exec(c, `UPDATE attendance SET check_out=CURRENT_TIMESTAMP WHERE id=(SELECT id FROM attendance WHERE employee_id=$1 AND check_out IS NULL ORDER BY id DESC LIMIT 1)`, id)
	if e == nil && res.RowsAffected() == 0 {
		return fmt.Errorf("no open check-in found")
	}
	return e
}
func (r *PostgresAttendanceRepository) ShowAttendance(c context.Context) error {
	rows, e := r.db.Query(c, `SELECT a.id,a.employee_id,e.name,a.check_in::text,COALESCE(a.check_out::text,'') FROM attendance a JOIN employees e ON e.id=a.employee_id ORDER BY a.id`)
	if e != nil {
		return e
	}
	defer rows.Close()
	for rows.Next() {
		var a model.Attendance
		if e = rows.Scan(&a.ID, &a.EmployeeID, &a.EmployeeName, &a.CheckIn, &a.CheckOut); e != nil {
			return e
		}
		fmt.Printf("%d | %d | %s | %s | %s\n", a.ID, a.EmployeeID, a.EmployeeName, a.CheckIn, a.CheckOut)
	}
	return rows.Err()
}
func (r *PostgresAttendanceRepository) ApplyLeave(c context.Context, id int64, d, reason string) error {
	_, e := r.db.Exec(c, `INSERT INTO leaves(employee_id,leave_date,reason,status) VALUES($1,$2,$3,'Pending')`, id, d, reason)
	return e
}
func (r *PostgresAttendanceRepository) UpdateLeave(c context.Context, id int64, status string) error {
	if status != "Approved" && status != "Rejected" {
		return fmt.Errorf("status must be Approved or Rejected")
	}
	res, e := r.db.Exec(c, `UPDATE leaves SET status=$1 WHERE id=$2`, status, id)
	if e == nil && res.RowsAffected() == 0 {
		return fmt.Errorf("leave request not found")
	}
	return e
}
func (r *PostgresAttendanceRepository) ShowLeaves(c context.Context) error {
	rows, e := r.db.Query(c, `SELECT l.id,l.employee_id,e.name,l.leave_date::text,l.reason,l.status FROM leaves l JOIN employees e ON e.id=l.employee_id ORDER BY l.id`)
	if e != nil {
		return e
	}
	defer rows.Close()
	for rows.Next() {
		var l model.Leave
		if e = rows.Scan(&l.ID, &l.EmployeeID, &l.EmployeeName, &l.LeaveDate, &l.Reason, &l.Status); e != nil {
			return e
		}
		fmt.Printf("%d | %d | %s | %s | %s | %s\n", l.ID, l.EmployeeID, l.EmployeeName, l.LeaveDate, l.Reason, l.Status)
	}
	return rows.Err()
}
