package repository

import (
	"database/sql"
	"task14_attendance_leave_management/models"
	"time"
)

type AttendanceRepository struct {
	DB *sql.DB
}

func NewAttendanceRepository(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{DB: db}
}

func (r *AttendanceRepository) CreateEmployee(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`INSERT INTO att_employees (name, email, department) VALUES ($1, $2, $3)
		RETURNING id, name, email, department, created_at`,
		emp.Name, emp.Email, emp.Department).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.CreatedAt)
	return emp, err
}

func (r *AttendanceRepository) GetAllEmployees() ([]models.Employee, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, department, created_at FROM att_employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var emps []models.Employee
	for rows.Next() {
		var e models.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.CreatedAt); err != nil {
			return nil, err
		}
		emps = append(emps, e)
	}
	return emps, rows.Err()
}

func (r *AttendanceRepository) CheckIn(employeeID int) (models.Attendance, error) {
	now := time.Now()
	var att models.Attendance
	err := r.DB.QueryRow(
		`INSERT INTO attendance (employee_id, check_in, date, status) VALUES ($1, $2, CURRENT_DATE, 'present')
		RETURNING id, employee_id, check_in, date, status`,
		employeeID, now).Scan(&att.ID, &att.EmployeeID, &att.CheckIn, &att.Date, &att.Status)
	return att, err
}

func (r *AttendanceRepository) CheckOut(employeeID int) error {
	now := time.Now()
	_, err := r.DB.Exec(
		`UPDATE attendance SET check_out = $1 WHERE employee_id = $2 AND date = CURRENT_DATE AND check_out IS NULL`,
		now, employeeID)
	return err
}

func (r *AttendanceRepository) GetAttendanceReport(employeeID int) ([]models.Attendance, error) {
	rows, err := r.DB.Query(
		`SELECT a.id, a.employee_id, e.name, a.check_in, a.check_out, a.date, a.status
		FROM attendance a
		JOIN att_employees e ON a.employee_id = e.id
		WHERE a.employee_id = $1 ORDER BY a.date DESC`, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []models.Attendance
	for rows.Next() {
		var a models.Attendance
		if err := rows.Scan(&a.ID, &a.EmployeeID, &a.EmployeeName, &a.CheckIn, &a.CheckOut, &a.Date, &a.Status); err != nil {
			return nil, err
		}
		records = append(records, a)
	}
	return records, rows.Err()
}
