package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"day12/model"
)

type EmployeeRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &EmployeeRepositoryImpl{DB: db}
}

func (r *EmployeeRepositoryImpl) Create(
	ctx context.Context,
	e *model.Employee,
) error {
	return r.DB.QueryRow(ctx,
		`INSERT INTO employees (name, email, department, salary)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		e.Name, e.Email, e.Department, e.Salary,
	).Scan(&e.ID)
}

func (r *EmployeeRepositoryImpl) GetAll(
	ctx context.Context,
	department, name, sortBy, sortOrder string,
	page, pageSize int,
) ([]model.Employee, int, error) {
	orderBy := "id"

	switch sortBy {
	case "name":
		orderBy = "name"
	case "email":
		orderBy = "email"
	case "department":
		orderBy = "department"
	case "salary":
		orderBy = "salary"
	}

	direction := "ASC"
	if sortOrder == "desc" {
		direction = "DESC"
	}

	var total int

	err := r.DB.QueryRow(ctx,
		`SELECT COUNT(*) FROM employees
		 WHERE ($1 = '' OR department = $1)
		 AND ($2 = '' OR name ILIKE '%' || $2 || '%')`,
		department, name,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	query := fmt.Sprintf(
		`SELECT id, name, email, department, salary
		 FROM employees
		 WHERE ($1 = '' OR department = $1)
		 AND ($2 = '' OR name ILIKE '%%' || $2 || '%%')
		 ORDER BY %s %s
		 LIMIT $3 OFFSET $4`,
		orderBy, direction,
	)

	rows, err := r.DB.Query(
		ctx, query, department, name, pageSize, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	employees := []model.Employee{}

	for rows.Next() {
		var e model.Employee

		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Email,
			&e.Department,
			&e.Salary,
		); err != nil {
			return nil, 0, err
		}

		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

func (r *EmployeeRepositoryImpl) GetByID(
	ctx context.Context,
	id int64,
) (*model.Employee, error) {
	var e model.Employee

	err := r.DB.QueryRow(ctx,
		`SELECT id, name, email, department, salary
		 FROM employees WHERE id = $1`,
		id,
	).Scan(
		&e.ID,
		&e.Name,
		&e.Email,
		&e.Department,
		&e.Salary,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &e, nil
}

func (r *EmployeeRepositoryImpl) Update(
	ctx context.Context,
	id int64,
	e *model.Employee,
) (*model.Employee, error) {
	var updated model.Employee

	err := r.DB.QueryRow(ctx,
		`UPDATE employees
		 SET name = $1, email = $2, department = $3, salary = $4
		 WHERE id = $5
		 RETURNING id, name, email, department, salary`,
		e.Name, e.Email, e.Department, e.Salary, id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Email,
		&updated.Department,
		&updated.Salary,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &updated, nil
}

func (r *EmployeeRepositoryImpl) Delete(
	ctx context.Context,
	id int64,
) error {
	commandTag, err := r.DB.Exec(
		ctx,
		`DELETE FROM employees WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("employee not found")
	}

	return nil
}
