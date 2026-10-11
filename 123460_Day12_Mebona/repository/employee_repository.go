package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/model"
)

type EmployeeRepository interface {
	Create(ctx context.Context, emp *model.Employee) error
	GetAll(ctx context.Context, filter model.EmployeeFilter) ([]model.Employee, int, error)
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	Ping(ctx context.Context) error
}

type postgresEmployeeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &postgresEmployeeRepository{db: db}
}

func (r *postgresEmployeeRepository) Ping(ctx context.Context) error {
	return r.db.Ping(ctx)
}

func (r *postgresEmployeeRepository) Create(ctx context.Context, emp *model.Employee) error {
	query := `INSERT INTO employees (name, email, department, salary) 
	          VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(ctx, query, emp.Name, emp.Email, emp.Department, emp.Salary).Scan(&emp.ID)
}

func (r *postgresEmployeeRepository) GetAll(ctx context.Context, filter model.EmployeeFilter) ([]model.Employee, int, error) {
	var conditions []string
	var args []interface{}
	argIdx := 1

	if filter.Department != "" {
		conditions = append(conditions, fmt.Sprintf("department = $%d", argIdx))
		args = append(args, filter.Department)
		argIdx++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM employees %s", whereClause)
	var total int
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sortBy := "id"
	if filter.SortBy == "name" || filter.SortBy == "department" || filter.SortBy == "salary" {
		sortBy = filter.SortBy
	}
	order := "ASC"
	if strings.ToUpper(filter.Order) == "DESC" {
		order = "DESC"
	}

	offset := (filter.Page - 1) * filter.Limit
	query := fmt.Sprintf("SELECT id, name, email, department, salary FROM employees %s ORDER BY %s %s LIMIT $%d OFFSET $%d",
		whereClause, sortBy, order, argIdx, argIdx+1)
	args = append(args, filter.Limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, 0, err
		}
		employees = append(employees, e)
	}

	return employees, total, nil
}

func (r *postgresEmployeeRepository) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	query := `SELECT id, name, email, department, salary FROM employees WHERE id = $1`
	var e model.Employee
	err := r.db.QueryRow(ctx, query, id).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if err != nil {
		return nil, err
	}
	return &e, nil
}