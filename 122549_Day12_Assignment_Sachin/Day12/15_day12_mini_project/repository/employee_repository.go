package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"

	"employeeapi/model"
)

var (
	ErrNotFound       = errors.New("employee not found")
	ErrDuplicateEmail = errors.New("an employee with that email already exists")
)

type EmployeeRepository interface {
	Create(e model.Employee) (model.Employee, error)
	GetByID(id int) (model.Employee, error)
	List(f model.ListFilter) ([]model.Employee, int, error)
	Update(e model.Employee) error
	Delete(id int) error
	Ping() error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Init() error {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		department VARCHAR(100) NOT NULL,
		salary NUMERIC(12,2) NOT NULL
	)`)
	return err
}

func (r *PostgresRepository) Ping() error {
	return r.db.Ping()
}

// 23505 is Postgres' code for "unique constraint violated"
func isDuplicate(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == "23505"
}

func (r *PostgresRepository) Create(e model.Employee) (model.Employee, error) {
	err := r.db.QueryRow(
		`INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`,
		e.Name, e.Email, e.Department, e.Salary,
	).Scan(&e.ID)
	if isDuplicate(err) {
		return model.Employee{}, ErrDuplicateEmail
	}
	if err != nil {
		return model.Employee{}, fmt.Errorf("create employee: %w", err)
	}
	return e, nil
}

func (r *PostgresRepository) GetByID(id int) (model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(
		`SELECT id, name, email, department, salary FROM employees WHERE id = $1`, id,
	).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Employee{}, ErrNotFound
	}
	if err != nil {
		return model.Employee{}, fmt.Errorf("get employee: %w", err)
	}
	return e, nil
}

// only these columns may be used for sorting - ORDER BY can't take a
// placeholder, so we never paste user input straight into the query
var sortColumns = map[string]string{
	"id":         "id",
	"name":       "name",
	"department": "department",
	"salary":     "salary",
}

func (r *PostgresRepository) List(f model.ListFilter) ([]model.Employee, int, error) {
	var conditions []string
	var args []interface{}

	if f.Department != "" {
		args = append(args, f.Department)
		conditions = append(conditions, fmt.Sprintf("department ILIKE $%d", len(args)))
	}
	if f.Search != "" {
		args = append(args, "%"+f.Search+"%")
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(args)))
	}
	if f.MinSalary > 0 {
		args = append(args, f.MinSalary)
		conditions = append(conditions, fmt.Sprintf("salary >= $%d", len(args)))
	}
	if f.MaxSalary > 0 {
		args = append(args, f.MaxSalary)
		conditions = append(conditions, fmt.Sprintf("salary <= $%d", len(args)))
	}

	where := ""
	if len(conditions) > 0 {
		where = " WHERE " + strings.Join(conditions, " AND ")
	}

	var total int
	if err := r.db.QueryRow("SELECT COUNT(*) FROM employees"+where, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count employees: %w", err)
	}

	sortColumn, ok := sortColumns[f.SortBy]
	if !ok {
		sortColumn = "id"
	}
	order := "ASC"
	if strings.EqualFold(f.Order, "desc") {
		order = "DESC"
	}

	args = append(args, f.Limit, (f.Page-1)*f.Limit)
	query := fmt.Sprintf(
		"SELECT id, name, email, department, salary FROM employees%s ORDER BY %s %s LIMIT $%d OFFSET $%d",
		where, sortColumn, order, len(args)-1, len(args),
	)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list employees: %w", err)
	}
	defer rows.Close()

	employees := []model.Employee{}
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, 0, fmt.Errorf("list employees: %w", err)
		}
		employees = append(employees, e)
	}
	return employees, total, rows.Err()
}

func (r *PostgresRepository) Update(e model.Employee) error {
	result, err := r.db.Exec(
		`UPDATE employees SET name = $1, email = $2, department = $3, salary = $4 WHERE id = $5`,
		e.Name, e.Email, e.Department, e.Salary, e.ID,
	)
	if isDuplicate(err) {
		return ErrDuplicateEmail
	}
	if err != nil {
		return fmt.Errorf("update employee: %w", err)
	}
	if affected, _ := result.RowsAffected(); affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *PostgresRepository) Delete(id int) error {
	result, err := r.db.Exec(`DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}
	if affected, _ := result.RowsAffected(); affected == 0 {
		return ErrNotFound
	}
	return nil
}
