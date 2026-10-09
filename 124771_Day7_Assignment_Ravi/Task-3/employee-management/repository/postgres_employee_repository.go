package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/model"
	"employee-management/view"
)

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

// Constructor
func NewPostgresEmployeeRepository(
	db *pgxpool.Pool,
) *PostgresEmployeeRepository {

	return &PostgresEmployeeRepository{
		db: db,
	}
}

func (r *PostgresEmployeeRepository) Search(
	ctx context.Context,
	req view.EmployeeSearchRequest,
) ([]model.Employee, int64, error) {

	var conditions []string
	var args []interface{}

	argIndex := 1

	// --------------------------------------------
	// Search by Name
	// --------------------------------------------

	if req.Name != "" {

		conditions = append(
			conditions,
			fmt.Sprintf(
				"name ILIKE $%d",
				argIndex,
			),
		)

		args = append(
			args,
			"%"+req.Name+"%",
		)

		argIndex++
	}

	// --------------------------------------------
	// Search by Department
	// --------------------------------------------

	if req.Department != "" {

		conditions = append(
			conditions,
			fmt.Sprintf(
				"department ILIKE $%d",
				argIndex,
			),
		)

		args = append(
			args,
			"%"+req.Department+"%",
		)

		argIndex++
	}

	// --------------------------------------------
	// Minimum Salary
	// --------------------------------------------

	if req.SalaryMin != nil {

		conditions = append(
			conditions,
			fmt.Sprintf(
				"salary >= $%d",
				argIndex,
			),
		)

		args = append(
			args,
			*req.SalaryMin,
		)

		argIndex++
	}

	// --------------------------------------------
	// Maximum Salary
	// --------------------------------------------

	if req.SalaryMax != nil {

		conditions = append(
			conditions,
			fmt.Sprintf(
				"salary <= $%d",
				argIndex,
			),
		)

		args = append(
			args,
			*req.SalaryMax,
		)

		argIndex++
	}

	// --------------------------------------------
	// WHERE Clause
	// --------------------------------------------

	whereClause := ""

	if len(conditions) > 0 {

		whereClause =
			"WHERE " +
				strings.Join(
					conditions,
					" AND ",
				)
	}

	// --------------------------------------------
	// Count Total Records
	// --------------------------------------------

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM employees
		%s
	`, whereClause)

	var total int64

	err := r.db.QueryRow(
		ctx,
		countQuery,
		args...,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	// --------------------------------------------
	// Pagination
	// --------------------------------------------

	offset :=
		(req.Page - 1) *
			req.PageSize

	// --------------------------------------------
	// Employee Search Query
	// --------------------------------------------

	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			department,
			salary
		FROM employees
		%s
		ORDER BY %s %s
		LIMIT $%d
		OFFSET $%d
	`,
		whereClause,
		req.SortBy,
		req.SortOrder,
		argIndex,
		argIndex+1,
	)

	args = append(
		args,
		req.PageSize,
		offset,
	)

	rows, err := r.db.Query(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	employees :=
		make([]model.Employee, 0)

	// --------------------------------------------
	// Read Rows
	// --------------------------------------------

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Department,
			&employee.Salary,
		)

		if err != nil {
			return nil, 0, err
		}

		employees = append(
			employees,
			employee,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}
