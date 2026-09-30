package repository
import (
	"context"
	"errors"
	// "time"

	"attendance_leave/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LeaveRepository interface {
	Apply(leave model.Leave) error
	Approve(id int) error
	Reject(id int) error
	FindByEmployee(employeeID int) ([]model.LeaveReport, error)
	FindAll() ([]model.LeaveReport, error)
}

type LeaveRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewLeaveRepository(
	db *pgxpool.Pool,
) LeaveRepository {

	return &LeaveRepositoryImpl{
		db: db,
	}
}

func (r *LeaveRepositoryImpl) Apply(
	leave model.Leave,
) error {

	query := `
		INSERT INTO leaves
		(employee_id, from_date, to_date, reason)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		leave.EmployeeID,
		leave.FromDate,
		leave.ToDate,
		leave.Reason,
	)

	return err
}

func (r *LeaveRepositoryImpl) Approve(
	id int,
) error {

	query := `
		UPDATE leaves
		SET status = 'APPROVED'
		WHERE id = $1
		AND status = 'PENDING'
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New(
			"leave not found or already processed",
		)
	}

	return nil
}

func (r *LeaveRepositoryImpl) Reject(
	id int,
) error {

	query := `
		UPDATE leaves
		SET status = 'REJECTED'
		WHERE id = $1
		AND status = 'PENDING'
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New(
			"leave not found or already processed",
		)
	}

	return nil
}

func (r *LeaveRepositoryImpl) FindByEmployee(
	employeeID int,
) ([]model.LeaveReport, error) {

	query := `
		SELECT
			l.id,
			l.employee_id,
			e.name,
			l.from_date,
			l.to_date,
			l.reason,
			l.status
		FROM leaves l
		JOIN employees e
			ON l.employee_id = e.id
		WHERE l.employee_id = $1
		ORDER BY l.id DESC
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

	reports := make(
		[]model.LeaveReport,
		0,
	)

	for rows.Next() {

		var report model.LeaveReport

		err := rows.Scan(
			&report.ID,
			&report.EmployeeID,
			&report.EmployeeName,
			&report.FromDate,
			&report.ToDate,
			&report.Reason,
			&report.Status,
		)

		if err != nil {
			return nil, err
		}

		reports = append(
			reports,
			report,
		)
	}

	return reports, rows.Err()
}

func (r *LeaveRepositoryImpl) FindAll() (
	[]model.LeaveReport,
	error,
) {

	query := `
		SELECT
			l.id,
			l.employee_id,
			e.name,
			l.from_date,
			l.to_date,
			l.reason,
			l.status
		FROM leaves l
		JOIN employees e
			ON l.employee_id = e.id
		ORDER BY l.id DESC
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reports := make(
		[]model.LeaveReport,
		0,
	)

	for rows.Next() {

		var report model.LeaveReport

		err := rows.Scan(
			&report.ID,
			&report.EmployeeID,
			&report.EmployeeName,
			&report.FromDate,
			&report.ToDate,
			&report.Reason,
			&report.Status,
		)

		if err != nil {
			return nil, err
		}

		reports = append(
			reports,
			report,
		)
	}

	return reports, rows.Err()
}
