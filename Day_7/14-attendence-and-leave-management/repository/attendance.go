package repository
import (
	"context"
	"errors"
	"time"

	"attendance_leave/model"

	"github.com/jackc/pgx/v5/pgxpool"
)


type AttendanceRepository interface {
	CheckIn(employeeID int) error
	CheckOut(employeeID int) error
	FindByEmployee(employeeID int) ([]model.AttendanceReport, error)
	FindAll() ([]model.AttendanceReport, error)
}

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

func (r *AttendanceRepositoryImpl) CheckIn(
	employeeID int,
) error {

	ctx := context.Background()

	today := time.Now()

	query := `
		INSERT INTO attendance
		(employee_id, attendance_date, check_in)
		VALUES ($1, CURRENT_DATE, $2)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		employeeID,
		today,
	)

	return err
}

func (r *AttendanceRepositoryImpl) CheckOut(
	employeeID int,
) error {

	ctx := context.Background()

	query := `
		UPDATE attendance
		SET check_out = $1
		WHERE employee_id = $2
		AND attendance_date = CURRENT_DATE
		AND check_in IS NOT NULL
		AND check_out IS NULL
	`

	result, err := r.db.Exec(
		ctx,
		query,
		time.Now(),
		employeeID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New(
			"check-in not found or employee already checked out",
		)
	}

	return nil
}

func (r *AttendanceRepositoryImpl) FindByEmployee(
	employeeID int,
) ([]model.AttendanceReport, error) {

	query := `
		SELECT
			a.employee_id,
			e.name,
			a.attendance_date,
			a.check_in,
			a.check_out
		FROM attendance a
		JOIN employees e
			ON a.employee_id = e.id
		WHERE a.employee_id = $1
		ORDER BY a.attendance_date DESC
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
		[]model.AttendanceReport,
		0,
	)

	for rows.Next() {

		var report model.AttendanceReport

		err := rows.Scan(
			&report.EmployeeID,
			&report.EmployeeName,
			&report.AttendanceDate,
			&report.CheckIn,
			&report.CheckOut,
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

func (r *AttendanceRepositoryImpl) FindAll() (
	[]model.AttendanceReport,
	error,
) {

	query := `
		SELECT
			a.employee_id,
			e.name,
			a.attendance_date,
			a.check_in,
			a.check_out
		FROM attendance a
		JOIN employees e
			ON a.employee_id = e.id
		ORDER BY a.attendance_date DESC
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
		[]model.AttendanceReport,
		0,
	)

	for rows.Next() {

		var report model.AttendanceReport

		err := rows.Scan(
			&report.EmployeeID,
			&report.EmployeeName,
			&report.AttendanceDate,
			&report.CheckIn,
			&report.CheckOut,
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
