
package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"secure-gin-api/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNotFound is returned when an employee does not exist.
var ErrNotFound = errors.New("employee not found")

// EmployeeRepository communicates with PostgreSQL.
type EmployeeRepository struct {
	DB *pgxpool.Pool
}

// NewEmployeeRepository creates the repository.
func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{
		DB: db,
	}
}

// Create inserts a new employee into the database.
func (r *EmployeeRepository) Create(
	ctx context.Context,
	e model.EmployeeInput,
) (model.Employee, error) {
	var out model.Employee

	query := `
		INSERT INTO employees (name, email, department, salary)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, department, salary
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		strings.TrimSpace(e.Name),
		strings.TrimSpace(e.Email),
		strings.TrimSpace(e.Department),
		e.Salary,
	).Scan(
		&out.ID,
		&out.Name,
		&out.Email,
		&out.Department,
		&out.Salary,
	)

	return out, err
}

// Get retrieves one employee using their ID.
func (r *EmployeeRepository) Get(
	ctx context.Context,
	id int64,
) (model.Employee, error) {
	var e model.Employee

	query := `
		SELECT id, name, email, department, salary
		FROM employees
		WHERE id = $1
	`

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&e.ID,
		&e.Name,
		&e.Email,
		&e.Department,
		&e.Salary,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return e, ErrNotFound
	}

	return e, err
}

// List retrieves employees using pagination, sorting and total count.
func (r *EmployeeRepository) List(
	ctx context.Context,
	page, limit int,
	sort, order string,
) ([]model.Employee, int, error) {
	// Only allow known column names for sorting.
	allowedColumns := map[string]string{
		"id":         "id",
		"name":       "name",
		"email":      "email",
		"department": "department",
		"salary":     "salary",
	}

	column, ok := allowedColumns[sort]
	if !ok {
		column = "id"
	}

	direction := "ASC"
	if strings.EqualFold(order, "desc") {
		direction = "DESC"
	}

	// Count all employees.
	var total int

	err := r.DB.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM employees`,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	// Calculate the offset for pagination.
	offset := (page - 1) * limit

	// Column and direction come from the allowlists above.
	query := fmt.Sprintf(`
		SELECT id, name, email, department, salary
		FROM employees
		ORDER BY %s %s
		LIMIT $1 OFFSET $2
	`, column, direction)

	rows, err := r.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)

	for rows.Next() {
		var e model.Employee

		err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Email,
			&e.Department,
			&e.Salary,
		)
		if err != nil {
			return nil, 0, err
		}

		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

// Update modifies an existing employee.
func (r *EmployeeRepository) Update(
	ctx context.Context,
	id int64,
	e model.EmployeeInput,
) (model.Employee, error) {
	var out model.Employee

	query := `
		UPDATE employees
		SET name = $1,
		    email = $2,
		    department = $3,
		    salary = $4
		WHERE id = $5
		RETURNING id, name, email, department, salary
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		strings.TrimSpace(e.Name),
		strings.TrimSpace(e.Email),
		strings.TrimSpace(e.Department),
		e.Salary,
		id,
	).Scan(
		&out.ID,
		&out.Name,
		&out.Email,
		&out.Department,
		&out.Salary,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return out, ErrNotFound
	}

	return out, err
}

// Delete removes an employee using their ID.
func (r *EmployeeRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	tag, err := r.DB.Exec(
		ctx,
		`DELETE FROM employees WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

// User represents a user account, if your application uses a users table.
type User struct {
	ID           int64
	Username     string
	PasswordHash string
	Role         string
}

// FindUser retrieves a user by username.
// This method requires a users table with the specified columns.
func (r *EmployeeRepository) FindUser(
	ctx context.Context,
	username string,
) (User, error) {
	var u User

	query := `
		SELECT id, username, password_hash, role
		FROM users
		WHERE username = $1
	`

	err := r.DB.QueryRow(ctx, query, username).Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.Role,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return u, ErrNotFound
	}

	return u, err
}
