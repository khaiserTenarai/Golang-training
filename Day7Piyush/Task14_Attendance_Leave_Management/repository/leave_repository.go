package repository

import (
	"database/sql"
	"task14_attendance_leave_management/models"
)

type LeaveRepository struct {
	DB *sql.DB
}

func NewLeaveRepository(db *sql.DB) *LeaveRepository {
	return &LeaveRepository{DB: db}
}

func (r *LeaveRepository) ApplyLeave(leave models.LeaveRequest) (models.LeaveRequest, error) {
	err := r.DB.QueryRow(
		`INSERT INTO leave_requests (employee_id, leave_type, start_date, end_date, reason)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, employee_id, leave_type, start_date, end_date, reason, status, created_at`,
		leave.EmployeeID, leave.LeaveType, leave.StartDate, leave.EndDate, leave.Reason).
		Scan(&leave.ID, &leave.EmployeeID, &leave.LeaveType, &leave.StartDate, &leave.EndDate,
			&leave.Reason, &leave.Status, &leave.CreatedAt)
	return leave, err
}

func (r *LeaveRepository) GetLeavesByEmployee(employeeID int) ([]models.LeaveRequest, error) {
	rows, err := r.DB.Query(
		`SELECT lr.id, lr.employee_id, e.name, lr.leave_type, lr.start_date, lr.end_date,
		lr.reason, lr.status, COALESCE(lr.reviewed_by, ''), lr.created_at
		FROM leave_requests lr
		JOIN att_employees e ON lr.employee_id = e.id
		WHERE lr.employee_id = $1 ORDER BY lr.created_at DESC`, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var leaves []models.LeaveRequest
	for rows.Next() {
		var l models.LeaveRequest
		if err := rows.Scan(&l.ID, &l.EmployeeID, &l.EmployeeName, &l.LeaveType,
			&l.StartDate, &l.EndDate, &l.Reason, &l.Status, &l.ReviewedBy, &l.CreatedAt); err != nil {
			return nil, err
		}
		leaves = append(leaves, l)
	}
	return leaves, rows.Err()
}

func (r *LeaveRepository) GetPendingLeaves() ([]models.LeaveRequest, error) {
	rows, err := r.DB.Query(
		`SELECT lr.id, lr.employee_id, e.name, lr.leave_type, lr.start_date, lr.end_date,
		lr.reason, lr.status, COALESCE(lr.reviewed_by, ''), lr.created_at
		FROM leave_requests lr
		JOIN att_employees e ON lr.employee_id = e.id
		WHERE lr.status = 'pending' ORDER BY lr.created_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var leaves []models.LeaveRequest
	for rows.Next() {
		var l models.LeaveRequest
		if err := rows.Scan(&l.ID, &l.EmployeeID, &l.EmployeeName, &l.LeaveType,
			&l.StartDate, &l.EndDate, &l.Reason, &l.Status, &l.ReviewedBy, &l.CreatedAt); err != nil {
			return nil, err
		}
		leaves = append(leaves, l)
	}
	return leaves, rows.Err()
}

func (r *LeaveRepository) ApproveLeave(leaveID int, reviewedBy string) error {
	_, err := r.DB.Exec(
		`UPDATE leave_requests SET status = 'approved', reviewed_by = $1 WHERE id = $2`,
		reviewedBy, leaveID)
	return err
}

func (r *LeaveRepository) RejectLeave(leaveID int, reviewedBy string) error {
	_, err := r.DB.Exec(
		`UPDATE leave_requests SET status = 'rejected', reviewed_by = $1 WHERE id = $2`,
		reviewedBy, leaveID)
	return err
}
